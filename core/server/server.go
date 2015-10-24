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
	"github.com/hypebeast/gojistatic"
)

func Start(settings config.Config, context config.Context) {
	fmt.Println("こんにちは, listening on port :" + strconv.Itoa(settings.Port))

	masterRouter := web.New()


	//Get routes
	apiRouter :=  web.New()
	renderRouter := web.New()
	angularRouter := web.New()

	masterRouter.Handle("/api/*", apiRouter)
	masterRouter.Handle("/client/*", renderRouter)
	masterRouter.Handle("/*", angularRouter)

//	angularRouter.Get("/*", http.FileServer(http.Dir("client/app")))

	var options gojistatic.StaticOptions

	angularRouter.Use(gojistatic.Static("client/app/", gojistatic.StaticOptions{
		SkipLogging: true,
		Expires: nil,
	}))
	angularRouter.Use(gojistatic.Static("/*", options))



	apiRouter.Use(middleware.SubRouter)
	angularRouter.Use(middleware.SubRouter)
	renderRouter.Use(middleware.SubRouter)

	routes.CreateRoutes(router.API{Router:apiRouter, Context:context})

	routes.CreateRenderRoutes(router.API{Router:renderRouter, Context:context})

	// Gracefully Serve
	err := graceful.ListenAndServe(":" + strconv.Itoa(settings.Port), masterRouter)
	if err != nil {
		panic(err)
	}
}

