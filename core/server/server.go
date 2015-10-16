package server

import (
	"fmt"
	"strconv"
	"../routes"
	"github.com/zenazn/goji/graceful"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
	"../config"
	"../router"
)

func Start(settings config.Config, context config.Context) {
	fmt.Println("こんにちは, listening on port :" + strconv.Itoa(settings.Port))

	masterRouter := web.New()

	//Get routes
	apiRouter :=  web.New()
	renderRouter := web.New()
	masterRouter.Handle("/api/*", apiRouter)
	masterRouter.Handle("/client/*", renderRouter)

	apiRouter.Use(middleware.SubRouter)
	renderRouter.Use(middleware.SubRouter)

	routes.CreateRoutes(router.API{Router:apiRouter, Context:context})

	routes.CreateRenderRoutes(router.API{Router:renderRouter, Context:context})

	// Gracefully Serve
	err := graceful.ListenAndServe(":" + strconv.Itoa(settings.Port), masterRouter)
	if err != nil {
		panic(err)
	}
}

