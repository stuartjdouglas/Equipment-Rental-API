package database

import (
	"time"
	"log"
	"github.com/remony/Equipment-Rental-API/core/router"
	"strings"
)

type Items struct {
	Items []Item `json:"items"`
	Total int        `json:"total"`
}

type Item struct {
	Product_name                string                `json:"title"`
	Product_id                  string                `json:"id"`
	Date_added                  time.Time        `json:"date_added"`
	Date_updated                time.Time        `json:"date_updated"`
	Product_description         string                `json:"description"`
	Product_rental_period_limit int64                `json:"product_rental_period_limit"`
	Owner                       User                `json:"owner"`
	Image                       Image           `json:"image"`
	Tags                        []Tag                `json:"tags"`
}

type Tag struct {
	Title string `json:"tag"`
}

type Result struct {
	Results Items                `json:"results"` // The returned results
	Total   int                `json:"total"`     // The total number of results
}

func CreateProduct(api router.API, product_name string, product_description string, product_rental_period_limit int, token string, file_name string, product_id string) bool {
	userid := GetUserIdFromToken(api, token)
	log.Println(userid)

	//	stmt, err := api.Context.Session.Prepare("INSERT INTO products (product_name, product_id, date_added, date_updated, product_description, product_rental_period_limit, product_image_id, owner_id) values (?,?,?,?,?,?,?,?)")
	stmt, err := api.Context.Session.Prepare("CALL createProduct(?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(product_name, product_id, time.Now(), time.Now(), product_description, product_rental_period_limit, file_name, userid)
	if (err != nil) {
		log.Println(err)
		return false
	}

	_ = res

	defer stmt.Close()
	return true
}

func GetProducts(api router.API) Items {
	var content = []Item{}
	stmt, err := api.Context.Session.Prepare("CALL getListing()")

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
		var postid int
		err := rows.Scan(
			&result.Owner.Username,
			&result.Owner.Gravatar,
			&result.Product_name,
			&result.Product_id,
			&result.Date_added,
			&result.Date_updated,
			&result.Product_description,
			&result.Product_rental_period_limit,
			&postid,
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

func filterTags(tag string) []Tag {
	var tags []Tag
	tmp := strings.Split(tag, ", ")
	for i := 0; i < len(tmp); i++ {
		var newTag Tag
		newTag.Title = tmp[i]
		tags = append(tags, newTag)
	}
	return tags
}

func getTags(api router.API, pid string) []Tag {
	var tags = []Tag{}

	stmt, err := api.Context.Session.Prepare("CALL GetTags(?)")
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
		var result Tag
		err := rows.Scan(
			&result.Title,
		)

		if err != nil {
			log.Println("Getting tags scanning")
			panic(err)
		}

		tags = append(tags, result)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return tags
}

func GetProductsPaging(api router.API, step int, count int) Items {

	var content = []Item{}

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
		var image_id int
		//var tmpuserid string
		err := rows.Scan(
			&result.Product_id,
			&result.Product_name,
			&result.Product_description,
			&result.Date_added,
			&result.Date_updated,
			&result.Product_rental_period_limit,
			&image_id,
			&result.Owner.Username,
			&result.Owner.Gravatar,
		)

		log.Println(result.Product_id)
		result.Tags = getTags(api, result.Product_id)

		result.Image = GetImage(api, image_id)

		if err != nil {
			log.Println("Getting paged results error scanning")
			panic(err)
		}

		content = append(content, result)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	var items Items

	items.Items = content
	items.Total = len(content)

	return items
}

type RentItems struct {
	ID          string                `json:"id"`
	Title       string                `json:"title"`
	Description string                `json:"description"`
	Due         time.Time        `json:"due"`
	Received    time.Time        `json:"received"`
	Images      Image                `json:"images"`
	Owner       User                `json:"owner"`
}

type RentResult struct {
	Items []RentItems        `json:"items"`
	Total int                `json:"total"`
}

func GetCurrentlyRentedProducts(api router.API, token string, step int, count int) RentResult {

	var content = []RentItems{}
	username := GetUserNameFromToken(api, token)
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
		//		userid, err := strconv.Atoi(tmpuserid)
		if err != nil {
			log.Println("Getting paged results error scanning")
			panic(err)
		}
		result.Owner = GetUser(api, username)

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

func GetProductFromOwner(api router.API, username string) Items {
	var content = []Item{}
	//	user:= getUserID(api, username)
	//	stmt, err := api.Context.Session.Prepare("SELECT product_name, product_id, date_added, date_updated, product_description, product_rental_period_limit, owner_id, product_image_id FROM products where users_id=? ")
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	defer stmt.Close()
	//	rows, err := stmt.Query(user)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	defer rows.Close()
	//
	//
	//	for rows.Next() {
	//		var result Item
	//		var image_filename string
	//		var tmpuserid string
	//		err := rows.Scan(
	//			&result.Product_name,
	//			&result.Product_id,
	//			&result.Date_added,
	//			&result.Date_updated,
	//			&result.Product_description,
	//			&result.Product_rental_period_limit,
	//			&tmpuserid,
	//			&image_filename,
	//		)
	//
	//		if (image_filename != "nil") {
	//			result.Image = GetImage(api, image_filename)
	//		}
	//
	//		userid, err := strconv.Atoi(tmpuserid)
	//		if err != nil {
	//			panic(err)
	//		}
	//		result.Owner = GetUser(api, getUsername(api, userid))
	//
	//		if err != nil {
	//			panic(err)
	//		}
	//
	//		content = append(content, result)
	//	}
	//	if err = rows.Err(); err != nil {
	//		log.Fatal(err)
	//	}

	return Items{Items: content, Total: len(content)}
}

func GetProductFromID(api router.API, id string) Items {
	var content = []Item{}
	log.Println("getting product " + id)
	stmt, err := api.Context.Session.Prepare("Call getProduct(?)")
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
		var imageid int
		var username string
		var tags string
		err := rows.Scan(
			&result.Product_name,
			&result.Product_id,
			&result.Date_added,
			&result.Date_updated,
			&result.Product_description,
			&result.Product_rental_period_limit,
			&username,
			&imageid,
			&tags,
		)

		if err != nil {
			panic(err)
		}

		result.Tags = filterTags(tags)
		result.Image = GetImage(api, imageid)
		if err != nil {
			panic(err)
		}
		result.Owner = GetUser(api, username)
		content = append(content, result)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return Items{Items: content, Total: len(content)}
}

type Availability struct {
	Available bool `json:"available"`
	Date      time.Time `json:"date"`
}

type RentalStatus struct {
	Owner      bool            `json:"owner"`
	Available  bool        `json:"available"`
	Date_taken time.Time   `json:"date_taken"`
	Date_due   time.Time   `json:"date_due"`
}

type OwnerRentalStatus struct {
	Owner      string                `json:"owner"`
	Available  bool                `json:"available"`
	Date_taken time.Time        `json:"date_taken"`
	Date_due   time.Time        `json:"date_due"`
}

func GetAvailability(api router.API, product string) Availability {
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

func GetAuthedAvailability(api router.API, product string, token string) RentalStatus {
	var currentProductRenter string
	var available bool

	username := GetUserNameFromToken(api, token)

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

func RentItem(api router.API, product string, token string) Availability {
	username := GetUserNameFromToken(api, token)
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

func ReturnItem(api router.API, product string, token string) Availability {
	stmt, err := api.Context.Session.Prepare("CALL ReturnItem(?, ?)")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(token, product)
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

func ReturnItemAsOwner(api router.API, product string, token string) Availability {
	stmt, err := api.Context.Session.Prepare("CALL ReturnItemAsOwner(?, ?)")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(token, product)
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
			log.Println("Error returning as owner")
			panic(err)
		}
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return result;
}

func IsOwner(api router.API, token string, product string) bool {
	stmt, err := api.Context.Session.Prepare("CALL isOwner(?, ?)")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(token, product)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var result bool
	for rows.Next() {
		err := rows.Scan(
			&result,
		)

		if err != nil {
			panic(err)
		}
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return result;
}

func RemoveProduct(api router.API, pid string, token string) {
	stmt, err := api.Context.Session.Prepare("CALL removeProduct(?, ?)")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(token, pid)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
}

// Returns products that belongs to the owner
func GetOwnerProductsPaging(api router.API, token string, step int, count int) Items {

	var content = []Item{}

	stmt, err := api.Context.Session.Prepare("CALL getOwnerProducts(?, ?, ?)")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(token, step, count)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var result Item
		var image_id int
		var tmpuserid string
		err := rows.Scan(
			&result.Product_id,
			&result.Product_name,
			&result.Product_description,
			&result.Date_added,
			&result.Date_updated,
			&result.Product_rental_period_limit,
			&image_id,
			&tmpuserid,
		)

		result.Image = GetImage(api, image_id)

		if err != nil {
			log.Println("Getting paged results error scanning")
			panic(err)
		}

		content = append(content, result)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return Items{Items:content, Total:len(content)}
}

func GetOwnerProductAvailability(api router.API, product string, token string) OwnerRentalStatus {
	stmt, err := api.Context.Session.Prepare("CALL ownerProductStatus(?, ?)")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(token, product)

	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var result OwnerRentalStatus

	for rows.Next() {
		err := rows.Scan(
			&result.Available,
			&result.Date_due,
			&result.Date_taken,
			&result.Owner,
		)

		if err != nil {
			panic(err)
		}
	}
	return result
}
