package routes
import (
	"github.com/zenazn/goji/web"
	"net/http"
	"encoding/json"
	"github.com/remony/Equipment-Rental-API/core/router"
	"github.com/remony/Equipment-Rental-API/core/models"
	"log"
	"github.com/remony/Equipment-Rental-API/core/utils"
	"strconv"
	"github.com/remony/Equipment-Rental-API/core/models/sessions"
	"encoding/base64"
	"strings"
)

type Product struct {
	Title 			string 	`json:"title"`
	Description		string	`json:"description"`
	Rental_period_limit 	int 	`json:"rental_period_limit"`
	Image			string 	`json:"image"`
}

func generateProductRoutes (api router.API) {
	api.Router.Post("/p", func (c web.C, res http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")
		if token != "" {

			limit, err := strconv.Atoi(r.FormValue("rental_period_limit"))

			if err != nil {
				log.Println(err)
			}
			product := Product {
				Title:r.FormValue("title"),
				Description:r.FormValue("description"),
				Rental_period_limit:limit,
				Image:r.FormValue("image"),
			}

			log.Println(product.Title)
			log.Println(product.Description)
			log.Println(product.Rental_period_limit)
			log.Println(product.Image)

			_ = product

//			file, header, err:= r.FormFile("image")
//			if err != nil {
//				panic(err)
//			}
			file := base64.NewDecoder(base64.StdEncoding, strings.NewReader(product.Image))




			imageCode := utils.RandomString(10) // create random string

			for models.DoesImageExist(api, imageCode) { // For each time the file exists
				imageCode = utils.RandomString(10)	// create new random string
			}

			fileExt := ".jpg"
			filename := imageCode + fileExt
			// If write is success then add image details to db
			if utils.WriteBase64Image(file, nil, imageCode, fileExt) {
				models.AddImageLocationToDb(api, filename, filename, filename, token)
			} else {
				// Otherwise we should call is nil
				filename = "nil"
			}

			models.CreateProduct(api, product.Title, product.Description, product.Rental_period_limit, token, filename)

		} else {
			http.Error(res, "", http.StatusUnauthorized)
		}
	})



	//	Get all Products
	api.Router.Get("/p", func (c web.C, res http.ResponseWriter, r *http.Request) {

		if (r.Header.Get("Start") != ""|| r.Header.Get("Count") != "") {
			step, err :=  strconv.Atoi(r.Header.Get("Start"))
			count, err :=  strconv.Atoi(r.Header.Get("Count"))

			result := models.GetProductsPaging(api, step, count)
			data, err := json.Marshal(result)
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}

			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(200)
			res.Write(data)
		} else {
			result := models.GetProducts(api)
			data, err := json.Marshal(result)
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}

			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(200)
			res.Write(data)
		}



	})





	api.Router.Get("/products/:username", func (c web.C, res http.ResponseWriter, r *http.Request) {

		result := models.GetProductFromOwner(api, c.URLParams["username"])

		data, err := json.Marshal(result)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(200)
		res.Write(data)
	})

	api.Router.Get("/product/:id", func (c web.C, res http.ResponseWriter, r *http.Request) {

		result := models.GetProductFromID(api, c.URLParams["id"])

		data, err := json.Marshal(result)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(200)
		res.Write(data)
	})


	//	Get all Currently Rented Products
	api.Router.Get("/p/rent/current", func (c web.C, res http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")
		if (r.Header.Get("Step") != ""|| r.Header.Get("Count") != "") {
			step, err :=  strconv.Atoi(r.Header.Get("Step"))
			count, err :=  strconv.Atoi(r.Header.Get("Count"))

			result := models.GetCurrentlyRentedProducts(api, token, step, count)
			data, err := json.Marshal(result)
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}

			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(200)
			res.Write(data)
		} else {
			result := models.GetProducts(api)
			data, err := json.Marshal(result)
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}

			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(200)
			res.Write(data)
		}
	})

	//	Get all Currently Rented Products
	api.Router.Get("/p/rent/past", func (c web.C, res http.ResponseWriter, r *http.Request) {
//		token := r.Header.Get("token")
		if (r.Header.Get("Step") != ""|| r.Header.Get("Count") != "") {
//			step, err :=  strconv.Atoi(r.Header.Get("Step"))
//			count, err :=  strconv.Atoi(r.Header.Get("Count"))

//			result := models.GetPastRentedProducts(api, token, step, count)
//			data, err := json.Marshal(result)
//			if err != nil {
//				http.Error(res, err.Error(), http.StatusInternalServerError)
//				return
//			}
//
//			res.Header().Set("Content-Type", "application/json")
//			res.WriteHeader(200)
//			res.Write(data)
		} else {
			result := models.GetProducts(api)
			data, err := json.Marshal(result)
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}

			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(200)
			res.Write(data)
		}
	})

	api.Router.Get("/p/:id/availability", func (c web.C, res http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")
		log.Println(token)
		if token != "" {
			// Check that the token is a valid token
			if sessions.IsSessionValid(api, token) {
				result := models.GetAuthedAvailability(api, c.URLParams["id"], token)
				data, err := json.Marshal(result)
				if err != nil {
					http.Error(res, err.Error(), http.StatusInternalServerError)
					return
				}

				res.Header().Set("Content-Type", "application/json")
				res.WriteHeader(200)
				res.Write(data)
			}
		} else {
			result := models.GetAvailability(api, c.URLParams["id"])
			data, err := json.Marshal(result)
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}

			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(200)
			res.Write(data)
		}
//		log.Println(token)


	})

	api.Router.Post("/p/:id/rent", func (c web.C, res http.ResponseWriter, r *http.Request) {

		result := models.GetAvailability(api, c.URLParams["id"])
		token := r.Header.Get("token")



		if result.Available {
			if token != "" {
				if (sessions.IsSessionValid(api, token)) {

					//Rent the product
					log.Println(token)
					models.RentItem(api, c.URLParams["id"], token)

				} else {
					log.Println("invalid, expired token")
					http.Error(res, "Unauthorized: invalid or expired token", http.StatusInternalServerError)
				}


			} else {
				http.Error(res, "Unauthorized: missing token", http.StatusInternalServerError)
			}
		} else {
			// Return nil
			data, err := json.Marshal(result)
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}

			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(http.StatusConflict)

			res.Write(data)

		}






	})

	api.Router.Post("/p/:id/return", func (c web.C, res http.ResponseWriter, r *http.Request) {

		result := models.GetAvailability(api, c.URLParams["id"])
		token := r.Header.Get("token")



		if !result.Available {
			if token != "" {
				if (sessions.IsSessionValid(api, token)) {

					//Rent the product
					log.Println(token)
					models.ReturnItem(api, c.URLParams["id"], token)

				} else {
					log.Println("invalid, expired token")
					http.Error(res, "Unauthorized: invalid or expired token", http.StatusInternalServerError)
				}


			} else {
				http.Error(res, "Unauthorized: missing token", http.StatusInternalServerError)
			}
		} else {
			// Return nil
			data, err := json.Marshal(result)
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}

			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(http.StatusConflict)

			res.Write(data)

		}






	})

}