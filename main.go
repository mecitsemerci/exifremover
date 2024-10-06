package main

import (
	"bytes"
	"image/jpeg"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/disintegration/imaging"
	"github.com/rwcarlsen/goexif/exif"
)

var logger *slog.Logger

func init() {
	logger = slog.Default()
}

func main() {

	// Serve the HTML page on the root path
	http.HandleFunc("/", htmlHandler)

	// Handle file upload and EXIF removal
	http.HandleFunc("/upload", uploadHandler)

	// Start the server on port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server listening on port %s", port)

	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Fatal("Server error:", err)
	}
}

// Upload handler to process the image and return it as a downloadable file
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		logger.Error("Invalid request method", "method", r.Method)
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the multipart form
	err := r.ParseMultipartForm(10 << 20) // Limit to 10MB
	if err != nil {
		logger.Error("Failed to parse form", "error", err)
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Get the uploaded file
	file, _, err := r.FormFile("photo")
	if err != nil {
		logger.Error("Failed to get file from form", "error", err)
		http.Error(w, "Unable to read file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Check if the image contains EXIF data
	ex, err := exif.Decode(file)

	if err != nil {
		logger.Error("Failed to decode EXIF data", "error", err)
		http.Error(w, "EXIF data not found", http.StatusBadRequest)
		return
	}
	logger.Info("EXIF data found", slog.Any("metadata", ex))

	// Seek the file back to the beginning to read the image for further processing
	file.Seek(0, 0)

	// Decode the image
	img, err := imaging.Decode(file)

	if err != nil {
		logger.Error("Failed to decode image", "error", err)
		http.Error(w, "Failed to decode image", http.StatusInternalServerError)
		return
	}

	// Prepare the cleaned image in memory
	var imgBuffer bytes.Buffer
	err = jpeg.Encode(&imgBuffer, img, nil)
	if err != nil {
		logger.Error("Failed to encode image", "error", err)
		http.Error(w, "Failed to encode image", http.StatusInternalServerError)
		return
	}
	// Set headers to download the image
	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Disposition", "attachment; filename=cleaned_image.jpg")
	w.Header().Set("Cache-Control", "no-store")

	// Write the image to the response
	w.Write(imgBuffer.Bytes())

	logger.Info("EXIF image removed successfully")
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./views/index.html")
}
