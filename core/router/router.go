package router

import (
	"github.com/zenazn/goji/web"
	"github.com/remony/Equipment-Rental-API/core/config"
)

// API contains the context and router
type API struct {
	Context config.Context
	Router  *web.Mux
	Config  config.Config
}

// New creates a new API
func New(context config.Context, router *web.Mux, config config.Config) *API {
	return &API{
		context,
		router,
		config,
	}
}