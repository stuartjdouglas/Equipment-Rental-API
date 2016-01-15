package routes
import (
"net/http"
	"encoding/json"
	"github.com/remony/Equipment-Rental-API/core/router"
	"github.com/remony/Equipment-Rental-API/core/models"
	"github.com/remony/Equipment-Rental-API/core/models/sessions"
	"github.com/zenazn/goji/web"
	"github.com/remony/Equipment-Rental-API/core/database"
)

func generateUserRoutes(api router.API) {
	api.Router.Get("/user/:name", func (c web.C, res http.ResponseWriter, r *http.Request) {
		result := models.GetUser(api, c.URLParams["name"])
		data, err := json.Marshal(result)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
		}
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(200)
		res.Write(data)
	})

	api.Router.Post("/user/register", func (c web.C, res http.ResponseWriter, r *http.Request) {
		newdata := register{}
		err := json.NewDecoder(r.Body).Decode(&newdata)
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
		}




		if !models.CheckIfUserExists(api, newdata.Username) {
			//hash := secure.SaltPassword(newdata.Password)
			if database.RegisterUser(api, newdata.Username, newdata.Password, newdata.Email) {
				res.Header().Set("Content-Type", "application/json")
				res.WriteHeader(http.StatusCreated)
				json.NewEncoder(res).Encode(error_response{Message:"User Created"})
			} else {
				res.Header().Set("Content-Type", "application/json")
				res.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(res).Encode(error_response{Message:"User not created: Something went wrong"})
			}
		} else {
			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(http.StatusConflict)
			json.NewEncoder(res).Encode(error_response{Message:"User not created: Already exists"})
		}
	})

	api.Router.Get("/users", func (c web.C, res http.ResponseWriter, r *http.Request) {

		result := models.GetUsers(api)

		data, err := json.Marshal(result)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(200)
		res.Write(data)
	})

	api.Router.Get("/profile", func (c web.C, res http.ResponseWriter, r *http.Request) {

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

	api.Router.Get("/profile/sessions", func (c web.C, res http.ResponseWriter, r *http.Request) {
		if r.Header.Get("token") != "" {
			if (sessions.IsSessionValid(api, r.Header.Get("token"))) {
				result := sessions.GetSessions(api, r.Header.Get("token"))
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
