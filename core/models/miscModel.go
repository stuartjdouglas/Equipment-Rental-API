package models

import (
	"github.com/remony/Equipment-Rental-API/core/router"
	"github.com/remony/Equipment-Rental-API/core/database"
)

func Like(api router.API, pid string, token string) bool {
	if IsSessionValid(api, token) {
		database.LikeProduct(api, pid, token)
		return true
	}
	return false
}
func UnLike(api router.API, pid string, token string) bool {
	if IsSessionValid(api, token) {
		database.UnLikeProduct(api, pid, token)
		return true
	}
	return false
}

