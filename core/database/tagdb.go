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

func GetTags(api router.API, pid string) []Tag {
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

func GetProductsWithTag(api router.API, tag string, start int, count int) Items {
	var content = []Item{}
	stmt, err := api.Context.Session.Prepare("CALL getListingOfTag(?, ?, ?)")

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
			&result.Age_Rating,
		)

		result.Tags = getTags(api, result.Product_id);

		result.Images = GetImage(api, postid)

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

type MostUsedTag struct {
	Title  string `json:"title"`
	Amount int `json:"amount"`
}

func GetTagsMostUsed(api router.API, step int, count int, token string, order bool) []MostUsedTag {
	var content = []MostUsedTag{}
	stmt, err := api.Context.Session.Prepare("CALL getMostUsedTags(?, ?, ?)")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(step, count, order)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var result MostUsedTag

		err := rows.Scan(
			&result.Title,
			&result.Amount,
		)

		if err != nil {
			panic(err)
		}
		content = append(content, result)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return content
}
