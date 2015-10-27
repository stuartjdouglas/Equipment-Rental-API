package router
import (
	"github.com/zenazn/goji/web"

	"../config/database"
)


type API struct {
	Context database.Context
	Router *web.Mux
}

func New (context database.Context, router *web.Mux) *API{
	return &API {
		context,
		router,
	}
}