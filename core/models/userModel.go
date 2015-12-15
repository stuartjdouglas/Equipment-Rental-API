package models

import (
	"fmt"
	"time"
	"log"
	"crypto/md5"
	"encoding/hex"
	"github.com/remony/Equipment-Rental-API/core/router"
	"github.com/remony/Equipment-Rental-API/core/utils"
	"github.com/remony/Equipment-Rental-API/core/models/sessions"
)

type user struct {
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

type Auth struct {
	Success		bool 		`json:"success"`
	Username 	string 		`json:"username"`
	Gravatar	string      	`json:"gravatar"`
	Token 		string 		`json:"token"`
	Expiry		time.Time   	`json:"expiry"`
}

type userProfile struct {
	Username 		string 		`json:"username"`
	Email			string 		`json:"email"`
	First_name		string 		`json:"first_name"`
	Last_name		string 		`json:"last_name"`
	Location 		string 		`json:"location"`
	Date_registered	time.Time 	`json:"date_registered"`
	Gravatar		string    	`json:"gravatar"`
}

type profile struct {
	Profile	userProfile    `json:"profile"`
}
// Logins in the user and returns an access token
func LoginUser(api router.API, username string, password string) Auth {
	var login Auth
	if getAuthUser(api, username, password) {
		// Generate the token
		token := utils.GenerateToken(username)
		indef := utils.GenerateToken(username)
		// Add the token to the database
		return addUserToken(api, username, token, indef, true);

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
func addUserToken(api router.API, username string, token string, idenf string, active bool) Auth {
	var login Auth
	userid := getUserID(api, username)
	stmt, err := api.Context.Session.Prepare("INSERT INTO tokens (token, user_id, date_generated, date_expires, idenf, active) values (?, ?, ?, ?, ?, ?)")

	if err != nil {
		panic(err)
	}

	res, err:= stmt.Exec(token, userid, time.Now(), time.Now().AddDate(0, 0, 7), idenf, active)
	if (err != nil) {
		panic(err)
	}
//	TODO Remove this
	fmt.Println(res);

	defer stmt.Close()
	login.Success = true
	login.Username = username
	login.Token = token
	login.Gravatar = getGravatarString(api, token)
	login.Expiry = time.Now().AddDate(0, 0, 7);

	return login;
}

func getGravatarString(api router.API, token string) string {
	userid := sessions.GetUserIdFromToken(api, token)

	stmt, err := api.Context.Session.Prepare("SELECT email FROM users where id=?")

	if err != nil {
		panic(err)
	}

	rows, err := stmt.Query(userid)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	user := user{}

	for rows.Next() {
		var result tempUser
		err := rows.Scan(
			&result.Email,
		)



		if err != nil {
			panic(err)
		}

		sum := md5.Sum([]byte(result.Email))
		gravatar := hex.EncodeToString(sum[:])
		user.Username = result.Username;
		user.Gravatar = gravatar;

	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return user.Gravatar;
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
func GetUser(api router.API, username string) user {
	stmt, err := api.Context.Session.Prepare("SELECT username, email FROM users WHERE username = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(username)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var user user;
	for rows.Next() {
		var result tempUser
		err := rows.Scan(
			&result.Username,
			&result.Email,
		)

		if err != nil {
			panic(err)
		}
		user.Username = result.Username
		sum := md5.Sum([]byte(result.Email))
		user.Gravatar = hex.EncodeToString(sum[:])
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return user;
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
	stmt, err := api.Context.Session.Prepare("INSERT INTO users (username, password, email, first_name, last_name, location, date_registered) VALUES (?,?,?,?,?,?,?)")

	if err != nil {
		log.Fatal(err)
		return false;
	}


	res, err:= stmt.Exec(username,  utils.Sha512Me([]byte(password)), email, "first_name", "last_name", "location", time.Now())
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

	users := []user{}

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
		users = append(users, user{
			Username:result.Username,
			Gravatar:gravatar,
		})
	}
	if err = rows.Err(); err != nil {
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

	user := userProfile{}

	for rows.Next() {
		err := rows.Scan(
			&user.Username,
			&user.Email,
			&user.First_name,
			&user.Last_name,
			&user.Location,
			&user.Date_registered,
		)


		if err != nil {
			panic(err)
		}
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	sum := md5.Sum([]byte(user.Email))
	user.Gravatar = hex.EncodeToString(sum[:])


	return profile{Profile:user};
}

type hello struct {
	Message string `json:"message"`
}

func GetHello(api router.API, token string) hello {

	author := getUsername(api, sessions.GetUserIdFromToken(api, token))
	message := fmt.Sprintf("こんにちは, %s!", author)

	return hello{Message:message}

}



