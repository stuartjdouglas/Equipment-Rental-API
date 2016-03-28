package models

import (
	"github.com/remony/Equipment-Rental-API/core/router"
	"github.com/remony/Equipment-Rental-API/core/database"
	"io"
	"github.com/remony/Equipment-Rental-API/core/utils"
	"encoding/base64"
	"strings"
	"bytes"
	"gitlab.com/remon/lemon-swear-detector"
	"html"
	"log"
)

type Product struct {
	Title                     string        `json:"title"`
	Description               string        `json:"description"`
	Rental_period_limit       int        `json:"rental_period_limit"`
	Image                     string        `json:"image"`
	Filetype                  string        `json:"filetype"`
	Condition                 string        `json:"condition"`
	Comments_enabled          bool `json:"comments_enabled"`
	Comments_require_approval bool `json:"comments_require_approval"`
	Content                   string `json:"content"`
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

	for imageExists(api, imageCode) {
		// For each time the file exists
		imageCode = utils.RandomString(10)        // create new random string
	}

	if (product.Filetype != "") {
		file = base64.NewDecoder(base64.StdEncoding, strings.NewReader(product.Image))
	} else {
		mime := strings.SplitN(product.Image, ",", 2)
		mime = strings.SplitN(string(mime[0]), ":", 2)
		mime = strings.SplitN(mime[1], ";", 2)
		product.Filetype = mime[0]

		b64data := product.Image[strings.IndexByte(product.Image, ',') + 1:]

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

	//imageGood := utils.CheckImageIsSafe(file, product.Filetype)
	filename := imageCode + fileExt
	// If write is success then add image details to db
	if utils.WriteBase64Image(file, product.Filetype, imageCode, fileExt) {
		database.AddImageLocationToDb(api, filename, filename, filename, token)
	} else {
		// Otherwise we should call is nil
		filename = "nil"
	}
	product_id := utils.GenerateUUID();

	requires_approval := lemon_swear_detector.CheckSentence(product.Title + " " + product.Description)

	if len(product.Content) <= 0 {
		product.Content = " "
	}

	product.Title = html.EscapeString(product.Title)
	product.Description = html.EscapeString(product.Description)
	product.Condition = html.EscapeString(product.Condition)
	product.Content = html.EscapeString(product.Content)

	database.CreateProduct(api, product.Title, product.Description, product.Rental_period_limit, token, filename, product_id, product.Condition, requires_approval, product.Content)

	return database.GetProductFromID(api, product_id, token)
}

func IsOwner(api router.API, token string, product_id string) bool {
	return database.IsOwner(api, token, product_id)
}

func GetProductFromID(api router.API, product_id string, token string) database.Items {
	return database.GetProductFromID(api, product_id, token)
}

func RemoveProduct(api router.API, product_id string, token string, item database.Items) bool {

	for i := 0; i < len(item.Items[0].Images); i++ {
		utils.BinFiles("image", item.Items[0].Images[i].Title)
		database.RemoveImages(api, product_id)
	}
	database.RemoveProduct(api, product_id, token)

	return true;
}

func GetCurrentlyRentedProducts(api router.API, token string, step int, count int) database.RentResult {
	return database.GetCurrentlyRentedProducts(api, token, step, count)
}

func GetProducts(api router.API) database.Items {
	//return database.GetProducts(api)
	return database.Items{}
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
		database.ReturnItemAsOwner(api, product_id, token)
	} else {
		database.ReturnItem(api, product_id, token)
	}
}

func GetProductsPaging(api router.API, step int, count int, token string, order bool) database.Items {
	return database.GetProductsPaging(api, step, count, token, order)
}

func GetProductsPagingSortedByAdded(api router.API, step int, count int, token string, order bool) database.Items {
	return database.GetProductsPagingSortedByAdded(api, step, count, token, order)
}

func GetProductsPagingSortedByUpdated(api router.API, step int, count int, token string, order bool) database.Items {
	return database.GetProductsPagingSortedByUpdated(api, step, count, token, order)
}
func GetProductsPagingSortedByLikes(api router.API, step int, count int, token string, order bool) database.Items {
	return database.GetProductsPagingSortedByLikes(api, step, count, token, order)
}
func GetProductsPagingRandom(api router.API, step int, count int, token string) database.Items {
	return database.GetProductsPagingRandom(api, step, count, token)
}

func GetOwnerProductsPaging(api router.API, token string, step int, count int) database.OwnerItems {
	return database.GetOwnerProductsPaging(api, token, step, count)
}

func GetProductFromOwner(api router.API, username string) database.Items {
	return database.GetProductFromOwner(api, username)
}

type OwnerRes struct {
	Owner bool `json:"owner"`
}

func AmITheOwner(api router.API, pid string, token string) OwnerRes {
	return OwnerRes{Owner:IsOwner(api, token, pid)}
}

func EditProduct(api router.API, pid string, product Product, token string) bool {
	if IsOwner(api, token, pid) {
		if len(product.Title) > 0 && len(product.Description) > 0 && product.Rental_period_limit > 0 {
			return database.UpdateProduct(api, pid, product.Title, product.Description, product.Rental_period_limit, product.Condition, product.Comments_enabled, product.Comments_require_approval, product.Content)
		}
	}
	return false
}