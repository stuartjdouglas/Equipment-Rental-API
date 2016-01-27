package routes

import (
	"net/http"
	"encoding/json"
	"github.com/remony/Equipment-Rental-API/core/router"
	"github.com/remony/Equipment-Rental-API/core/models"
	"github.com/zenazn/goji/web"
)

func generateUserRoutes(api router.API) {
	api.Router.Get("/user/:name", func(c web.C, res http.ResponseWriter, r *http.Request) {
		result := models.GetUserData(api, c.URLParams["name"])
		data, err := json.Marshal(result)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(200)
		res.Write(data)
	})

	api.Router.Get("/users", func(c web.C, res http.ResponseWriter, r *http.Request) {

		result := models.GetUsersData(api)

		data, err := json.Marshal(result)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(200)
		res.Write(data)
	})

	api.Router.Get("/profile", func(c web.C, res http.ResponseWriter, r *http.Request) {

		if r.Header.Get("token") != "" {
			result := models.GetProfile(api, r.Header.Get("token"))
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

	api.Router.Get("/profile/sessions", func(c web.C, res http.ResponseWriter, r *http.Request) {
		if r.Header.Get("token") != "" {
			if (models.IsSessionValid(api, r.Header.Get("token"))) {
				result := models.GetSessions(api, r.Header.Get("token"))
				data, err := json.Marshal(result)
				if err != nil {
					http.Error(res, err.Error(), http.StatusInternalServerError)
					return
				}

				res.Header().Set("Content-Type", "application/json")
				res.WriteHeader(200)
				res.Write(data)
			} else {
				http.Error(res, "Unauthorized", http.StatusUnauthorized)
			}
		} else {
			http.Error(res, "", http.StatusUnauthorized)
		}
	})

	//	api.Router.Get("/profile/session", func (c web.C, res http.ResponseWriter, r *http.Request) {
	//		if r.Header.Get("token") != "" {
	//			if (sessions.IsSessionValid(api, r.Header.Get("token"))) {
	//				result := sessions.GetSession(api, r.Header.Get("token"))
	//				data, err := json.Marshal(result)
	//				if err != nil {
	//					http.Error(res, err.Error(), http.StatusInternalServerError)
	//					return
	//				}
	//
	//				res.Header().Set("Content-Type", "application/json")
	//				res.WriteHeader(200)
	//				res.Write(data)
	//			} else {
	//				http.Error(res, "Unauthorized", http.StatusUnauthorized)
	//			}
	//
	//		} else {
	//			http.Error(res, "", http.StatusUnauthorized)
	//		}
	//	})
}
