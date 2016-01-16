package routes
import (
	"github.com/zenazn/goji/web"
	"net/http"
	"encoding/json"
	"github.com/remony/Equipment-Rental-API/core/router"
	"github.com/remony/Equipment-Rental-API/core/models"
	"log"
	"strconv"
)



type DeleteResponse struct {
	message string
}

func generateProductRoutes (api router.API) {
	api.Router.Post("/p", func (c web.C, res http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")

		if models.ValidToken(token) {
			limit, err := strconv.Atoi(r.FormValue("rental_period_limit"))

			if err != nil {
				log.Println(err)
			}

			product := models.Product {
				Title:r.FormValue("title"),
				Description:r.FormValue("description"),
				Rental_period_limit:limit,
				Image:r.FormValue("image"),
				Filetype:r.FormValue("filetype"),
			}



			result := models.CreateProduct(api, product, token)

			data, err := json.Marshal(result)
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}

			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(200)
			res.Write(data)



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

	api.Router.Get("/owner/products", func(c web.C, res http.ResponseWriter, req *http.Request) {
		token :=  req.Header.Get("token")
		step, err :=  strconv.Atoi(req.Header.Get("step"))
		count, err :=  strconv.Atoi(req.Header.Get("count"))

		result := models.GetOwnerProductsPaging(api, token, step, count)

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




	api.Router.Delete("/product/:id/delete", func (c web.C, res http.ResponseWriter, req *http.Request) {
		token := req.Header.Get("token")


		if (models.IsOwner(api, token, c.URLParams["id"])) {
			item := models.GetProductFromID(api, c.URLParams["id"])
			models.RemoveProduct(api, c.URLParams["id"], token, item)
			result := DeleteResponse {
				message: "Has been deleted",
			}
			data, err := json.Marshal(result)
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}

			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(200)
			res.Write(data)
		} else {
			http.Error(res, "", http.StatusUnauthorized)
		}

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
		if token != "" {
			// Check that the token is a valid token
			if models.IsSessionValid(api, token) {
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
	})

	api.Router.Get("/owner/products/:id/availability", func (c web.C, res http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")
		if token != "" {
			// Check that the token is a valid token
			if models.IsSessionValid(api, token) {
				result := models.GetOwnerProductAvailability(api, c.URLParams["id"], token)
				data, err := json.Marshal(result)
				if err != nil {
					http.Error(res, err.Error(), http.StatusInternalServerError)
					return
				}

				res.Header().Set("Content-Type", "application/json")
				res.WriteHeader(200)
				res.Write(data)
			} else {
				http.Error(res, "", http.StatusUnauthorized)
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


	})

	api.Router.Post("/p/:id/rent", func (c web.C, res http.ResponseWriter, r *http.Request) {

		result := models.GetAvailability(api, c.URLParams["id"])
		token := r.Header.Get("token")



		if result.Available {
			if token != "" {
				if (models.IsSessionValid(api, token)) {
					models.RentItem(api, c.URLParams["id"], token)
				} else {
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
				if (models.IsSessionValid(api, token)) {
					models.ReturnItem(api, c.URLParams["id"], token)

				} else {
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