package database

import (
	"github.com/remony/Equipment-Rental-API/core/router"
	"log"
)

func SendRequestProduct(api router.API, pid string, token string) bool {
	stmt, err := api.Context.Session.Prepare("CALL RequestToBorrowItem(?, ?)")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(pid, token)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()


	if err = rows.Err(); err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func SendCancelRequestProduct(api router.API, pid string, token string) bool {
	stmt, err := api.Context.Session.Prepare("CALL cancelProductRequest(?, ?)")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(pid, token)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()


	if err = rows.Err(); err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
