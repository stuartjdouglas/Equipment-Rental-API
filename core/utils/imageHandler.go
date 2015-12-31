package utils

import (
  "fmt"
	"os"
	"mime/multipart"
	"github.com/nfnt/resize"
	"log"
	"image"
	"net/textproto"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
)

func WriteBase64Image (filez io.Reader, header textproto.MIMEHeader, filename string, extension string) bool {
	file, err := jpeg.Decode(filez)
	filetype := "image/jpeg"
	if err != nil {
		log.Println(err)
	}
	// Write original
	Write(file, filename + extension, filetype)
	// Write Large
	largeImage := ResizeImage(file, 1000)
	Write(largeImage, filename + "_large" + extension, filetype)

	// Write medium
	mediumImage := ResizeImage(file, 600)
	Write(mediumImage, filename + "_medium" + extension, filetype)

	// Write small
	smallImage := ResizeImage(file, 300)
	Write(smallImage, filename + "_small" + extension, filetype)

	// Write thumbnail
	thumbnail := ResizeImage(file, 150)
	Write(thumbnail, filename + "_thumb" + extension, filetype)
	return true;
}

func WriteImage (file multipart.File, header textproto.MIMEHeader, filename string, extension string) bool {
	filetype := header.Get("Content-Type")

	if (filetype == "image/jpeg") {
		file, err := jpeg.Decode(file)
		if err != nil {
			log.Println(err)
		}


		// Write original
		Write(file, filename + extension, filetype)
		// Write Large
		largeImage := ResizeImage(file, 1000)
		Write(largeImage, filename + "_large" + extension, filetype)

		// Write medium
		mediumImage := ResizeImage(file, 600)
		Write(mediumImage, filename + "_medium" + extension, filetype)

		// Write small
		smallImage := ResizeImage(file, 300)
		Write(smallImage, filename + "_small" + extension, filetype)

		// Write thumbnail
		thumbnail := ResizeImage(file, 150)
		Write(thumbnail, filename + "_thumb" + extension, filetype)

	} else if (filetype == "image/gif") {
		file, err := gif.Decode(file)
		if err != nil {
			log.Println(err)
		}
		// write original
		Write(file, filename + extension, filetype)
		// Write Large
		largeImage := ResizeImage(file, 1000)
		Write(largeImage, filename + "_large" + extension, filetype)

		// Write medium
		mediumImage := ResizeImage(file, 600)
		Write(mediumImage, filename + "_medium" + extension, filetype)

		// Write small
		smallImage := ResizeImage(file, 300)
		Write(smallImage, filename + "_small" + extension, filetype)

		// Write thumbnail
		thumbnail := ResizeImage(file, 150)
		Write(thumbnail, filename + "_thumb" + extension, filetype)


	} else if (filetype == "image/png") {
		file, err := png.Decode(file)
		if err != nil {
			log.Println(err)
		}
		// Write original
		Write(file, filename + extension, filetype)
		// Write Large
		largeImage := ResizeImage(file, 1000)
		Write(largeImage, filename + "_large" + extension, filetype)

		// Write medium
		mediumImage := ResizeImage(file, 600)
		Write(mediumImage, filename + "_medium" + extension, filetype)

		// Write small
		smallImage := ResizeImage(file, 300)
		Write(smallImage, filename + "_small" + extension, filetype)

		// Write thumbnail
		thumbnail := ResizeImage(file, 150)
		Write(thumbnail, filename + "_thumb" + extension, filetype)
	} else {
		file, err := jpeg.Decode(file)
		if err != nil {
			log.Println(err)
		}


		// Write original
		Write(file, filename + extension, filetype)
		// Write Large
		largeImage := ResizeImage(file, 1000)
		Write(largeImage, filename + "_large" + extension, filetype)

		// Write medium
		mediumImage := ResizeImage(file, 600)
		Write(mediumImage, filename + "_medium" + extension, filetype)

		// Write small
		smallImage := ResizeImage(file, 300)
		Write(smallImage, filename + "_small" + extension, filetype)

		// Write thumbnail
		thumbnail := ResizeImage(file, 150)
		Write(thumbnail, filename + "_thumb" + extension, filetype)
	}

	return true
}

func ResizeImage(image image.Image, height uint) image.Image{

	return resize.Resize(0, height, image, resize.Lanczos3)
}

// Write writes a given file to the file system
func Write(image image.Image, filename string, filetype string) bool {
	if (filetype == "image/jpeg") {
		writeJPEG(image, filename, 100)
	} else if (filetype == "image/gif") {
		writeGIF(image, filename)
	} else if (filetype == "image/png") {
		writePNG(image, filename)
	} else {
		return false;
	}
	//success
	return true
}
func WriteImageType(image image.Image, filename string, filetype string) bool {
	if (filetype == "image/jpeg") {
		writeJPEG(image, filename, 100)
	} else if (filetype == "image/gif") {
		writeGIF(image, filename)
	} else if (filetype == "image/png") {
		writePNG(image, filename)
	} else {
		return false;
	}
	//success
	return true
}

const imagelocation = "data/images/"

func writeJPEG(image image.Image, filename string, quality int) {

	out, err := os.Create(imagelocation + filename)
	if err != nil {
		log.Println(err)
	}
	var opt jpeg.Options

	opt.Quality = quality

	err = jpeg.Encode(out, image, &opt)
	if err != nil {
		log.Println(err)
	}
	out.Close()
}

func writeGIF(image image.Image, filename string) {

	out, err := os.Create(imagelocation + filename)
	if err != nil {
		log.Println(err)
	}

	var opt gif.Options
	opt.NumColors = 256

	err = gif.Encode(out, image, &opt) // put num of colors to 256
	if err != nil {
		log.Println(err)
	}
	out.Close()
}

func writePNG(image image.Image, filename string) {

	out, err := os.Create(imagelocation + filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = png.Encode(out, image)
	if err != nil {
		log.Println(err)
	}

	out.Close()
}
