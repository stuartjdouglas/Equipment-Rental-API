package models

import (
	"fmt"
	"time"
	"log"
	"crypto/md5"
	"encoding/hex"
	"github.com/remony/Equipment-Rental-API/core/router"
	"github.com/remony/Equipment-Rental-API/core/models/sessions"
)

type User struct {
	Username 	string 	`json:"username"`
	Gravatar	string  `json:"gravatar"`
}

type tempUser struct {
	Username	string    `json:"username"`
	Email		string    `json:"email"`
}

type Profile struct {
	First_name 		string `json:"first_name"`
	Last_name 		string `json:"last_name"`
	Date_registered string `json:"date_registered"`
}

type fullUser struct {
	Username 	string `json:"username", bson:"username"`
	Email 		string `json:"email", bson:"email"`
	Password 	string `json:"password", bson:"password"`
}



type UserProfile struct {
	ID int `json:"id"`
	Username 		string 		`json:"username"`
	Email			string 		`json:"email"`
	First_name		string 		`json:"first_name"`
	Last_name		string 		`json:"last_name"`
	Location 		string 		`json:"location"`
	Date_registered	time.Time 	`json:"date_registered"`
	Gravatar		string    	`json:"gravatar"`
}

type profile struct {
	Profile	UserProfile    `json:"profile"`
}

// Returns the userid when given a username
func getUserID(api router.API, username string) int {
	var id int
	err := api.Context.Session.QueryRow("SELECT id FROM users WHERE username=?", username).Scan(&id)
	if (err != nil) {
		panic(err)
	}
	return id
}

// Returns User information when given a username
//noinspection GoUnusedFunction
func GetUser(api router.API, id string) User {
	stmt, err := api.Context.Session.Prepare("SELECT username, email FROM users WHERE username = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var User User;
	for rows.Next() {
		var result tempUser
		err := rows.Scan(
			&result.Username,
			&result.Email,
		)

		if err != nil {
			log.Println(err)
		}
		User.Username = result.Username
		sum := md5.Sum([]byte(result.Email))
		User.Gravatar = hex.EncodeToString(sum[:])
	}
	if err != nil {
		log.Println(err)
	}
	return User;
}

// Checks if a User already exists
func CheckIfUserExists(api router.API, username string) bool {
	var exist bool
	err := api.Context.Session.QueryRow("CALL doesUserExist(?)", username).Scan(&exist)
	if (err != nil) {
		panic(err)
	}

	if exist {
		return true
	}
	return false
}





//noinspection GoUnusedFunction
func GetUsers(api router.API) []User{
//	SELECT username, bio FROM users;
	stmt, err := api.Context.Session.Prepare("SELECT username, email FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	users := []User{}

	for rows.Next() {
		var result tempUser
		err := rows.Scan(
			&result.Username,
			&result.Email,
		)

		if err != nil {
			panic(err)
		}

		sum := md5.Sum([]byte(result.Email))
		gravatar := hex.EncodeToString(sum[:])
		users = append(users, User{
			Username:result.Username,
			Gravatar:gravatar,
		})
	}
	if err != nil {
		log.Fatal(err)
	}
	return users;
}





// TODO fix json so you don't have to parse [0] to get values
func GetProfile(api router.API, token string) profile {
	userid := sessions.GetUserIdFromToken(api, token)
	stmt, err := api.Context.Session.Prepare("SELECT username, email, first_name, last_name, location, date_registered FROM users WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(userid)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	User := UserProfile{}

	for rows.Next() {
		err := rows.Scan(
			&User.Username,
			&User.Email,
			&User.First_name,
			&User.Last_name,
			&User.Location,
			&User.Date_registered,
		)


		if err != nil {
			panic(err)
		}
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	sum := md5.Sum([]byte(User.Email))
	User.Gravatar = hex.EncodeToString(sum[:])
	User.ID = 0;

	return profile{Profile:User};
}

type hello struct {
	Message string `json:"message"`
}

func GetHello(api router.API, token string) hello {

	author := getUsername(api, sessions.GetUserIdFromToken(api, token))
	message := fmt.Sprintf("こんにちは, %s!", author)

	return hello{Message:message}

}



