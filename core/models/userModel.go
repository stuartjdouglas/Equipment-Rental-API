package models

import (
	"../router"
	"../utils"
	"fmt"
	"time"
	"log"
)

type user struct {
	Username 	string 	`json:"username"`
	Bio 		string    `json:"bio"`
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

type auth struct {
	Username 	string `json:"username"`
	Token 		string `json:"token"`
	Expiry		time.Time    `json:"expiry"`
}

type userProfile struct {
	Username 		string `json:"username"`
	Bio 			string `json:"bio"`
	Email			string `json:"email"`
	First_name		string `json:"first_name"`
	Last_name		string `json:"last_name"`
	Location 		string `json:"location"`
	Date_registered	string `json:"date_registered"`
}

// Logins in the user and returns an access token
func LoginUser(api router.API, username string, password string, ip_address string) auth {
	login := auth{}
	if getAuthUser(api, username, password) {
		// Generate the token
		token := utils.GenerateToken(username);
		// Add the token to the database
		return addUserToken(api, username, token, ip_address);

	} else {
		return login
	}
	return login
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


// Adds a token into the database
func addUserToken(api router.API, username string, token string, ip_address string) auth {
	login := auth{}
	userid := getUserID(api, username)
	stmt, err := api.Context.Session.Prepare("INSERT INTO tokens (token, user_id, date_generated, date_expires, ip_address) values (?, ?, ?, ?, ?)")

	if err != nil {
		panic(err)
	}

	res, err:= stmt.Exec(token, userid, time.Now(), time.Now().AddDate(0, 0, 7), ip_address)
	if (err != nil) {
		panic(err)
	}
//	TODO Remove this
	fmt.Println(res);

	defer stmt.Close()

	login.Username = username
	login.Token = token
	login.Expiry = time.Now().AddDate(0, 0, 7);

	return login;
}

// Checks if a user has given the correct details or not
// TODO change name
func getAuthUser(api router.API, username string, password string) bool {
	var exist bool
	err := api.Context.Session.QueryRow("SELECT EXISTS(SELECT * FROM users WHERE username = '" + username + "' AND password = '" + password + "')").Scan(&exist)
	if (err != nil) {
		panic(err)
	}

	if exist {
		return true
	}
	return false


	fmt.Println(exist)

	if exist {
		return true
	}
	return false


}


// Returns user information when given a username
//noinspection GoUnusedFunction
func GetUser(api router.API, username string) []user {
	stmt, err := api.Context.Session.Prepare("SELECT username, bio FROM users WHERE username = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(username)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	users := []user{}

	for rows.Next() {
		var result user
		err := rows.Scan(
			&result.Username,
			&result.Bio,
		)

		if err != nil {
			panic(err)
		}
		users = append(users, result)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return users;
}

// Checks if a user already exists
func CheckIfUserExists(api router.API, username string) bool {
	var exist bool
	err := api.Context.Session.QueryRow("SELECT EXISTS (SELECT 1 FROM users WHERE username = ?)", username).Scan(&exist)
	if (err != nil) {
		panic(err)
	}

	if exist {
		return true
	}
	return false
}


// Registers the user
func RegisterUser(api router.API, username string, password string, email string) bool {
	stmt, err := api.Context.Session.Prepare("INSERT INTO users (username, password, email, first_name, last_name, location, bio, date_registered) VALUES (?,?,?,?,?,?,?,?)")

	if err != nil {
		return false;
	}


	res, err:= stmt.Exec(username,  utils.ShaSum([]byte(password)), email, "first_name", "last_name", "location", "bio", time.Now())
	if (err != nil) {
		panic(err)
		return false;
	}
//	TODO Remove this: res should not be used or printed
	fmt.Println(res)
	defer stmt.Close()

	return true;
}


//noinspection GoUnusedFunction
func GetUsers(api router.API) []user{
//	SELECT username, bio FROM users;
	stmt, err := api.Context.Session.Prepare("SELECT username, bio FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	users := []user{}

	for rows.Next() {
		var result user
		err := rows.Scan(
			&result.Username,
			&result.Bio,
		)

		if err != nil {
			panic(err)
		}
		users = append(users, result)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return users;
}

func getUserIdFromToken(api router.API, token string) int {
	stmt, err := api.Context.Session.Prepare("SELECT user_id FROM tokens where token=?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(token)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var id int

	for rows.Next() {
		err := rows.Scan(
			&id,
		)

		if err != nil {
			panic(err)
		}
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return id;
}

type session struct{
	Date_generated 	time.Time `json:"date_generated"`
	Date_expires	time.Time `json:"date_expires"`
	Ip_address		string    `json:"ip_address"`
}

func GetSessions(api router.API, token string) []session {
	sessions := []session{}
	userid := getUserIdFromToken(api, token)

	stmt, err := api.Context.Session.Prepare("SELECT date_generated, date_expires, ip_address FROM tokens WHERE user_id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(userid)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var sess session
		err := rows.Scan(
			&sess.Date_generated,
			&sess.Date_expires,
			&sess.Ip_address,
		)

		if err != nil {
			panic(err)
		}
		sessions = append(sessions, sess)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return sessions
}

func GetSession(api router.API, token string) []session {
	sessions := []session{}

	stmt, err := api.Context.Session.Prepare("SELECT date_generated, date_expires, ip_address FROM tokens WHERE token = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(token)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var sess session
		err := rows.Scan(
			&sess.Date_generated,
			&sess.Date_expires,
			&sess.Ip_address,
		)

		if err != nil {
			panic(err)
		}
		sessions = append(sessions, sess)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return sessions
}

func GetProfile(api router.API, token string) []userProfile {
	userid := getUserIdFromToken(api, token)
	stmt, err := api.Context.Session.Prepare("SELECT username, bio, email, first_name, last_name, location, date_registered FROM users WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(userid)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	user := []userProfile{}

	for rows.Next() {
		var result userProfile
		err := rows.Scan(
			&result.Username,
			&result.Bio,
			&result.Email,
			&result.First_name,
			&result.Last_name,
			&result.Location,
			&result.Date_registered,
		)

		if err != nil {
			panic(err)
		}
		user = append(user, result)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return user;
}



