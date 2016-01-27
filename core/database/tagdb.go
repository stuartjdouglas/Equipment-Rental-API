package database

import (
	"github.com/remony/Equipment-Rental-API/core/router"
	"log"
)

func AddTagToProduct(api router.API, pid string, tag string) {
	stmt, err := api.Context.Session.Prepare("CALL addTag(?, ?)")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(pid, tag)
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

}
