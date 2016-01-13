package models

import (
	"github.com/remony/Equipment-Rental-API/core/router"
	"log"
)

type Index struct {
	Title string `json:"title"`
	Description string `json:"description"`
}

func GetIndex(api router.API) Index {
	stmt, err := api.Context.Session.Prepare("CALL getIndex()")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()


	var index Index
	for rows.Next() {
		err := rows.Scan(
			&index.Title,
			&index.Description,
		)

		if err != nil {
			log.Println("Getting site index")
			panic(err)
		}
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return index;
}
