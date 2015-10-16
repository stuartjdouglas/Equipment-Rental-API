package routes

import (
//	"net/http"
//	"github.com/zenazn/goji/web"
//	"strings"
//	"../controllers"
//	"../config"
//	"../server"
	"../router"
"net/http"
	"encoding/json"
	"fmt"
	"github.com/zenazn/goji/web"
	"../models"
	"strings"
)


/*

	 // Sets the Header
w.Header().Set("Content-Type", "application/json")

// Returns the JSON
encoder := json.NewEncoder(w)
encoder.Encode(message)

 */


type hello struct {
	Message string `json:"message"`
}

type test_struct struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
}

func processUser(username string, password string) test_struct {
	user:= test_struct{}
	if strings.Replace(username, " ", "+", -1) != "" && strings.Replace(password, " ", "+", -1) != "" {
		user.Username = username;
		user.Password = password;
		return user;
	}	else	{
		return user
	}
}

func CreateRoutes (api router.API) {

//	api.Router.Get("/hello/:name", controllers.Hello)
//	api.Router.Get("/person", controllers.Person)

	api.Router.Get("/hello/:name", func (c web.C, res http.ResponseWriter, r *http.Request) {
		message := hello{
			Message: fmt.Sprintf("こんにちは, %s!", c.URLParams["name"]),
		}
		//	encoder := json.NewEncoder(w)

//		data, err := json.Marshal(message)
//		if err != nil {
//			http.Error(res, err.Error(), http.StatusInternalServerError)
//			return
//		}

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

	type error_response struct {
		Message string `json:"message"`
	}





	api.Router.Post("/user/register", func (c web.C, res http.ResponseWriter, r *http.Request) {
		newdata := test_struct{}
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
		data, err := json.Marshal(result)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(200)
		res.Write(data)
//		json.NewEncoder(res).Encode(data)
	})
}