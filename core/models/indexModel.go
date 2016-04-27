package models

import (
	"github.com/remony/Equipment-Rental-API/core/router"
	"github.com/remony/Equipment-Rental-API/core/database"
)

// Returns Index data
func GetIndex(api router.API) database.Index {
	return database.GetIndexFromDB(api)
}

