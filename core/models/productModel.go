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

type Product struct {
	Title 			string 	`json:"title"`
	Description		string	`json:"description"`
	Rental_period_limit 	int 	`json:"rental_period_limit"`
	Image			string 	`json:"image"`
	Filetype		string 	`json:"filetype"`
}

func ValidToken(token string) bool {
	return true
}

func imageExists(api router.API, imageCode string) bool {
	return database.DoesImageExist(api, imageCode)
}

func GetUsername(api router.API, userid int) string {
	return database.GetUsername(api, userid)
}

func CreateProduct(api router.API, product Product, token string) database.Items {
	var file io.Reader

	imageCode := utils.RandomString(10) // create random string

	for imageExists(api, imageCode) { // For each time the file exists
		imageCode = utils.RandomString(10)	// create new random string
	}

	if (product.Filetype != "") {
		file = base64.NewDecoder(base64.StdEncoding, strings.NewReader(product.Image))
	} else {
		mime := strings.SplitN(product.Image, ",", 2)
		mime = strings.SplitN(string(mime[0]), ":", 2)
		mime = strings.SplitN(mime[1], ";", 2)
		product.Filetype = mime[0]

		b64data := product.Image[strings.IndexByte(product.Image, ',')+1:]

		data, err := base64.StdEncoding.DecodeString(b64data)
		if err != nil {
			log.Println("error:", err)
		}

		file = bytes.NewReader(data)
	}

	var fileExt string

	if (product.Filetype == "image/jpeg") {
		fileExt = ".jpg"
	} else if (product.Filetype == "image/gif") {
		fileExt = ".gif"
	} else if (product.Filetype == "image/png") {
		fileExt = ".png"
	}

	filename := imageCode + fileExt
	// If write is success then add image details to db
	if utils.WriteBase64Image(file, product.Filetype, imageCode, fileExt) {
		database.AddImageLocationToDb(api, filename, filename, filename, token)
	} else {
		// Otherwise we should call is nil
		filename = "nil"
	}
	product_id := utils.GenerateUUID();
	database.CreateProduct(api, product.Title, product.Description, product.Rental_period_limit, token, filename, product_id)

	return database.GetProductFromID(api, product_id)
}

func IsOwner(api router.API, token string, product_id string) bool {
	return database.IsOwner(api, token, product_id)
}

func GetProductFromID(api router.API, product_id string) database.Items {
	return database.GetProductFromID(api, product_id)
}

func RemoveProduct(api router.API, product_id string, token string, item database.Items) bool {
	database.RemoveProduct(api, product_id, token)
	utils.BinFiles("image", item.Items[0].Image.Title)
	return true;
}

func GetCurrentlyRentedProducts(api router.API, token string, step int, count int) database.RentResult {
	return database.GetCurrentlyRentedProducts(api, token, step, count)
}

func GetProducts(api router.API) database.Items {
	return database.GetProducts(api)
}

func GetAuthedAvailability(api router.API, product_id string, token string) database.RentalStatus {
	return database.GetAuthedAvailability(api, product_id, token)
}

func GetAvailability(api router.API, product_id string) database.Availability {
	return database.GetAvailability(api, product_id)
}

func GetOwnerProductAvailability(api router.API, product_id string, token string) database.OwnerRentalStatus {
	return database.GetOwnerProductAvailability(api, product_id, token)
}

func RentItem(api router.API, product_id string, token string) database.Availability {
	return database.RentItem(api, product_id, token)
}

func ReturnItem(api router.API, product_id string, token string) {
	if (database.IsOwner(api, token, product_id)) {
		log.Println("returning as owner")
		database.ReturnItemAsOwner(api, product_id, token)
	} else {
		log.Println("returning as user")
		database.ReturnItem(api, product_id, token)
	}
}

func GetProductsPaging(api router.API, step int, count int) database.Items {
	return database.GetProductsPaging(api, step, count)
}

func GetOwnerProductsPaging(api router.API, token string, step int, count int) database.Items {
	return database.GetOwnerProductsPaging(api, token, step, count)
}

func GetProductFromOwner(api router.API, username string) database.Items{
	return database.GetProductFromOwner(api, username)
}
