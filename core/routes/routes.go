package routes

import (
	"../router"
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/zenazn/goji/web"
	"../models"
	"net"
)


type login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type hello struct {
	Message string `json:"message"`
}

type register struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
}

type error_response struct {
	Message string `json:"message"`
}


func CreateRoutes (api router.API) {
	createPostRoutes(api)
	// A test route
	api.Router.Get("/hello/:name", func (c web.C, res http.ResponseWriter, r *http.Request) {
		message := hello{
			Message: fmt.Sprintf("こんにちは, %s!", c.URLParams["name"]),
		}
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(200)
		json.NewEncoder(res).Encode(message)
	})


	api.Router.Get("/user/:name", func (c web.C, res http.ResponseWriter, r *http.Request) {
		result := models.GetUser(api, c.URLParams["name"])
		data, err := json.Marshal(result)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
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
			if models.RegisterUser(api, newdata.Username, newdata.Password, newdata.Email) {
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

		fmt.Println(result)
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
			http.Error(res, "", http.StatusUnauthorized)
		}
	})

	api.Router.Get("/profile/session", func (c web.C, res http.ResponseWriter, r *http.Request) {
		if r.Header.Get("token") != "" {
			result := models.GetSession(api, r.Header.Get("token"))
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

	api.Router.Post("/logout", func (c web.C, res http.ResponseWriter, r *http.Request) {
//		Not yet implemented
//		Call method to remove token
	})

	api.Router.Post("/login", func (c web.C, res http.ResponseWriter, r *http.Request) {
		loginDetails := login{}
		err := json.NewDecoder(r.Body).Decode(&loginDetails)
		if err != nil {
			http.Error(res, err.Error(), http.StatusBadRequest)
		}
		ip_address, _, _ := net.SplitHostPort(r.RemoteAddr)
		result := models.LoginUser(api, loginDetails.Username, loginDetails.Password, ip_address)
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