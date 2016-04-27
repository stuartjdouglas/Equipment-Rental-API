package models

import (
	"github.com/remony/Equipment-Rental-API/core/router"
	"github.com/remony/Equipment-Rental-API/core/database"
	"time"
	"github.com/remony/Equipment-Rental-API/core/utils/email"
)

func GetProductsRequests(api router.API, token string, start string, count string) database.ItemRequests {
	if ValidToken(token) {
		return database.GetProductsRequests(api, token, parseStringToInt(start), parseStringToInt(count))
	} else {
		return database.ItemRequests{}
	}
}

func GetUserRequests(api router.API, token string, start string, count string) database.UserItemRequests {
	if ValidToken(token) {
		return database.GetUserRequests(api, token, parseStringToInt(start), parseStringToInt(count))
	} else {
		return database.UserItemRequests{}
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
			result := database.RequestToRent(api, pid, username)
			if result {
				productdata := GetProductFromID(api, pid, token)
				SendNotificationToUser(api, username, Notification{Title: "Your request for " + productdata.Items[0].Product_name + " has been accepted", Message: productdata.Items[0].Product_name + " is now ready to collect"})
			}
			return result
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
			result := database.SendRequestProduct(api, pid, token)
			username := database.GetUserNameFromToken(api, token)
			user := database.GetUserDetails(api, username)
			product := database.GetProductFromID(api, pid, token)

			if result.Requested {
				email.SendEmail(api,
					user.Username,
					user.Email,
					"Someone has requested " + product.Items[0].Product_name,
					"Hello, " + user.Username + "\nSomeone has requested " + result.Title + "\n <img src=\"https://www.karite.xyz" + product.Items[0].Images[0].Size.Medium + "\">")

			}
			return result

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

func CancelRequestProduct(api router.API, pid string, token string, username string) bool {
	if ValidToken(token) {
		if database.GetAvailability(api, pid).Available {
			if IsOwner(api, token, pid) {
				database.SendCancelRequestProductAsOwner(api, pid, username)
			} else {
				database.SendCancelRequestProduct(api, pid, token)
			}
			return true
		} else {
			return false
		}
	} else {
		return false
	}
	return false
}

func GetAdminProductsRequests(api router.API, token string, step int, count int) database.OwnerItems {
	return database.GetAdminProductsRequests(api, step, count)
}

func AdminAuthorizeListingRequest(api router.API, pid string, token string) bool {
	//if isAdmin(api, token) {
	return database.AdminAuthorizeListingRequest(api, pid, token)
	//}

	return false
}