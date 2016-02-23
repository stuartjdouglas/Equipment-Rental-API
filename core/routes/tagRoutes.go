package routes

import (
	"github.com/zenazn/goji/web"
	"net/http"
	"github.com/remony/Equipment-Rental-API/core/router"
	"github.com/remony/Equipment-Rental-API/core/models"
	"encoding/json"
"strconv"
)

func generateTagRoutes(api router.API) {
	api.Router.Delete("/product/:pid/tag/:tag/remove", func(c web.C, res http.ResponseWriter, req *http.Request) {
		if models.RemoveTag(api, c.URLParams["pid"], c.URLParams["tag"], req.Header.Get("token")) {
			message := hello{
				Message: "Tag Removed",
			}
			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(200)
			json.NewEncoder(res).Encode(message)

		} else {
			message := hello{
				Message: "Tag Remove Failed",
			}
			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(res).Encode(message)
		}
	})

	api.Router.Get("/product/:pid/tags", func(c web.C, res http.ResponseWriter, req *http.Request) {
		result := models.GetTags(api, c.URLParams["pid"])
		data, err := json.Marshal(result)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(200)
		res.Write(data)
	})

	api.Router.Get("/tags/:tid", func(c web.C, res http.ResponseWriter, req *http.Request) {
		query := c.URLParams["tid"]
		if len(query) > 3 {
			result := models.GetProductsOfTag(api, query, req.Header.Get("start"), req.Header.Get("count"))
			data, err := json.Marshal(result)
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}

			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(200)
			res.Write(data)
		} else {
			message := hello{
				Message: "No Tags exist",
			}
			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(res).Encode(message)
		}

	})

	api.Router.Post("/product/:pid/tag", func(c web.C, res http.ResponseWriter, req *http.Request) {
		if models.AddTag(api, c.URLParams["pid"], req.Header.Get("tags"), req.Header.Get("token")) {
			message := hello{
				Message: "Tag Added",
			}
			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(200)
			json.NewEncoder(res).Encode(message)

		} else {
			message := hello{
				Message: "Tag Add Failed",
			}
			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(res).Encode(message)
		}
	})


	//	Get All most used tags
	api.Router.Get("/filter/tags/:order", func(c web.C, res http.ResponseWriter, r *http.Request) {

		if (r.Header.Get("Start") != "" || r.Header.Get("Count") != "") {
			var order bool

			t_order := c.URLParams["order"]

			order = t_order == "popular"

			step, err := strconv.Atoi(r.Header.Get("Start"))
			count, err := strconv.Atoi(r.Header.Get("Count"))
			token := r.Header.Get("token")
			result := models.GetTagsMostUsed(api, step, count, token, order)

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

}
