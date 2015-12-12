package models
import (
	"time"
	"log"
	"github.com/remony/Equipment-Rental-API/core/router"
	"github.com/remony/Equipment-Rental-API/core/utils"
)

type Items struct {
	Item []Item `json:"item"`
	Total int 	`json:"total"`
}

type Item struct {
	Product_name					string 		`json:"title"`
	Product_id	 					string 		`json:"id"`
	Date_added 						time.Time	`json:"date_added"`
	Date_updated					time.Time 	`json:"date_updated"`
	Product_description				string		`json:"description"`
	Product_rental_period_limit		int64 		`json:"product_rental_period_limit"`
	Owner 							int    		`json:"owner"`
}

func CreateProduct(api router.API, product_name string, product_description string, product_rental_period_limit int, token string) bool {
	userid := getUserIdFromToken(api, token)
	product_id := utils.GenerateUUID();

	stmt, err := api.Context.Session.Prepare("INSERT INTO products (product_name, product_id, date_added, date_updated, product_description, product_rental_period_limit, product_image_id, product_rented_to, users_id, owner_id) values (?,?,?,?,?,?,?,?,?,?)")

	if err != nil {
		panic(err)
	}

	res, err:= stmt.Exec(product_name, product_id, time.Now(), time.Now(), product_description, product_rental_period_limit, 0, userid, userid, userid)
	if (err != nil) {
		log.Println(err)
		return false
	}

	log.Println(res)

	defer stmt.Close()
	return true
}

func GetProducts (api router.API) Items {
	var content = []Item{}
	stmt, err := api.Context.Session.Prepare("SELECT product_name, product_id, date_added, date_updated, product_description, product_rental_period_limit, owner_id FROM products")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()


	for rows.Next() {
		var result Item
		err := rows.Scan(
			&result.Product_name,
			&result.Product_id,
			&result.Date_added,
			&result.Date_updated,
			&result.Product_description,
			&result.Product_rental_period_limit,
			&result.Owner,
		)

		if err != nil {
			panic(err)
		}
		content = append(content, result)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return Items{Item: content, Total: len(content)}
}

func GetProductFromOwner (api router.API, username string) Items {
	var content = []Item{}
	user:= getUserID(api, username)
	stmt, err := api.Context.Session.Prepare("SELECT product_name, product_id, date_added, date_updated, product_description, product_rental_period_limit, owner_id FROM products where users_id=? ")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(user)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()


	for rows.Next() {
		var result Item
		err := rows.Scan(
			&result.Product_name,
			&result.Product_id,
			&result.Date_added,
			&result.Date_updated,
			&result.Product_description,
			&result.Product_rental_period_limit,
			&result.Owner,
		)

		if err != nil {
			panic(err)
		}
		content = append(content, result)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return Items{Item: content, Total: len(content)}
}

func GetProductFromID (api router.API, id string) Items {
	var content = []Item{}
	stmt, err := api.Context.Session.Prepare("SELECT product_name, product_id, date_added, date_updated, product_description, product_rental_period_limit, owner_id FROM products where product_id=? ")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()


	for rows.Next() {
		var result Item
		err := rows.Scan(
			&result.Product_name,
			&result.Product_id,
			&result.Date_added,
			&result.Date_updated,
			&result.Product_description,
			&result.Product_rental_period_limit,
			&result.Owner,
		)

		if err != nil {
			panic(err)
		}
		content = append(content, result)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return Items{Item: content, Total: len(content)}
}