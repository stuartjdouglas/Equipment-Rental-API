package models

import (
	"github.com/remony/Equipment-Rental-API/core/router"
	"github.com/remony/Equipment-Rental-API/core/database"
	"io"
	"github.com/remony/Equipment-Rental-API/core/utils"
	"encoding/base64"
	"strings"
	"log"
	"bytes"
)

func IsImageAvailable(api router.API, url string) bool {
	return database.DoesImageExist(api, url)
}

func RemoveImage(api router.API, pid string, title string, token string) bool {
	if IsOwner(api, token, pid) {
		utils.BinFiles("image", title)
		return database.DeleteImage(api, title);
	}
	return false

}

func AddImageToProduct(api router.API, pid string, token string, Filetype string, Image string) bool {
	if IsOwner(api, token, pid) {
		var file io.Reader

		product := GetProductFromID(api, pid, token)

		imageCode := utils.RandomString(10) // create random string

		for imageExists(api, imageCode) {
			// For each time the file exists
			imageCode = utils.RandomString(10)        // create new random string
		}

		if (Filetype != "") {
			file = base64.NewDecoder(base64.StdEncoding, strings.NewReader(Image))
		} else {
			mime := strings.SplitN(Image, ",", 2)
			mime = strings.SplitN(string(mime[0]), ":", 2)
			mime = strings.SplitN(mime[1], ";", 2)
			Filetype = mime[0]

			b64data := Image[strings.IndexByte(Image, ',') + 1:]

			data, err := base64.StdEncoding.DecodeString(b64data)
			if err != nil {
				log.Println("error:", err)
			}

			file = bytes.NewReader(data)
		}

		var fileExt string

		if (Filetype == "image/jpeg") {
			fileExt = ".jpg"
		} else if (Filetype == "image/gif") {
			fileExt = ".gif"
		} else if (Filetype == "image/png") {
			fileExt = ".png"
		}

		filename := imageCode + fileExt
		// If write is success then add image details to db
		if utils.WriteBase64Image(file, Filetype, imageCode, fileExt) {
			database.AddImageToProduct(api, filename, token, product.Items[0].Product_id)
			return true
		} else {
			// Otherwise we should call is nil
			filename = "nil"
		}

	}
	return false
}