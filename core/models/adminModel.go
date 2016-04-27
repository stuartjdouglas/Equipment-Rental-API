package models

import (
	"github.com/remony/Equipment-Rental-API/core/router"
	"github.com/remony/Equipment-Rental-API/core/database"
)

// Update the index (the site name)
func UpdateIndex(api router.API, title string, Description string, token string) bool {
	return database.UpdateIndex(api, title, Description, token)
}

// Delete a user
func DeleteUser(api router.API, uid string, token string) bool {
	user := database.GetUserRoleFromToken(api, token)
	// Only if the user role is admin should it delete the user
	if user.Role == "admin" {
		return database.DeleteUser(api, uid, token)
	}
	return false
}