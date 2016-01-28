package database

import (
	"github.com/remony/Equipment-Rental-API/core/router"
	"log"
)

func SearchTag(api router.API, tag string, start int, count int) Items {
	var content = []Item{}
	stmt, err := api.Context.Session.Prepare("CALL searchFilterByTag(?, ?, ?)")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(tag, start, count)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var result Item
		var postid int
		err := rows.Scan(
			&result.Product_id,
			&result.Product_name,
			&result.Product_description,
			&result.Date_added,
			&result.Date_updated,
			&result.Product_rental_period_limit,
			&postid,
			&result.Owner.Username,
			&result.Owner.Gravatar,
		)

		result.Tags = getTags(api, result.Product_id);

		result.Image = GetImage(api, postid)

		if err != nil {
			panic(err)
		}
		content = append(content, result)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return Items{Items: content, Total: len(content)}
}
