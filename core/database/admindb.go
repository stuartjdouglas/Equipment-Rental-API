package database

import (
	"github.com/remony/Equipment-Rental-API/core/router"
	"log"
)

func UpdateIndex(api router.API, title string, description string, token string) bool {
	stmt, err := api.Context.Session.Prepare("CALL updateSite(?, ?, ?)")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(title, description, token)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	return true
}

func DeleteUser(api router.API, uid string, token string) bool {
	stmt, err := api.Context.Session.Prepare("CALL removeUserAsAdmin(?, ?)")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(uid, token)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	return true
}