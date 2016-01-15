package database

import (
	"log"
	"github.com/remony/Equipment-Rental-API/core/utils"
	"github.com/remony/Equipment-Rental-API/core/router"
	"time"
)

type Auth struct {
	Success  bool                `json:"success"`
	Username string                `json:"username"`
	Gravatar string        `json:"gravatar"`
	Token    string                `json:"token"`
	Expiry   time.Time        `json:"expiry"`
}

func LoginUser(api router.API, username string, password string) Auth {
	token := utils.GenerateToken(username)
	indef := utils.GenerateToken(username)

	stmt, err := api.Context.Session.Prepare("CALL login(?, ?, ?, ?)")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(username, password, token, indef)

	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var result Auth

	for rows.Next() {
		err := rows.Scan(
			&result.Success,
			&result.Username,
			&result.Gravatar,
			&result.Token,
			&result.Expiry,
		)

		if err != nil {
			panic(err)
		}
	}
	return result
}

// Registers the User
func RegisterUser(api router.API, username string, password string, email string) bool {
	stmt, err := api.Context.Session.Prepare("CALL register(?,?,?,?, ?)")

	if err != nil {
		log.Fatal(err)
		return false;
	}

	res, err := stmt.Exec(username, password, email, "first_name", "last_name")
	if (err != nil) {
		panic(err)
		return false;
	}

	_ = res

	defer stmt.Close()

	return true;
}