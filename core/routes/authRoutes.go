package routes

import (
	"encoding/json"
	"net/http"
	"github.com/zenazn/goji/web"
	"github.com/remony/Equipment-Rental-API/core/router"
	"github.com/remony/Equipment-Rental-API/core/database"
	"github.com/remony/Equipment-Rental-API/core/models"
)

type tokenremoved struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

func generateAuthRoutes(api router.API) {

	api.Router.Post("/logout", func(c web.C, res http.ResponseWriter, r *http.Request) {
		//		Not yet implemented
		//		Call method to remove token
	})

	api.Router.Post("/user/register", func(c web.C, res http.ResponseWriter, r *http.Request) {
		newdata := models.Register{}
		err := json.NewDecoder(r.Body).Decode(&newdata)
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
		}

		if models.PerformRegister(api, newdata, false) {
			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(http.StatusCreated)
			json.NewEncoder(res).Encode(error_response{Message:"User Created"})
		} else {
			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(res).Encode(error_response{Message:"User not created: Something went wrong"})
		}
	})

	api.Router.Post("/login", func(c web.C, res http.ResponseWriter, r *http.Request) {
		// Get username and password from form
		var loginDetails = login{
			Username: r.FormValue("username"),
			Password: r.FormValue("password"),
		}
		// If any of the values are empty, check the json body for login and password
		if len(loginDetails.Username) == 0 || len(loginDetails.Password) == 0 {
			err := json.NewDecoder(r.Body).Decode(&loginDetails)
			if err != nil {
				http.Error(res, err.Error(), http.StatusBadRequest)
			}
		}

		var login database.Auth;
		login = models.PerformLogin(api, loginDetails.Username, loginDetails.Password)

		if (login.Success) {
			data, err := json.Marshal(login)
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}
			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(http.StatusOK)
			res.Write(data)
		} else {
			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(res).Encode(error_response{Message:"Invalid Username or Password"})
		}

	})

	api.Router.Delete("/session", func(c web.C, res http.ResponseWriter, r *http.Request) {
		var idenf = r.Header.Get("id")
		var token = r.Header.Get("token")

		if (token != "" && idenf != "") {
			if (models.IsSessionValid(api, token)) {

				removal := models.DisableToken(api, idenf)

				if (removal) {

					data, err := json.Marshal(tokenremoved{ID:idenf, Message:"Session removed."})
					if err != nil {
						http.Error(res, err.Error(), http.StatusInternalServerError)
						return
					}
					res.Header().Set("Content-Type", "application/json")
					res.WriteHeader(http.StatusOK)
					res.Write(data)
				} else {
					res.Header().Set("Content-Type", "application/json")
					res.WriteHeader(http.StatusInternalServerError)
					json.NewEncoder(res).Encode(error_response{Message:"Server error"})
				}
			} else {
				res.Header().Set("Content-Type", "application/json")
				res.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(res).Encode(error_response{Message:"Invalid Username or Password"})
			}
		}else {
			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(res).Encode(error_response{Message:"Missing parameters id and/or token"})
		}

	})
}