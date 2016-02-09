package models

import (
	"github.com/remony/Equipment-Rental-API/core/router"
	"github.com/remony/Equipment-Rental-API/core/database"
	"log"
)

func UpdateIndex(api router.API, title string, Description string, token string) bool {
	log.Println(title)
	//if isAdmin(api, token) {
		return database.UpdateIndex(api, title, Description, token)
	//}
	return false
}
