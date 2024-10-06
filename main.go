package main

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"log"
	"net/http"

	"github.com/disintegration/imaging"
	"github.com/rwcarlsen/goexif/exif"
)

// Upload handler to process the image and return it as a downloadable file
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
		http.Error(w, "EXIF data not found", http.StatusBadRequest)
        return
	}

	// Seek the file back to the beginning to read the image for further processing
	file.Seek(0, 0)

	// Decode the image
	img, err := imaging.Decode(file)
	if err != nil {
		http.Error(w, "Failed to decode image", http.StatusInternalServerError)
		return
	}

	// Prepare the cleaned image in memory
	var imgBuffer bytes.Buffer
	err = jpeg.Encode(&imgBuffer, img, nil)
	if err != nil {
		http.Error(w, "Failed to encode image", http.StatusInternalServerError)
		return
	}
	// Set headers to download the image
	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Disposition", "attachment; filename=cleaned_image.jpg")
	w.Header().Set("Cache-Control", "no-store")

	// Write the image to the response
	w.Write(imgBuffer.Bytes())

	// Optionally: Log whether EXIF data was found
	log.Printf("EXIF found: %v\n", true)
	log.Print(ex)

}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")
}

func main() {
	// Serve the HTML page on the root path
	http.HandleFunc("/", htmlHandler)

	// Handle file upload and EXIF removal
	http.HandleFunc("/upload", uploadHandler)

	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}