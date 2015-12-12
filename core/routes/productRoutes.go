package routes
import (
	"github.com/zenazn/goji/web"
	"net/http"
	"encoding/json"
	"github.com/remony/Equipment-Rental-API/core/router"
	"github.com/remony/Equipment-Rental-API/core/models"
)

type Product struct {
	Title 				string 	`json:"title"`
	Description			string	`json:"description"`
	Rental_period_limit int 	`json:"rental_period_limit"`
}

func generateProductRoutes (api router.API) {
	api.Router.Post("/p", func (c web.C, res http.ResponseWriter, r *http.Request) {
		if r.Header.Get("token") != "" {
			product := Product{}
			err := json.NewDecoder(r.Body).Decode(&product)
			if err != nil {
				http.Error(res, err.Error(), http.StatusBadRequest)
			}

//			fmt.Println(product.Rental_period_limit)

			if len(product.Title) > 140 {
				res.Header().Set("Content-Type", "application/json")
				res.WriteHeader(http.StatusNotAcceptable)
				json.NewEncoder(res).Encode(error_response{Message:"Product not created: Title too long"})
			}

			//if !models.CheckIfProductExists(api, slug) {
				if models.CreateProduct(api, product.Title, product.Description, product.Rental_period_limit,  r.Header.Get("token")) {
					res.Header().Set("Content-Type", "application/json")
					res.WriteHeader(http.StatusCreated)
					json.NewEncoder(res).Encode(error_response{Message:"Product Created"})
				} else {
					res.Header().Set("Content-Type", "application/json")
					res.WriteHeader(http.StatusInternalServerError)
					json.NewEncoder(res).Encode(error_response{Message:"Product not created: Something went wrong"})
				}
//			} else {
//				res.Header().Set("Content-Type", "application/json")
//				res.WriteHeader(http.StatusConflict)
//				json.NewEncoder(res).Encode(error_response{Message:"Product not created: Already exists"})
//			}
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