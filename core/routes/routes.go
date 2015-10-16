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

	api.Router.Get("/user/register", func (c web.C, res http.ResponseWriter, r *http.Request) {
		r.Body.Read()
	})

	api.Router.Get("/users", func (c web.C, res http.ResponseWriter, r *http.Request) {

		result := models.GetUsers(api)
		data, err := json.Marshal(result)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

//		fmt.Printf(data)
		
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(200)
		res.Write(data)
//		json.NewEncoder(res).Encode(data)
	})
}