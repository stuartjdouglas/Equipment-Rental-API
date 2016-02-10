package database

import (
	"github.com/remony/Equipment-Rental-API/core/router"
	"log"
)

func AddComment(api router.API, token string, pid string, comment string) bool {

	stmt, err := api.Context.Session.Prepare("CALL AddComment(?, ?, ?)")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(token, pid, comment)

	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	return true

}
func DeleteComment(api router.API, pid string, cid string, token string) bool {

	stmt, err := api.Context.Session.Prepare("CALL DeleteComment(?, ?)")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(token, cid)

	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	return true

}
