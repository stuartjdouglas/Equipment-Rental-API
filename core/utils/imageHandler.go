package utils

import (
  "fmt"
	"io/ioutil"
	"os"
	"io"
	"mime/multipart"
)

func writeToFile(image byte, filename string) {
  fmt.Println(filename);
  path:= "client/images/"
  ioutil.WriteFile(path + filename, []byte("hello"), 0644)
  fmt.Println("write file to desk");
}

// Write writes a given file to the file system
func Write(image multipart.File, filename string) bool {
	f, err := os.OpenFile("data/images/" + filename, os.O_WRONLY | os.O_CREATE, 0666)
  	if err != nil {
  	  return false;
  	}
	defer f.Close()
	// Write the file to the filesystem
	io.Copy(f, image)
  	image.Close()

	//success
	return true;
}
