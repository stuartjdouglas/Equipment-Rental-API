package routes

import (
	"github.com/zenazn/goji/web"
	"net/http"
	"encoding/json"
	"strings"
	"fmt"
	"github.com/remony/Equipment-Rental-API/core/router"
	"github.com/remony/Equipment-Rental-API/core/models"
)

func generateSlug(slug string) string {
	return strings.Replace(strings.ToLower(slug), " ", "_", -1)
}

func createPostRoutes(api router.API) {
	api.Router.Post("/post", func(c web.C, res http.ResponseWriter, r *http.Request) {
		if r.Header.Get("token") != "" {
			post := models.Post{}
			err := json.NewDecoder(r.Body).Decode(&post)
			if err != nil {
				http.Error(res, err.Error(), http.StatusBadRequest)
			}

			fmt.Println(post.Title + " " + post.Content)
			if len(post.Content) > 140 {
				res.Header().Set("Content-Type", "application/json")
				res.WriteHeader(http.StatusNotAcceptable)
				json.NewEncoder(res).Encode(error_response{Message:"Post not created: Content too long"})
			}

			if len(post.Title) > 140 {
				res.Header().Set("Content-Type", "application/json")
				res.WriteHeader(http.StatusNotAcceptable)
				json.NewEncoder(res).Encode(error_response{Message:"Post not created: Title too long"})
			}
			slug := generateSlug(post.Title)

			if !models.CheckIfPostExists(api, slug) {
				if models.CreatePost(api, post, r.Header.Get("token"), slug) {
					res.Header().Set("Content-Type", "application/json")
					res.WriteHeader(http.StatusCreated)
					json.NewEncoder(res).Encode(error_response{Message:"Post Created"})
				} else {
					res.Header().Set("Content-Type", "application/json")
					res.WriteHeader(http.StatusInternalServerError)
					json.NewEncoder(res).Encode(error_response{Message:"User not created: Something went wrong"})
				}
			} else {
				res.Header().Set("Content-Type", "application/json")
				res.WriteHeader(http.StatusConflict)
				json.NewEncoder(res).Encode(error_response{Message:"Post not created: Already exists"})
			}
		} else {
			http.Error(res, "", http.StatusUnauthorized)
		}
	})



	//	Get all posts
	api.Router.Get("/posts", func(c web.C, res http.ResponseWriter, r *http.Request) {

		result := models.GetPosts(api)

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

	api.Router.Get("/posts/:username", func(c web.C, res http.ResponseWriter, r *http.Request) {

		result := models.GetPostsFromUser(api, c.URLParams["username"])

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

	api.Router.Get("/post/:slug", func(c web.C, res http.ResponseWriter, r *http.Request) {

		result := models.GetPost(api, c.URLParams["slug"])

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
}