# EXIF Metadata Remover

A simple Go web application that removes EXIF metadata from uploaded images.

## Installation

To install the project, follow these steps:

1. Install Go (version 1.16 or higher) if you haven't already. You can download it from the [official Go website](https://golang.org/dl/).

2. Clone the project repository:

```bash
git clone https://github.com/mecitsemerci/exifremover.git
```

3. Navigate to the project directory:

```bash
cd exifremover
```

4. Build the project:

```bash
go build
```

## Usage

To run the application, execute the following command:

```bash
./exifremover
```

The application will start listening on port 8080 by default. You can access it in your web browser by navigating to `http://localhost:8080`.

To upload an image and remove its EXIF metadata, follow these steps:

1. Click on the "Choose File" button.
2. Select an image file from your computer.
3. Click on the "Upload" button.
4. After the upload is complete, you will see a download link below the image. Click on the link to download the image with the EXIF metadata removed.

## Environment Variables

You can customize the behavior of the application by setting environment variables. Here are some available options:

- `PORT`: The port number on which the application will listen. Default is `8080`.

## Acknowledgments

This project was inspired by the need to remove EXIF metadata from images for privacy and security reasons. The code uses the [goexif](https://github.com/rwcarlsen/goexif) library for reading and writing EXIF metadata, and the [imaging](https://github.com/disintegration/imaging) library for image processing.