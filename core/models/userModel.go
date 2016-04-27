package models

import (
	"github.com/remony/Equipment-Rental-API/core/router"
	"github.com/remony/Equipment-Rental-API/core/database"
)

func GetUserData(api router.API, username string) database.User {

	return database.GetUser(api, username, false)
}

func GetUsersData(api router.API, token string) []database.User {
	return database.GetUsers(api, token)
}

func CheckIfUserExists(api router.API, username string) bool {
	return database.CheckIfUserExists(api, username)
}

func GetProfile(api router.API, token string) database.UserProfileContainer {
	return database.GetProfile(api, token)
}

func ChangeRole(api router.API, username string, role string, token string) bool {
	if database.GetUserNameFromToken(api, token) != username {
		if IsSessionValid(api, token) {
			database.ChangeRole(api, username, role, token)
			return true
		}
	}
	return false
}