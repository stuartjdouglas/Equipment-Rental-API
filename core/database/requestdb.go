package database

import (
	"github.com/remony/Equipment-Rental-API/core/router"
	"log"
	"time"
)

type ItemRequests struct {
	Items []ItemRequest `json:"items"`
	Total int        `json:"total"`
}

type ItemRequest struct {
	Product_name                string                `json:"title"`
	Product_id                  string                `json:"id"`
	Date_added                  time.Time        `json:"date_added"`
	Date_updated                time.Time        `json:"date_updated"`
	Product_description         string                `json:"description"`
	Product_rental_period_limit int64                `json:"product_rental_period_limit"`
	Owner                       User                `json:"owner"`
	Image                       Image           `json:"image"`
	Tags                        []Tag                `json:"tags"`
	Requests                    int `json:"request"`
}

type Request struct {
	Requested       bool `json:"requested"`
	Title           string `json:"title"`
	Date_requested  time.Time `json:"date_requested"`
	Num_of_Requests int `json:"num_of_request"`
}

type UserRequests struct {
	Requests []UserRequest `json:"requests"`
	Total    int `json:"total"`
}

type UserRequest struct {
	Username       string `json:"username"`
	Gravatar       string `json:"gravatar"`
	Date_requested time.Time `json:"date_requested"`
}

func GetProductRequestStatus(api router.API, pid string, token string) Request {
	var request Request
	stmt, err := api.Context.Session.Prepare("Call GetRequestStatus(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(pid, token)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&request.Requested,
			&request.Title,
			&request.Date_requested,
			&request.Num_of_Requests,
		)

		if err != nil {
			panic(err)
		}

	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return request
}

func RequestToRent(api router.API, pid string, username string) bool {

	stmt, err := api.Context.Session.Prepare("Call RentFromRequest(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(pid, username)
	if err != nil {
		//log.Fatal(err)
		return false
	}
	defer rows.Close()

	if err = rows.Err(); err != nil {
		//log.Fatal(err)
		return false
	}

	return true
}

func GetProductRequests(api router.API, pid string, token string) UserRequests {
	var content = []UserRequest{}
	stmt, err := api.Context.Session.Prepare("Call OwnerGetProductRequests(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(token, pid)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var result UserRequest
		err := rows.Scan(
			&result.Username,
			&result.Gravatar,
			&result.Date_requested,
		)

		if err != nil {
			panic(err)
		}
		content = append(content, result)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return UserRequests{Requests: content, Total: len(content)}
}

func GetProductsRequests(api router.API, token string, start int, count int) ItemRequests {
	var content = []ItemRequest{}
	stmt, err := api.Context.Session.Prepare("Call OwnerGetRequests(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(token, start, count)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var result ItemRequest
		var imageid int
		err := rows.Scan(
			&result.Owner.Username,
			&result.Owner.Gravatar,
			&result.Product_name,
			&result.Product_id,
			&result.Date_added,
			&result.Date_updated,
			&result.Product_description,
			&result.Product_rental_period_limit,
			&imageid,
			&result.Requests,
		)

		if err != nil {
			panic(err)
		}

		result.Image = GetImage(api, imageid)
		if err != nil {
			panic(err)
		}
		content = append(content, result)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return ItemRequests{Items: content, Total: len(content)}
}

type UserItemRequests struct {
	Requests []UserRequestItem `json:"requests"`
	Total int `json:"total"`
}

type UserRequestItem struct {
	ID          string                `json:"id"`
	Title       string                `json:"title"`
	Description string                `json:"description"`
	Date_requested         time.Time        `json:"date_requested"`
	Images      Image                `json:"images"`
	Owner       User                `json:"owner"`
}

func GetUserRequests(api router.API, token string, start int, count int) UserItemRequests {
	var content = []UserRequestItem{}
	stmt, err := api.Context.Session.Prepare("Call UserGetOngoingRequests(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(token, start, count)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var result UserRequestItem
		var image_id int
		var username string
		err := rows.Scan(
			&result.ID,
			&result.Title,
			&result.Description,
			&result.Date_requested,
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

	return UserItemRequests{Requests: content, Total: len(content)}
}

func SendRequestProduct(api router.API, pid string, token string) Request {
	var req Request
	stmt, err := api.Context.Session.Prepare("CALL RequestToBorrowItem(?, ?)")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(pid, token)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&req.Requested,
			&req.Title,
			&req.Date_requested,
		)

		if err != nil {
			panic(err)
		}
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return req
}

func SendCancelRequestProduct(api router.API, pid string, token string) bool {
	stmt, err := api.Context.Session.Prepare("CALL CancelRequest(?, ?)")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(pid, token)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	if err = rows.Err(); err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func SendCancelRequestProductAsOwner(api router.API, pid string, username string) bool {
	stmt, err := api.Context.Session.Prepare("CALL OwnerDropRequest(?, ?)")

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(pid, username)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	if err = rows.Err(); err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
