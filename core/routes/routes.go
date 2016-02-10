package routes

import (
	"github.com/remony/Equipment-Rental-API/core/router"
)

type login struct {
	Username string `json:"username" param:"username"`
	Password string `json:"password" param:"password"`
}

type hello struct {
	Message string `json:"message"`
}

type error_response struct {
	Message string `json:"message"`
}

func CreateRoutes(api router.API) {
	createPostRoutes(api)
	generateRootRoutes(api)
	generateUserRoutes(api)
	generateAuthRoutes(api)
	generateImageRoutes(api)
	generateQrRoutes(api)
	generateProductRoutes(api)
	generateTagRoutes(api)
	generateSearchRoutes(api)
	generateRequestRouter(api)
	generatePushNotificationRoutes(api)
	generateAdminRoutes(api)
	generateCommentsRoutes(api)
	generateMiscRoutes(api)
}