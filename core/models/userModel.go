package models

import (
	"github.com/remony/Equipment-Rental-API/core/router"
	"github.com/remony/Equipment-Rental-API/core/database"
)

func GetUserData(api router.API, username string) database.User {
	return database.GetUser(api, username)
}

func GetUsersData(api router.API) []database.User {
	return database.GetUsers(api)
}


func CheckIfUserExists(api router.API, username string) bool {
	return database.CheckIfUserExists(api, username)
}

func GetProfile(api router.API, token string) database.UserProfileContainer {
	return database.GetProfile(api, token)
}
