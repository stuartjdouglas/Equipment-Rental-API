package router
import (
	"github.com/zenazn/goji/web"

	"github.com/remony/Equipment-Rental-API/core/config/database"
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