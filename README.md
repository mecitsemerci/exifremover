# EXIF Metadata Remover

A simple web application that removes EXIF metadata from uploaded images.

## Features

* Upload images via a web browser.
* Uploaded images are displayed on the page.
* EXIF metadata is removed from the uploaded images before they are displayed.

## Getting Started

### Prerequisites

* **Go:** Make sure you have Go installed on your system. You can download it from [https://go.dev/](https://go.dev/).
* **Docker (Optional):** If you prefer to run the application in a Docker container, you'll need Docker installed.

### Running the Application

1. **Clone the repository:**

   ```bash
   git clone https://github.com/mecitsemerci/exifremover.git
   cd exifremover
   ```

2. **Build the project:**

   ```bash
   go build
   ```

3. **Run the application:**

   ```bash
   ./exifremover
   ```

    or
    

    ```bash
    make docker-run
    ```

   The application will start listening on port 8080 by default. You can access it in your web browser by navigating to `http://localhost:8080`.

4. **Upload an image:**

   Click on the "Choose File" button, select an image file from your computer, and click on the "Upload" button.

5. **View the uploaded image with EXIF metadata removed:**

   After the upload is complete, you will see the uploaded image displayed on the page. The EXIF metadata has been removed from the image before it is displayed.

## Environment Variables

You can customize the behavior of the application by setting environment variables. Here are some available options:

- `PORT`: The port number on which the application will listen. Default is `8080`.

## Acknowledgments

This project was inspired by the need to remove EXIF metadata from images for privacy and security reasons. The code uses the [goexif](https://github.com/rwcarlsen/goexif) library for reading and writing EXIF metadata, and the [imaging](https://github.com/disintegration/imaging) library for image processing.