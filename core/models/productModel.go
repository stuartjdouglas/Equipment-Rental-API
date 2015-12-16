package models
import (
	"time"
	"log"
	"github.com/remony/Equipment-Rental-API/core/router"
	"github.com/remony/Equipment-Rental-API/core/utils"
	"github.com/remony/Equipment-Rental-API/core/models/sessions"
	"strconv"
)

type Items struct {
	Item []Item `json:"item"`
	Total int 	`json:"total"`
}

type Item struct {
	Product_name			string 		`json:"title"`
	Product_id			string 		`json:"id"`
	Date_added 			time.Time	`json:"date_added"`
	Date_updated			time.Time 	`json:"date_updated"`
	Product_description		string		`json:"description"`
	Product_rental_period_limit	int64 		`json:"product_rental_period_limit"`
	Owner 				user    		`json:"owner"`
	Image				Images           `json:"image"`
}

func CreateProduct(api router.API, product_name string, product_description string, product_rental_period_limit int, token string, file_name string) bool {
	userid := sessions.GetUserIdFromToken(api, token)
	product_id := utils.GenerateUUID();

	stmt, err := api.Context.Session.Prepare("INSERT INTO products (product_name, product_id, date_added, date_updated, product_description, product_rental_period_limit, product_image_id, product_rented_to, users_id, owner_id) values (?,?,?,?,?,?,?,?,?,?)")

	if err != nil {
		panic(err)
	}

	res, err:= stmt.Exec(product_name, product_id, time.Now(), time.Now(), product_description, product_rental_period_limit, file_name, userid, userid, userid)
	if (err != nil) {
		log.Println(err)
		return false
	}

	_ = res

	defer stmt.Close()
	return true
}

func GetProducts (api router.API) Items {
	var content = []Item{}
	stmt, err := api.Context.Session.Prepare("SELECT product_name, product_id, date_added, date_updated, product_description, product_rental_period_limit, owner_id, product_image_id FROM products")
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
		var image_filename string
		var tmpuserid string
		err := rows.Scan(
			&result.Product_name,
			&result.Product_id,
			&result.Date_added,
			&result.Date_updated,
			&result.Product_description,
			&result.Product_rental_period_limit,
			&tmpuserid,
			&image_filename,
		)

		result.Image = GetImage(api, image_filename)
		userid, err := strconv.Atoi(tmpuserid)
		if err != nil {
			panic(err)
		}
		result.Owner = GetUser(api, getUsername(api, userid))

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
	stmt, err := api.Context.Session.Prepare("SELECT product_name, product_id, date_added, date_updated, product_description, product_rental_period_limit, owner_id, product_image_id FROM products where users_id=? ")
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
		var image_filename string
		var tmpuserid string
		err := rows.Scan(
			&result.Product_name,
			&result.Product_id,
			&result.Date_added,
			&result.Date_updated,
			&result.Product_description,
			&result.Product_rental_period_limit,
			&tmpuserid,
			&image_filename,
		)

		if (image_filename != "nil") {
			result.Image = GetImage(api, image_filename)
		}

		userid, err := strconv.Atoi(tmpuserid)
		if err != nil {
			panic(err)
		}
		result.Owner = GetUser(api, getUsername(api, userid))

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
	stmt, err := api.Context.Session.Prepare("SELECT product_name, product_id, date_added, date_updated, product_description, product_rental_period_limit, owner_id, product_image_id FROM products where product_id=? ")
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
		var image_filename string
		var useridtmp string
		err := rows.Scan(
			&result.Product_name,
			&result.Product_id,
			&result.Date_added,
			&result.Date_updated,
			&result.Product_description,
			&result.Product_rental_period_limit,
			&useridtmp,
			&image_filename,
		)
		if err != nil {
			panic(err)
		}
		result.Image = GetImage(api, image_filename)

		userid, err := strconv.Atoi(useridtmp)
		if err != nil {
			panic(err)
		}
		result.Owner = GetUser(api, getUsername(api, userid))
		content = append(content, result)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return Items{Item: content, Total: len(content)}
}