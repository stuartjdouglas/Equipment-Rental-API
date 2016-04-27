package routes

import (
	"github.com/remony/Equipment-Rental-API/core/router"
	"github.com/zenazn/goji/web"
	"net/http"
	"github.com/remony/Equipment-Rental-API/core/models"
	"encoding/json"
)

type Error struct {
	Message string `json:"message"`
}

func generateImageRoutes(api router.API) {
	api.Router.Post("/product/:pid/image/add", func(c web.C, res http.ResponseWriter, req *http.Request) {
		pid := c.URLParams["pid"]
		fileType := req.FormValue("filetype")
		image := req.FormValue("image")
		if len(pid) > 3 {
			result := models.AddImageToProduct(api, pid, req.Header.Get("token"), fileType, image)
			var message hello
			if (result) {
				res.WriteHeader(200)
				message.Message = "Image added"
			} else {
				res.WriteHeader(http.StatusInternalServerError)
				message.Message = "Unable to add image"
			}
			data, err := json.Marshal(message)
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}

			res.Header().Set("Content-Type", "application/json")

			res.Write(data)
		} else {
			message := hello{
				Message: "No product exist",
			}
			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(res).Encode(message)
		}
	})

	api.Router.Delete("/product/:pid/image/:title/delete", func(c web.C, res http.ResponseWriter, req *http.Request) {
		pid := c.URLParams["pid"]
		title := c.URLParams["title"]

		if len(pid) > 3 {
			result := models.RemoveImage(api, pid, title, req.Header.Get("token"))
			var message hello
			if (result) {
				res.WriteHeader(200)
				message.Message = "Image added"
			} else {
				res.WriteHeader(http.StatusInternalServerError)
				message.Message = "Unable to add image"
			}
			data, err := json.Marshal(message)
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}

			res.Header().Set("Content-Type", "application/json")

			res.Write(data)
		} else {
			message := hello{
				Message: "No product exist",
			}
			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(res).Encode(message)
		}
	})

}
