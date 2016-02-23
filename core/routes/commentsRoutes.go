package routes

import (
	"github.com/zenazn/goji/web"
	"github.com/remony/Equipment-Rental-API/core/router"
	"net/http"
	"encoding/json"
	"github.com/remony/Equipment-Rental-API/core/models"
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

	api.Router.Post("/product/:pid/comment", func(c web.C, res http.ResponseWriter, req *http.Request) {
		pid := c.URLParams["pid"]
		comment := req.Header.Get("comment")
		if len(pid) > 5 {
			if (len(comment) < 140 && len(comment) > 5) {
				if models.AddComment(api, req.Header.Get("token"), pid, comment) {
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