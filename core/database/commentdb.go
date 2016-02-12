package database

import (
	"github.com/remony/Equipment-Rental-API/core/router"
	"log"
)

func AddComment(api router.API, token string, pid string, comment string, requiresApproval bool) bool {

	stmt, err := api.Context.Session.Prepare("CALL AddComment(?, ?, ?, ?)")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(token, pid, comment, requiresApproval)

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
func DisableComments(api router.API, pid string) bool {

	stmt, err := api.Context.Session.Prepare("CALL DisableComments(?)")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(pid)

	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	return true

}

func EnableComments(api router.API, pid string) bool {

	stmt, err := api.Context.Session.Prepare("CALL EnableComments(?)")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(pid)

	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	return true

}

func ApproveComment(api router.API, pid string, cid string) bool {

	stmt, err := api.Context.Session.Prepare("CALL ApproveComment(?)")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(cid)

	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	return true

}


