package routes

import (
	"github.com/remony/Equipment-Rental-API/core/router"
	"github.com/zenazn/goji/web"
	"net/http"
	"github.com/remony/Equipment-Rental-API/core/models"
	"encoding/json"
)

func generateRootRoutes(api router.API) {
	api.Router.Get("/", func (c web.C, res http.ResponseWriter, r *http.Request) {
		result := models.GetIndex(api)
		data, err := json.Marshal(result)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(200)
		res.Write(data)
	})
}
