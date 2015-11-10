package router
import (
	"github.com/zenazn/goji/web"

	"../config/database"
)

// API contains the context and router
type API struct {
	Context database.Context
	Router *web.Mux
}

// New creates a new API
func New (context database.Context, router *web.Mux) *API{
	return &API {
		context,
		router,
	}
}