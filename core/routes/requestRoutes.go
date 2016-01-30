package routes

import (
	"github.com/zenazn/goji/web"
	"github.com/remony/Equipment-Rental-API/core/router"
	"net/http"
	"encoding/json"
	"github.com/remony/Equipment-Rental-API/core/models"
)

func generateRequestRouter(api router.API) {
	api.Router.Post("/product/:pid/request", func(c web.C, res http.ResponseWriter, req *http.Request) {
		pid := c.URLParams["pid"]
		if len(pid) > 3 {
			result := models.RequestProduct(api, pid, req.Header.Get("token"))
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
			res.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(res).Encode(message)
		}
	})

	api.Router.Post("/product/:pid/request/cancel", func(c web.C, res http.ResponseWriter, req *http.Request) {
		pid := c.URLParams["pid"]
		if len(pid) > 3 {
			result := models.CancelRequestProduct(api, pid, req.Header.Get("token"))
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
			res.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(res).Encode(message)
		}
	})
}
