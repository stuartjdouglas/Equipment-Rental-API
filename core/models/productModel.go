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
	Image				Image           `json:"image"`
}

type Result struct {
	Results	Items 		`json:"results"` // The returned results
	Total	int        	`json:"total"` // The total number of results
}

func CreateProduct(api router.API, product_name string, product_description string, product_rental_period_limit int, token string, file_name string) bool {
	userid := sessions.GetUserIdFromToken(api, token)
	product_id := utils.GenerateUUID();

	stmt, err := api.Context.Session.Prepare("INSERT INTO products (product_name, product_id, date_added, date_updated, product_description, product_rental_period_limit, product_image_id, owner_id) values (?,?,?,?,?,?,?,?)")

	if err != nil {
		panic(err)
	}

	res, err:= stmt.Exec(product_name, product_id, time.Now(), time.Now(), product_description, product_rental_period_limit, file_name, userid)
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
	stmt, err := api.Context.Session.Prepare("SELECT product_name, product_id, date_added, date_updated, product_description, product_rental_period_limit, owner_id, product_image_id FROM products. ORDER BY date_updated DESC")

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

func GetProductsPaging (api router.API, step int, count int) Result {

	var content = []Item{}

//	stmt, err := api.Context.Session.Prepare("SELECT product_name, product_id, date_added, date_updated, product_description, product_rental_period_limit, owner_id, product_image_id FROM products ORDER BY date_added DESC LIMIT ?, ?")
	stmt, err := api.Context.Session.Prepare("CALL getPagedProducts(?, ?)")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(step, count)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()


	for rows.Next() {
		var result Item
		var image_filename string
		var tmpuserid string
		err := rows.Scan(
			&result.Product_id,
			&result.Product_name,
			&result.Product_description,
			&result.Date_added,
			&result.Date_updated,
			&result.Product_rental_period_limit,
			&image_filename,
			&tmpuserid,
		)

		result.Image = GetImage(api, image_filename)
//		userid, err := strconv.Atoi(tmpuserid)
		if err != nil {
			log.Println("Getting paged results error scanning")
			panic(err)
		}
//		result.Owner = GetUser(api, getUsername(api, userid))
//
//		if err != nil {
//			panic(err)
//		}
		content = append(content, result)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	var items Items

	items.Item = content
	items.Total = len(content)
	total := getCount(api, "all")
	return Result{Results:items, Total:total}
}

func GetCurrentlyRentedProducts (api router.API, token string, step int, count int) Result {

	var content = []Item{}
	username := sessions.GetUserNameFromToken(api, token)
	//	stmt, err := api.Context.Session.Prepare("SELECT product_name, product_id, date_added, date_updated, product_description, product_rental_period_limit, owner_id, product_image_id FROM products ORDER BY date_added DESC LIMIT ?, ?")
	stmt, err := api.Context.Session.Prepare("CALL getCurrentlyRentingProducts(?, ?, ?)")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(username, step, count)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()


	for rows.Next() {
		var result Item
		var image_filename string
		var tmpuserid string
		err := rows.Scan(
			&result.Product_id,
			&result.Product_name,
			&result.Product_description,
			&result.Date_added,
			&result.Date_updated,
			&result.Product_rental_period_limit,
			&image_filename,
			&tmpuserid,
		)

		result.Image = GetImage(api, image_filename)
		//		userid, err := strconv.Atoi(tmpuserid)
		if err != nil {
			log.Println("Getting paged results error scanning")
			panic(err)
		}
		//		result.Owner = GetUser(api, getUsername(api, userid))
		//
		//		if err != nil {
		//			panic(err)
		//		}
		content = append(content, result)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	var items Items

	items.Item = content
	items.Total = len(content)
	total := getCount(api, "all")
	return Result{Results:items, Total:total}
}

//func GetAvailability (api router.API, token string, step int, count int) Result {
//
//	var content = []Item{}
//	username := sessions.GetUserNameFromToken(api, token)
//	//	stmt, err := api.Context.Session.Prepare("SELECT product_name, product_id, date_added, date_updated, product_description, product_rental_period_limit, owner_id, product_image_id FROM products ORDER BY date_added DESC LIMIT ?, ?")
//	stmt, err := api.Context.Session.Prepare("CALL getRentedProducts(?, ?, ?)")
//	if err != nil {
//		log.Println(err)
//	}
//	defer stmt.Close()
//	rows, err := stmt.Query(username, step, count)
//	if err != nil {
//		log.Println(err)
//	}
//	defer rows.Close()
//
//
//	for rows.Next() {
//		var result Item
//		var image_filename string
//		var tmpuserid string
//		err := rows.Scan(
//			&result.Product_id,
//			&result.Product_name,
//			&result.Product_description,
//			&result.Date_added,
//			&result.Date_updated,
//			&result.Product_rental_period_limit,
//			&image_filename,
//			&tmpuserid,
//		)
//
//		result.Image = GetImage(api, image_filename)
//		//		userid, err := strconv.Atoi(tmpuserid)
//		if err != nil {
//			log.Println("Getting paged results error scanning")
//			panic(err)
//		}
//		//		result.Owner = GetUser(api, getUsername(api, userid))
//		//
//		//		if err != nil {
//		//			panic(err)
//		//		}
//		content = append(content, result)
//	}
//	if err = rows.Err(); err != nil {
//		log.Fatal(err)
//	}
//
//	var items Items
//
//	items.Item = content
//	items.Total = len(content)
//	total := getCount(api, "all")
//	return Result{Results:items, Total:total}
//}

func getCount(api router.API, query string) int {
	count := 0;
	if (query == "all") {
		stmt, err := api.Context.Session.Prepare("SELECT COUNT(*) from products")
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
			err = rows.Scan(
				&count,
			)

			if err != nil {
				log.Fatal(err)
			}
		}
		if err = rows.Err(); err != nil {
			log.Fatal(err)
		}


	}
	return count;
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

type Availability struct {
	Available bool `json:"available"`
	Date time.Time `json:"date"`
}

type RentalStatus struct {
	Owner		bool  	    `json:"owner"`
	Available	bool        `json:"available"`
	Date_taken	time.Time   `json:"date_taken"`
	Date_due	time.Time   `json:"date_due"`
}

func GetAvailability (api router.API, product string) Availability {

	stmt, err := api.Context.Session.Prepare("CALL checkProductAvailability(?)")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(product)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()


	var result Availability
	for rows.Next() {
		err := rows.Scan(
			&result.Available,
			&result.Date,
		)

		if err != nil {
			log.Println("Getting paged results error scanning")
			panic(err)
		}
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return result;
}

func GetAuthedAvailability (api router.API, product string, token string) RentalStatus {
	var currentProductRenter string
	var available bool

	username := sessions.GetUserNameFromToken(api, token)

	stmt, err := api.Context.Session.Prepare("CALL checkAuthedProductAvailability(?)")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(product)

	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var result RentalStatus

	for rows.Next() {
		err := rows.Scan(
			&available,
			&result.Date_due,
			&result.Date_taken,
			&currentProductRenter,
		)

		if err != nil {
			log.Println("Getting Authed availability scanning")
			panic(err)
		}
		result.Available = available
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	if username != currentProductRenter {
		result.Owner = false;
		result.Date_taken = time.Now()
	} else {
		result.Owner = true;
	}


	return result


}

func RentItem (api router.API, product string, token string) Availability {
	username := sessions.GetUserNameFromToken(api, token)
	stmt, err := api.Context.Session.Prepare("CALL RentItem(?, ?)")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(product, username)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()


	var result Availability
	for rows.Next() {
		err := rows.Scan(
			&result.Available,
			&result.Date,
		)

		if err != nil {
			log.Println("Getting paged results error scanning")
			panic(err)
		}
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return result;
}

func ReturnItem (api router.API, product string, token string) Availability {
	username := sessions.GetUserIdFromToken(api, token)
	stmt, err := api.Context.Session.Prepare("CALL ReturnItem(?, ?)")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(product, username)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()


	var result Availability
	for rows.Next() {
		err := rows.Scan(
			&result.Available,
			&result.Date,
		)

		if err != nil {
			log.Println("Getting paged results error scanning")
			panic(err)
		}
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return result;
}

