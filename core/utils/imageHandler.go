package utils

import (
  "fmt"
	"os"
	"github.com/nfnt/resize"
	"log"
	"image"
	"image/jpeg"
	"image/png"
	"io"
"strings"
	"image/gif"
)

func WriteBase64Image (filez io.Reader, filetype string, filename string, extension string) bool {
	if (filetype == "image/jpeg") {
		file, err := jpeg.Decode(filez)
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
		file, err := gif.DecodeAll(filez)
		if err != nil {
			log.Println(err)
		}
		// write original
		writeGIF(file, filename + extension, filetype)


	} else if (filetype == "image/png") {
		file, err := png.Decode(filez)
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
		file, err := jpeg.Decode(filez)
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
	return true;
}

func ResizeImage(image image.Image, height uint) image.Image {
	return resize.Resize(0, height, image, resize.Lanczos3)
}

// Write writes a given file to the file system
func Write(image image.Image, filename string, filetype string) bool {
	if (filetype == "image/jpeg") {
		writeJPEG(image, filename, 100)
	} else if (filetype == "image/gif") {
		//writeGIF(image, filename, filetype)
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
		//writeGIF(image, filename, filetype)
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

func writeGIF(image *gif.GIF, filename string, filetype string) {

	out, err := os.Create(imagelocation + filename)
	if err != nil {
		log.Println(err)
	}

	var opt gif.Options
	opt.NumColors = 256

	err = gif.EncodeAll(out, image) // put num of colors to 256
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

func BinFiles(file_type string, filename string) {
	if (file_type == "image") {
		log.Println("Handling Images")
		log.Println(getFilename(filename))
		moveImages(getFilename(filename), getFileExt(filename));

	}
}

func moveImages(filename string, ext string) {

	err :=  os.Rename("data/images/" + filename + "." + ext, "data/images/bin/" + filename + "." + ext)
	err =  os.Rename("data/images/" + filename + "_large." + ext, "data/images/bin/" + filename + "_large." + ext)
	err =  os.Rename("data/images/" + filename + "_medium." + ext, "data/images/bin/" + filename + "_medium." + ext)
	err =  os.Rename("data/images/" + filename + "_small." + ext, "data/images/bin/" + filename + "_small." + ext)
	err =  os.Rename("data/images/" + filename + "_thumb." + ext, "data/images/bin/" + filename + "_thumb." + ext)

	if err != nil {
		fmt.Println(err)
		return
	}
}

func getFileExt(data string) string {
	splitfile := strings.Split(data, ".")
	return splitfile[len(splitfile) - 1]
}

func getFilename(data string) string {
	splitfile := strings.Split(data, ".")
	splitfile = strings.Split(splitfile[0], "/")
	return splitfile[len(splitfile)- 1]
}

func getFullFilename(data string) string {
	splitfile := strings.Split(data, "/")
	return splitfile[len(splitfile)- 1]
}