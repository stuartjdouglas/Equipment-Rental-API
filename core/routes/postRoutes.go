package routes

import (
	"../router"
	"github.com/zenazn/goji/web"
"net/http"
	"fmt"
	"encoding/json"
)


func createPostRoutes (api router.API) {
	api.Router.Post("/post", func (c web.C, res http.ResponseWriter, r *http.Request) {
		message := hello{
			Message: fmt.Sprintf("こんにちは, %s!", c.URLParams["name"]),
		}
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(200)
		json.NewEncoder(res).Encode(message)
	})
}