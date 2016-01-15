package routes

import (
	"github.com/remony/Equipment-Rental-API/core/router"
)

type Error struct {
	Message string `json:"message"`
}

func generateImageRoutes(api router.API) {
//	api.Router.Get("/image/:filename", func(c web.C, res http.ResponseWriter, r *http.Request) {
//		result := models.GetImage(api,c.URLParams["filename"])
//
//		data, err := json.Marshal(result)
//		if err != nil {
//			http.Error(res, err.Error(), http.StatusInternalServerError)
//			return
//		}
//
//		res.Header().Set("Content-Type", "application/json")
//		res.WriteHeader(http.StatusOK)
//		res.Write(data)
//	})
//
//	api.Router.Get("/images", func (c web.C, res http.ResponseWriter, r *http.Request) {
//		token := r.Header.Get("token")
//		res.Header().Set("Content-Type", "application/json")
//		if token != "" {
//			result := models.GetAllImages(api)
//
//			data, err := json.Marshal(result)
//			if err != nil {
//				http.Error(res, err.Error(), http.StatusInternalServerError)
//				return
//			}
//
//			res.Header().Set("Content-Type", "application/json")
//			res.WriteHeader(http.StatusOK)
//			res.Write(data)
//		} else {
//			res.WriteHeader(http.StatusUnauthorized)
//		}
//
//	})
//
//
//	api.Router.Post("/image/upload", func (c web.C, res http.ResponseWriter, r *http.Request) {
//		token := r.Header.Get("token")
//
//		if token != "" {
//			if sessions.IsSessionValid(api, token) {
//				file, header, err:= r.FormFile("image")
//				if err != nil {
//					panic(err)
//				}
//
//				imagecode := utils.RandomString(20)
//
//				for !models.DoesImageExist(api, imagecode) {
//					log.Println("creating new code")
//					imagecode = utils.RandomString(20)
//				}
//
//				filename := imagecode + path.Ext(header.Filename)
//
//
//
//				if utils.WriteImage(file, header.Header, imagecode, path.Ext(header.Filename)) {
//					models.AddImageLocationToDb(api, filename, header.Filename, header.Filename, token)
//
//					result := models.GetImage(api, filename)
//
//					data, err := json.Marshal(result)
//					if err != nil {
//						http.Error(res, err.Error(), http.StatusInternalServerError)
//						return
//					}
//					res.Header().Set("Content-Type", "application/json")
//					res.WriteHeader(http.StatusOK)
//					res.Write(data)
//				}
//			} else {
////				TODO Return UNAUTHORIZED
//			}
//
//		} else {
//			// TODO Return UNAUTHORIZED
//		}
//	})
//
////	TODO Delete
//
////	Should we store images raw or scale them (to save storage?)

}
