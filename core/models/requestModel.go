package models

import (
	"github.com/remony/Equipment-Rental-API/core/router"
	"github.com/remony/Equipment-Rental-API/core/database"
	"time"
)

func GetProductsRequests(api router.API, token string, start string, count string) database.ItemRequests {
	if ValidToken(token) {
		return database.GetProductsRequests(api, token, parseStringToInt(start), parseStringToInt(count))
	} else {
		return database.ItemRequests{}
	}
}

func RequestProductStatus(api router.API, pid string, token string) database.Request {
	if ValidToken(token) {
		return database.GetProductRequestStatus(api, pid, token)
	}
	return database.Request{
		Title: "null",
		Date_requested: time.Now(),
		Requested: false,
	}
}


func RequestToRent(api router.API, pid string, username string, token string) bool {
	if IsOwner(api, token, pid) {
		if database.GetAuthedAvailability(api, pid, token).Available {
			return database.RequestToRent(api, pid, username)
		} else {
			return false
		}

	} else {
		return false
	}
	return false
}

func GetProductRequests(api router.API, pid string, token string) database.UserRequests {
	error := database.UserRequest{
		Date_requested: time.Now(),
		Username: "null",
		Gravatar: "",
	}

	var errorresp []database.UserRequest
	errorresp = append(errorresp, error)
	if ValidToken(token) {
		if IsOwner(api, token, pid) {
			return database.GetProductRequests(api, pid, token)
		}
		return database.UserRequests{Total:0, Requests: errorresp}
	} else {
		return database.UserRequests{Total:0, Requests: errorresp}
	}
}

func RequestProduct(api router.API, pid string, token string) database.Request {
	if (ValidToken(token)) {
		if database.GetAvailability(api, pid).Available {
			return database.SendRequestProduct(api, pid, token)
		} else {
			return database.Request{
				Requested: false,
				Title: "null",
				Date_requested: time.Now(),
			}
		}
	} else {
		return database.Request{
			Requested: false,
			Title: "null",
			Date_requested: time.Now(),
		}
	}
	return database.Request{
		Requested: false,
		Title: "null",
		Date_requested: time.Now(),
	}
}

func CancelRequestProduct(api router.API, pid string, token string) bool {
	if ValidToken(token) {
		if database.GetAvailability(api, pid).Available {
			database.SendCancelRequestProduct(api, pid, token)
			return true
		} else {
			return false
		}
	} else {
		return false
	}
	return false
}
