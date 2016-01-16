package models

import (
	"github.com/remony/Equipment-Rental-API/core/router"
	"github.com/remony/Equipment-Rental-API/core/database"
)

func IsImageAvailable(api router.API, url string) bool {
	return database.DoesImageExist(api, url)
}