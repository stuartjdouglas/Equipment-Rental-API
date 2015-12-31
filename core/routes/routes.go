package routes

import (
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/zenazn/goji/web"
	"github.com/remony/Equipment-Rental-API/core/models"
	"github.com/remony/Equipment-Rental-API/core/router"
)


type login struct {
	Username string `json:"username" param:"username"`
	Password string `json:"password" param:"password"`
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
	generateUserRoutes(api)
	generateAuthRoutes(api)
	generateImageRoutes(api)
	generateQrRoutes(api)
	generateProductRoutes(api)
	// A test route

	api.Router.Get("/hello/:name", func (c web.C, res http.ResponseWriter, r *http.Request) {
		message := hello{
			Message: fmt.Sprintf("こんにちは, %s!", c.URLParams["name"]),
		}
//		email.SendEmail(api, "remonasebi@gmail.com", "hello", "<h1>" + message.Message + "</h1>")
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(200)
		json.NewEncoder(res).Encode(message)
	})

	api.Router.Get("/hello", func (c web.C, res http.ResponseWriter, r *http.Request) {
		if r.Header.Get("token") != "" {
			result := models.GetHello(api, r.Header.Get("token"))
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


}