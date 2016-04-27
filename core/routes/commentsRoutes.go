package routes

import (
	"github.com/zenazn/goji/web"
	"github.com/remony/Equipment-Rental-API/core/router"
	"net/http"
	"encoding/json"
	"github.com/remony/Equipment-Rental-API/core/models"
	"strconv"
)

func generateCommentsRoutes(api router.API) {

	api.Router.Post("/product/:pid/comments/enable", func(c web.C, res http.ResponseWriter, req *http.Request) {
		pid := c.URLParams["pid"]
		token := req.Header.Get("token")
		if len(pid) > 5 {
			if (models.IsOwner(api, token, pid)) {
				models.EnableComments(api, pid, token)
				message := hello{
					Message: "Comments enabled",
				}
				res.Header().Set("Content-Type", "application/json")
				res.WriteHeader(http.StatusOK)
				json.NewEncoder(res).Encode(message)

			} else {
				message := hello{
					Message: "Not authorized",
				}
				res.Header().Set("Content-Type", "application/json")
				res.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(res).Encode(message)
			}

		} else {
			message := hello{
				Message: "Unable to enable comments",
			}
			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(res).Encode(message)
		}
	})
	api.Router.Post("/product/:pid/comments/disable", func(c web.C, res http.ResponseWriter, req *http.Request) {
		pid := c.URLParams["pid"]
		token := req.Header.Get("token")
		if len(pid) > 5 {
			if (models.IsOwner(api, token, pid)) {
				models.DisableComments(api, pid, token)
				message := hello{
					Message: "Comments disabled",
				}
				res.Header().Set("Content-Type", "application/json")
				res.WriteHeader(http.StatusOK)
				json.NewEncoder(res).Encode(message)

			} else {
				message := hello{
					Message: "Not authorized",
				}
				res.Header().Set("Content-Type", "application/json")
				res.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(res).Encode(message)
			}

		} else {
			message := hello{
				Message: "Unable to disable comments",
			}
			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(res).Encode(message)
		}
	})
	api.Router.Post("/product/:pid/comment/:cid/approve", func(c web.C, res http.ResponseWriter, req *http.Request) {
		pid := c.URLParams["pid"]
		cid := c.URLParams["cid"]
		token := req.Header.Get("token")
		if len(pid) > 5 {
			if (models.IsOwner(api, token, pid)) {
				models.ApproveComment(api, pid, cid)
				message := hello{
					Message: "Comments approved",
				}
				res.Header().Set("Content-Type", "application/json")
				res.WriteHeader(http.StatusOK)
				json.NewEncoder(res).Encode(message)

			} else {
				message := hello{
					Message: "Not authorized",
				}
				res.Header().Set("Content-Type", "application/json")
				res.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(res).Encode(message)
			}

		} else {
			message := hello{
				Message: "Unable to approve comment",
			}
			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(res).Encode(message)
		}
	})

	api.Router.Post("/product/:pid/comment/:cid/edit", func(c web.C, res http.ResponseWriter, req *http.Request) {
		//token string, cid string, comment string, rating int
		cid := c.URLParams["cid"]
		comment := req.Header.Get("comment")
		token := req.Header.Get("token")
		rating, err := strconv.Atoi(req.Header.Get("rating"))
		if (err != nil) {
			rating = 3
		}

		if len(cid) > 5 {
			result := models.EditComment(api, token, cid, comment, rating)
			data, err := json.Marshal(result)
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}

			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(200)
			res.Write(data)
		} else {
			message := hello{
				Message: "Unable to edit comment",
			}
			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(res).Encode(message)
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
				newComment := models.AddComment(api, req.Header.Get("token"), pid, comment, rating)
				if newComment.Text != "null" {
					data, err := json.Marshal(newComment)
					if err != nil {
						http.Error(res, err.Error(), http.StatusInternalServerError)
						return
					}

					res.Header().Set("Content-Type", "application/json")
					res.WriteHeader(200)
					res.Write(data)
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
	api.Router.Delete("/product/:pid/comment/:cid", func(c web.C, res http.ResponseWriter, req *http.Request) {
		pid := c.URLParams["pid"]
		cid := c.URLParams["cid"]
		token := req.Header.Get("token")
		if len(pid) > 5 && len(cid) > 5 && len(token) > 5 {
			models.DeleteComment(api, pid, cid, token)
			message := hello{
				Message: "comment deleted",
			}
			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(http.StatusOK)
			json.NewEncoder(res).Encode(message)
		} else {
			message := hello{
				Message: "Unable to delete comment",
			}
			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(res).Encode(message)
		}
	})
}