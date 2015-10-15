package router
import (
	"github.com/zenazn/goji/web"
	"../config"
)


type API struct {
	Context config.Context
	Router *web.Mux
}

func New (context config.Context, router *web.Mux) *API{
	return &API {
		context,
		router,
	}
}