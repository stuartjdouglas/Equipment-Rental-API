package routes

import (
	"github.com/remony/Equipment-Rental-API/core/router"
	"github.com/zenazn/goji/web"
	"net/http"
	"github.com/remony/Equipment-Rental-API/core/models"
	"encoding/json"
	"log"
)

func generateAdminRoutes(api router.API) {
	api.Router.Post("/", func(c web.C, res http.ResponseWriter, r *http.Request) {
		log.Println(r.Body)
		title := r.FormValue("title")
		description := r.FormValue("description")
		token := r.Header.Get("token")
		result := models.UpdateIndex(api, title, description, token)
		data, err := json.Marshal(result)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(200)
		res.Write(data)
	})
}
