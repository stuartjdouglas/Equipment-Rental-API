package routes

import (
	"github.com/zenazn/goji/web"
	"net/http"
	"github.com/remony/Equipment-Rental-API/core/router"
	"github.com/remony/Equipment-Rental-API/core/models"
	"encoding/json"
	"log"
)

func generatePushNotificationRoutes(api router.API) {
	api.Router.Post("/notification/register", func(c web.C, res http.ResponseWriter, req *http.Request) {

		result := models.PushNotificationAddRegID(api, req.Header.Get("regid"), req.Header.Get("type"), req.Header.Get("token"))
		data, err := json.Marshal(result)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(200)
		res.Write(data)
	})

	api.Router.Post("/notification/product/:pid", func(c web.C, res http.ResponseWriter, req *http.Request) {
		decoder := json.NewDecoder(req.Body)
		var message models.Notification
		err := decoder.Decode(&message)
		if err != nil {
			log.Println(err)
		}
		result := models.SendNotificationProduct(api, req.Header.Get("token"), c.URLParams["pid"], message)
		data, err := json.Marshal(result)
		if err != nil {
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(200)
		res.Write(data)
	})

	api.Router.Post("/notification", func(c web.C, res http.ResponseWriter, req *http.Request) {
		log.Println("testing")
		models.TestNotification();
	})
}
