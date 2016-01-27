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

func RemoveTag(api router.API, pid string, tag string) {
	stmt, err := api.Context.Session.Prepare("CALL removeTag(?, ?)")
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

func GetTags(api router.API, pid string) []Tag{
	var tags []Tag
	stmt, err := api.Context.Session.Prepare("CALL getTags(?)")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(pid)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()


	for rows.Next() {
		var index Tag
		err := rows.Scan(
			&index.Title,
		)

		if err != nil {
			log.Println("Getting " + pid + " tags")
			panic(err)
		}
		tags = append(tags, index);
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return tags
}
