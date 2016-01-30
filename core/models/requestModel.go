package models

import (
	"github.com/remony/Equipment-Rental-API/core/router"
	"github.com/remony/Equipment-Rental-API/core/database"
	"log"
)



func RequestProduct(api router.API, pid string, token string) bool {
	if (ValidToken(token)) {
		if database.GetAvailability(api, pid).Available {
			log.Println("product " + pid + " is valid")
			database.SendRequestProduct(api, pid, token)
		} else {
			return false
		}
	} else {
		return false
	}
	return false
}

func CancelRequestProduct(api router.API, pid string, token string) bool {
	if ValidToken(token) {
		if database.GetAvailability(api, pid).Available {
			database.SendCancelRequestProduct(api, pid, token)
		} else {
			return false
		}
	} else {
		return false
	}
	return false
}
