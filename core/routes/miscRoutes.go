package routes

import (
	"github.com/zenazn/goji/web"
	"github.com/remony/Equipment-Rental-API/core/router"
	"net/http"
	"encoding/json"
	"github.com/remony/Equipment-Rental-API/core/models"
	"fmt"
	"github.com/remony/Equipment-Rental-API/core/database"
	"strconv"
)

func generateMiscRoutes(api router.API) {

	api.Router.Get("/hello/:name", func(c web.C, res http.ResponseWriter, r *http.Request) {
		message := hello{
			Message: fmt.Sprintf("こんにちは, %s!", c.URLParams["name"]),
		}
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(200)
		json.NewEncoder(res).Encode(message)
	})

	api.Router.Get("/hello", func(c web.C, res http.ResponseWriter, r *http.Request) {
		if r.Header.Get("token") != "" {
			result := database.GetHello(api, r.Header.Get("token"))
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

	api.Router.Post("/product/:pid/comment", func(c web.C, res http.ResponseWriter, req *http.Request) {
		pid := c.URLParams["pid"]
		comment := req.Header.Get("comment")
		rating, err := strconv.Atoi(req.Header.Get("rating"))
		if (err != nil) {
			rating = 3
		}

		if len(pid) > 5 {
			if (len(comment) < 140 && len(comment) > 5) {
				if models.AddComment(api, req.Header.Get("token"), pid, comment, rating).Text != "null" {
					message := hello{
						Message: "Comment added",
					}
					res.Header().Set("Content-Type", "application/json")
					res.WriteHeader(http.StatusOK)
					json.NewEncoder(res).Encode(message)
				} else {
					message := hello{
						Message: "Something went wrong on our end",
					}
					res.Header().Set("Content-Type", "application/json")
					res.WriteHeader(http.StatusInternalServerError)
					json.NewEncoder(res).Encode(message)
				}

			} else {
				message := hello{
					Message: "Comment too big, must be less than 140.",
				}
				res.Header().Set("Content-Type", "application/json")
				res.WriteHeader(http.StatusRequestEntityTooLarge)
				json.NewEncoder(res).Encode(message)
			}

		} else {
			message := hello{
				Message: "Unable to add comment",
			}
			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(res).Encode(message)
		}
	})
	api.Router.Post("/product/:pid/like", func(c web.C, res http.ResponseWriter, req *http.Request) {
		pid := c.URLParams["pid"]
		token := req.Header.Get("token")
		if len(pid) > 5 &&  len(token) > 5 {
			if models.Like(api, pid, token) {
				message := hello{
					Message: "liked",
				}
				res.Header().Set("Content-Type", "application/json")
				res.WriteHeader(http.StatusOK)
				json.NewEncoder(res).Encode(message)
			} else {
				message := hello{
					Message: "Unable to like",
				}
				res.Header().Set("Content-Type", "application/json")
				res.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(res).Encode(message)
			}

		} else {
			message := hello{
				Message: "Unable to like",
			}
			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(res).Encode(message)
		}
	})
	api.Router.Post("/product/:pid/unlike", func(c web.C, res http.ResponseWriter, req *http.Request) {
		pid := c.URLParams["pid"]
		token := req.Header.Get("token")
		if len(pid) > 5 &&  len(token) > 5 {
			if models.UnLike(api, pid, token) {
				message := hello{
					Message: "Unliked",
				}
				res.Header().Set("Content-Type", "application/json")
				res.WriteHeader(http.StatusOK)
				json.NewEncoder(res).Encode(message)
			} else {
				message := hello{
					Message: "Unable to unlike",
				}
				res.Header().Set("Content-Type", "application/json")
				res.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(res).Encode(message)
			}

		} else {
			message := hello{
				Message: "Unable to unlike",
			}
			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(res).Encode(message)
		}
	})
}