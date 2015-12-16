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
	"path"
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
			}

			_ = product

			file, header, err:= r.FormFile("image")
			if err != nil {
				panic(err)
			}

			filename := utils.RandomString(10) + path.Ext(header.Filename)

			// If write is success then add image details to db
			if utils.Write(file, filename) {
				models.AddImageLocationToDb(api, filename, header.Filename, header.Filename, token)
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
		result := models.GetProducts(api)
		data, err := json.Marshal(result)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(200)
		res.Write(data)
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
}