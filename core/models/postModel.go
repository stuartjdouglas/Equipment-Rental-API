package models
import (
	"time"
	"fmt"
	"log"
	"github/remony/Equipment-Rental-API/core/router"
)

type Posts struct {
	Post []Post `json:"post"`
	Total int 	`json:"total"`
}



type Post struct {
	Title			string 		`json:"title"`
	Slug 			string 		`json:"slug"`
	Author 			string 		`json:"author"`
	Content			string 		`json:"content"`
	Date_created	time.Time   `json:"date_Created"`
	Date_edited		time.Time 	`json:"date_edited"`
}

func CheckIfPostExists (api router.API, slug string) bool {
	var exist bool
	err := api.Context.Session.QueryRow("SELECT EXISTS (SELECT 1 FROM posts WHERE slug = ?)", slug).Scan(&exist)
	if (err != nil) {
		panic(err)
	}

	if exist {
		return true
	}
	return false
}

func getUsername(api router.API, userid int) string {
	stmt, err := api.Context.Session.Prepare("SELECT username FROM users where id=?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(userid)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var username string

	for rows.Next() {
		err := rows.Scan(
			&username,
		)

		if err != nil {
			panic(err)
		}
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return username;
}

func CreatePost(api router.API, post Post, token string, slug string) bool {
	userid := getUserIdFromToken(api, token)
	author := getUsername(api, userid)
	stmt, err := api.Context.Session.Prepare("INSERT INTO posts (title, slug, author, content, date_created, date_edited, users_id) values (?, ?, ?, ?, ?, ?, ?)")

	if err != nil {
		panic(err)
	}

	res, err:= stmt.Exec(post.Title, slug, author, post.Content, time.Now(), time.Now(), userid)
	if (err != nil) {
		return false
	}
	//	TODO Remove this
	fmt.Println(res);

	defer stmt.Close()
	return true
}

func GetPosts (api router.API) Posts {
	var content = []Post{}
	stmt, err := api.Context.Session.Prepare("SELECT title, slug, author, content, date_created, date_edited FROM posts")
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
		var result Post
		err := rows.Scan(
			&result.Title,
			&result.Slug,
			&result.Author,
			&result.Content,
			&result.Date_created,
			&result.Date_edited,
		)

		if err != nil {
			panic(err)
		}
		content = append(content, result)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return Posts{Post: content, Total: len(content)}
}
func GetPostsFromUser (api router.API, username string) Posts {
	var content = []Post{}
	stmt, err := api.Context.Session.Prepare("SELECT title, slug, author, content, date_created, date_edited FROM posts where author = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(username)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()


	for rows.Next() {
		var result Post
		err := rows.Scan(
			&result.Title,
			&result.Slug,
			&result.Author,
			&result.Content,
			&result.Date_created,
			&result.Date_edited,
		)

		if err != nil {
			panic(err)
		}
		content = append(content, result)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return Posts{Post: content, Total: len(content)}
}

func GetPost (api router.API, slug string) Posts {
	var content = []Post{}
	stmt, err := api.Context.Session.Prepare("SELECT title, slug, author, content, date_created, date_edited FROM posts where slug = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(slug)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()


	for rows.Next() {
		var result Post
		err := rows.Scan(
			&result.Title,
			&result.Slug,
			&result.Author,
			&result.Content,
			&result.Date_created,
			&result.Date_edited,
		)

		if err != nil {
			panic(err)
		}
		content = append(content, result)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return Posts{Post: content, Total: len(content)}
}