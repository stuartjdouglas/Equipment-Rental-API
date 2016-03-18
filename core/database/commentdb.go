package database

import (
	"github.com/remony/Equipment-Rental-API/core/router"
	"log"
)

func AddComment(api router.API, token string, pid string, comment string, requiresApproval bool, rating int) string {
	var id string
	stmt, err := api.Context.Session.Prepare("CALL AddComment(?, ?, ?, ?, ?)")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(token, pid, comment, requiresApproval, rating)

	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(
			&id,
		)

		if err != nil {
			panic(err)
		}
	}

	return id

}
func DeleteComment(api router.API, pid string, cid string, token string) Comment {
	var comment Comment
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

	for rows.Next() {
		var id string
		err := rows.Scan(
			&id,
		)

		if err != nil {
			panic(err)
		}

		comment = GetComment(api, id)
	}

	return comment

}
func GetComment(api router.API, cid string) Comment {
	var comment Comment
	stmt, err := api.Context.Session.Prepare("CALL GetComment(?)")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(cid)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&comment.Text,
			&comment.Author.Username,
			&comment.Author.Gravatar,
			&comment.Date_added,
			&comment.Date_updated,
			&comment.ID,
			&comment.Authorized,
			&comment.Rating,
		)

		if err != nil {
			log.Println("Getting comment scanning")
			panic(err)
		}
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return comment
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

func EditComment(api router.API, token string, cid string, comment string, rating int) string {
	var id string
	stmt, err := api.Context.Session.Prepare("CALL EditComment(?, ?, ?, ?)")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(token, cid, comment, rating)

	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(
			&id,
		)

		if err != nil {
			panic(err)
		}
	}

	return id
}

func HaveIReviewed(api router.API, pid string, token string) bool {
	var id int
	var reviewed bool
	stmt, err := api.Context.Session.Prepare("CALL HasUserReviewedListing(?, ?)")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(token, pid)

	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(
			&id,
		)

		log.Println(id)
		if (id == 0) {
			reviewed = false
			log.Println("not reviewed")
		} else {
			log.Println("reviewed")
			reviewed = true
		}

		if err != nil {
			panic(err)
		}
	}

	return reviewed
}


