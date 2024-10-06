package main

import (
	"encoding/json"
	"fmt"
	"image/jpeg"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/disintegration/imaging"
	"github.com/rwcarlsen/goexif/exif"
)

// Upload handler to process the image and return JSON response
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the multipart form
	err := r.ParseMultipartForm(10 << 20) // Limit to 10MB
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Get the uploaded file
	file, _, err := r.FormFile("photo")
	if err != nil {
		http.Error(w, "Unable to read file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Check if the image contains EXIF data
	ex, err := exif.Decode(file)

	if err != nil {
		response := map[string]interface{}{
			"status":  "failed",
			"message": "EXIF data is not available",
			"error": err.Error(),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return

	}
	log.Println(ex)

	// Seek the file back to the beginning to read the image for further processing
	file.Seek(0, 0)

	// Decode the image
	img, err := imaging.Decode(file)
	if err != nil {
		http.Error(w, "Failed to decode image", http.StatusInternalServerError)
		return
	}

	// Generate a filename and save the image without EXIF data
	timestamp := time.Now().Unix()
	fileName := fmt.Sprintf("cleaned_image_%d.jpg", timestamp)
	filePath := filepath.Join(".", "uploads", fileName)

	outFile, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Failed to save image", http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	// Encode and save the cleaned image
	err = jpeg.Encode(outFile, img, nil)
	if err != nil {
		http.Error(w, "Failed to encode image", http.StatusInternalServerError)
		return
	}

	// Create JSON response
	response := map[string]interface{}{
		"status":   "success",
		"imageUrl": fmt.Sprintf("/uploads/%s", fileName),
		"message":  "EXIF metadata removed successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}

func main() {
	// Ensure the upload directory exists
	err := os.MkdirAll("./uploads", os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create upload directory: %v", err)
	}

	// Serve static files for uploads
	http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	// Serve the HTML page on the root path
	http.HandleFunc("/", htmlHandler)

	// Handle file upload and EXIF removal
	http.HandleFunc("/upload", uploadHandler)

	fmt.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
