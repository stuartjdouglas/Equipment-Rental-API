package routes

import (
	"github.com/zenazn/goji/web"
	"github.com/remony/Equipment-Rental-API/core/router"
	"net/http"
	"encoding/json"
	"github.com/remony/Equipment-Rental-API/core/models"
	"strconv"
)

func generateRequestRouter(api router.API) {
	api.Router.Get("/owner/requests", func(c web.C, res http.ResponseWriter, req *http.Request) {
		token := req.Header.Get("token")
		start := req.Header.Get("start")
		count := req.Header.Get("count")

		if len(token) > 3 {
			result := models.GetProductsRequests(api, token, start, count)
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
				Message: "Unable to request item",
			}
			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(res).Encode(message)
		}
	})

	api.Router.Get("/requests", func(c web.C, res http.ResponseWriter, req *http.Request) {
		//pid := c.URLParams["pid"]
		token := req.Header.Get("token")
		start := req.Header.Get("start")
		count := req.Header.Get("count")

		if len(token) > 3 {
			result := models.GetUserRequests(api, token, start, count)
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
				Message: "Unable to request item",
			}
			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(res).Encode(message)
		}
	})

	api.Router.Get("/requests/:pid", func(c web.C, res http.ResponseWriter, req *http.Request) {
		pid := c.URLParams["pid"]
		token := req.Header.Get("token")

		if len(token) > 3 {
			result := models.GetProductRequests(api, pid, token)
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
				Message: "Unable to request item",
			}
			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(res).Encode(message)
		}
	})

	api.Router.Post("/product/:pid/request/authorize/:username", func(c web.C, res http.ResponseWriter, req *http.Request) {
		pid := c.URLParams["pid"]
		username := c.URLParams["username"]
		token := req.Header.Get("token")
		if len(pid) > 3 && len(username) > 3 {
			result := models.RequestToRent(api, pid, username, token)
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
				Message: "Unable to request item",
			}
			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(res).Encode(message)
		}
	})

	api.Router.Get("/product/:pid/request", func(c web.C, res http.ResponseWriter, req *http.Request) {
		pid := c.URLParams["pid"]
		if len(pid) > 3 {
			result := models.RequestProductStatus(api, pid, req.Header.Get("token"))
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
				Message: "Unable to request item",
			}
			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(res).Encode(message)
		}
	})

	api.Router.Post("/product/:pid/request", func(c web.C, res http.ResponseWriter, req *http.Request) {
		pid := c.URLParams["pid"]

		if len(pid) > 3 {
			result := models.RequestProduct(api, pid, req.Header.Get("token"))
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
				Message: "Unable to request item",
			}
			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(res).Encode(message)
		}
	})

	api.Router.Post("/product/:pid/request/cancel", func(c web.C, res http.ResponseWriter, req *http.Request) {
		pid := c.URLParams["pid"]
		if len(pid) > 3 {
			result := models.CancelRequestProduct(api, pid, req.Header.Get("token"), req.Header.Get("username"))
			var message hello
			if (result) {
				res.WriteHeader(200)
				message.Message = "Request deleted"
			} else {
				res.WriteHeader(http.StatusInternalServerError)
				message.Message = "Unable to delete request"
			}
			data, err := json.Marshal(message)
			if err != nil {
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}

			res.Header().Set("Content-Type", "application/json")

			res.Write(data)
		} else {
			message := hello{
				Message: "No Tags exist",
			}
			res.Header().Set("Content-Type", "application/json")
			res.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(res).Encode(message)
		}
	})

	api.Router.Get("/admin/requests/products", func(c web.C, res http.ResponseWriter, req *http.Request) {
		token := req.Header.Get("token")
		step, err := strconv.Atoi(req.Header.Get("step"))
		count, err := strconv.Atoi(req.Header.Get("count"))

		result := models.GetAdminProductsRequests(api, token, step, count)

		data, err := json.Marshal(result)

		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(200)
		res.Write(data)

	})

	api.Router.Post("/admin/requests/listing/:pid/authorize", func(c web.C, res http.ResponseWriter, req *http.Request) {
		token := req.Header.Get("token")
		pid := c.URLParams["pid"]

		result := models.AdminAuthorizeListingRequest(api, pid, token)

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
