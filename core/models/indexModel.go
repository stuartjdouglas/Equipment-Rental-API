package models

import (
	"github.com/remony/Equipment-Rental-API/core/router"
	"github.com/remony/Equipment-Rental-API/core/database"
)

func GetIndex(api router.API) database.Index {
	return database.GetIndexFromDB(api)
}

