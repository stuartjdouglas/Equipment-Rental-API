package models

import (
	"time"
	"log"
	"github.com/remony/Equipment-Rental-API/core/router"
)

type Images struct {
	Images 	[]Image `json:"image"`
	Total 	int 	`json:"total"`
}

type Image struct {
	Title 			string 		`json:"title"`
	Location 		string 		`json:"location"`
	Date_added 		time.Time 	`json:"date_added"`
	File_location 	string 		`json:"file_location"`
}

func GetImage(api router.API, filename string) Images {
	var images = []Image{}
	stmt, err := api.Context.Session.Prepare("SELECT file_name, title, date_added FROM images where file_name = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var image Image
		err := rows.Scan(
			&image.Location,
			&image.Title,
			&image.Date_added,
		)

		if err != nil {
			panic(err)
		}
		image.File_location = "/data/" + image.Location
		images = append(images, image)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return Images{Images:images, Total:len(images)}
}

func GetAllImages(api router.API) Images {
	var images = []Image{}
	stmt, err := api.Context.Session.Prepare("SELECT file_name, title, date_added FROM images ORDER BY date_added DESC")
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
		var image Image
		err := rows.Scan(
			&image.Location,
			&image.Title,
			&image.Date_added,
		)

		if err != nil {
			panic(err)
		}
		image.File_location = "/data/" + image.Location
		images = append(images, image)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return Images{Images:images, Total:len(images)}
}

// IsImageAvilable queries the database if the filename already exists
func IsImageAvailable(api router.API, token string) bool {
	var exist bool
	err := api.Context.Session.QueryRow("SELECT EXISTS (SELECT 1 FROM images WHERE file_name = ?)", token).Scan(&exist)
	if (err != nil) {
		// TODO remove panic
		panic(err)
	}
	// If it exists return true
	if exist {
		return true
	}
	// Otherwise return false
	return false
}

func AddImageLocationToDb (api router.API, filename string, title string, original_name string, token string) bool {
	userid := getUserIdFromToken(api, token)
	stmt, err := api.Context.Session.Prepare("INSERT INTO images (file_name, title, date_added, original_name, users_id) values (?, ?, ?, ?, ?)")

	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(filename, title, time.Now(), original_name, userid)
	if (err != nil) {
		return false
	}

	defer stmt.Close()
	return true
}


