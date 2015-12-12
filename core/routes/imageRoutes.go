package routes

import (

	"github.com/zenazn/goji/web"
"net/http"
	"path"
	"encoding/json"
	"github/remony/Equipment-Rental-API/core/router"
	"github/remony/Equipment-Rental-API/core/models"
	"github/remony/Equipment-Rental-API/core/utils"
)

type Error struct {
	Message string `json:"message"`
}

func generateImageRoutes(api router.API) {
	api.Router.Get("/image/:filename", func(c web.C, res http.ResponseWriter, r *http.Request) {
		result := models.GetImage(api,c.URLParams["filename"])

		data, err := json.Marshal(result)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		res.Write(data)
	})

	api.Router.Get("/images", func (c web.C, res http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")
		res.Header().Set("Content-Type", "application/json")
		if token != "" {
			result := models.GetAllImages(api)

			data, err := json.Marshal(result)
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}

			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(http.StatusOK)
			res.Write(data)
		} else {
//			res.Header().Set("Content-Type", "application/json")
							res.WriteHeader(http.StatusUnauthorized)
//							res.Write(data)
		}

//		res.Header().Set("Content-Type", "application/json")
						res.WriteHeader(http.StatusInternalServerError)
//						res.Write(data)
	})


	api.Router.Post("/image/upload", func (c web.C, res http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("token")

		if token != "" {
			if models.IsSessionValid(api, token) {
				file, header, err:= r.FormFile("image")
				if err != nil {
					panic(err)
				}

				// TODO Check if it already exists
				filename := utils.RandomString(10) + path.Ext(header.Filename)
				if utils.Write(file, filename) {
					models.AddImageLocationToDb(api, filename, header.Filename, header.Filename, token)

					result := models.GetImage(api, filename)

					data, err := json.Marshal(result)
					if err != nil {
						http.Error(res, err.Error(), http.StatusInternalServerError)
						return
					}
					res.Header().Set("Content-Type", "application/json")
					res.WriteHeader(http.StatusOK)
					res.Write(data)
				}
			} else {
//				TODO Return UNAUTHORIZED
			}

		} else {
			// TODO Return UNAUTHORIZED
		}
	})

//	TODO Delete

//	Should we store images raw or scale them (to save storage?)

}
