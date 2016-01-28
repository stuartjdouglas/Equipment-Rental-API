package models

import (
	"github.com/remony/Equipment-Rental-API/core/router"
	"github.com/remony/Equipment-Rental-API/core/database"
)

func SearchTag(api router.API, tag string, start string, count string) database.Items {
	return database.SearchTag(api, tag, parseStringToInt(start), parseStringToInt(count))
}