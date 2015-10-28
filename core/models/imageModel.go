package models

import (
	"../router"
)

// IsImageAvilable queries the database if the filename already exists
func IsImageAvailable(api router.API, token string) bool {
	var exist bool
	err := api.Context.Session.QueryRow("SELECT EXISTS (SELECT 1 FROM images WHERE file_name = ?)", token).Scan(&exist)
	if (err != nil) {
		// TODO remove panic
		panic(err)
	}
	// If it exists return true
	if exist {
		return true
	}
	// Otherwise return false
	return false
}
