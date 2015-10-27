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
	"github.com/rs/cors"
)

func Start(settings config.Properties, context config.Context) {
	fmt.Println("こんにちは, listening on port :" + strconv.Itoa(settings.Port))

	masterRouter := web.New()


	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE"},
		AllowCredentials: true,
		AllowedHeaders: []string{"*"},
	})


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

//	Apply the CORS options to the main route handler
	masterRouter.Use(c.Handler)


	apiRouter.Use(middleware.SubRouter)
	angularRouter.Use(middleware.SubRouter)
	renderRouter.Use(middleware.SubRouter)

	routes.CreateRoutes(router.API{Router:apiRouter, Context:context})

	routes.CreateRenderRoutes(router.API{Router:renderRouter, Context:context})

	// Gracefully Serve
	if portIsFree(settings.Port) {
		err := graceful.ListenAndServe(":" + strconv.Itoa(settings.Port), masterRouter)
		if err != nil {
			//		If an error occurs, normally is port is already in use

			//		Don't panic
			panic(err)
		}
	}
}

// Checks if a port is free
func portIsFree(port int) bool {
//	If the port is being used

//	Return false

//	if not in use
	return true
}


