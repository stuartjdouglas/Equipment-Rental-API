package routes

import (
	"github.com/zenazn/goji/web"
	"net/http"
	"github.com/remony/Equipment-Rental-API/core/router"
	"github.com/remony/Equipment-Rental-API/core/models"
	"encoding/json"
)

func generateSearchRoutes(api router.API) {
	api.Router.Get("/search/:terms", func(c web.C, res http.ResponseWriter, req *http.Request) {

		result := models.SearchTag(api, c.URLParams["terms"], req.Header.Get("start"), req.Header.Get("count"))
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
