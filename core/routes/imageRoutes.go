package routes

import (
	"../router"
	"github.com/zenazn/goji/web"
"net/http"
	"../models"
	"../utils"
	"path"
	"encoding/json"
)


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


				if utils.Write(file, utils.GenerateToken(header.Filename) + path.Ext(header.Filename)) {
//					models.AddImageLocationToDb(api, utils.GenerateToken(header.Filename) + path.Ext(header.Filename), header.Filename)
				} else {

				}

//				res.Header().Set("Content-Type", "application/json")
//				res.WriteHeader(200)
//				res.Write(data)
			} else {
				http.Error(res, "", http.StatusUnauthorized)
			}

		} else {
			http.Error(res, "", http.StatusUnauthorized)
		}
	})

//	TODO Delete

//	Should we store images raw or scale them (to save storage?)

}
