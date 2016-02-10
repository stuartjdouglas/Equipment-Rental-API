package models

import (
	"github.com/remony/Equipment-Rental-API/core/router"
	"github.com/remony/Equipment-Rental-API/core/database"
)

func AddComment(api router.API, token string, pid string, comment string) bool {
	if IsSessionValid(api, token) {
		database.AddComment(api, token, pid, comment)
		return true
	}
	return false
}

func DeleteComment(api router.API, pid string, cid string, token string) {
	if IsSessionValid(api, token) {
		database.DeleteComment(api, pid, cid, token)
	}
}
