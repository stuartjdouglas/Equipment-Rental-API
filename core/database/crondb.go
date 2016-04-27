package database

import (
	"log"
	"github.com/remony/Equipment-Rental-API/core/router"
)

func GetRentalsDueLessThreeDays(api router.API) RentResult {
	var content = []RentItems{}
	stmt, err := api.Context.Session.Prepare("CALL getRentalsDueThreeDays()")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var result RentItems
		var image_id int
		var username string
		err := rows.Scan(
			&result.ID,
			&result.Title,
			&result.Description,
			&result.Due,
			&result.Received,
			&image_id,
			&username,
		)

		result.Images = GetImage(api, image_id)
		if err != nil {
			log.Println("Getting paged results error scanning")
			panic(err)
		}
		result.Owner = GetUser(api, username, false)

		if err != nil {
			panic(err)
		}
		content = append(content, result)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return RentResult{
		Items: content,
		Total: len(content),
	}
}

