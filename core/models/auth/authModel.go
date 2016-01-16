package auth

import (
	"golang.org/x/crypto/bcrypt"
"strings"
"github.com/remony/Equipment-Rental-API/core/database"
	"github.com/remony/Equipment-Rental-API/core/router"
	"log"
	"regexp"
)


type Error_response struct {
	Message string `json:"message"`
}
func authLogin(password string, digest string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(digest), []byte(password)); err != nil {
		return false
	} else {
		return true
	}
}

func isValidUsername(username string) bool{
	if len(username) < 240 {
		if ok, _ := regexp.MatchString("^[A-Za-z0-9]+$", username); ok {
			log.Println(ok)
			return true
		} else {
			return false
		}
	}


	return false
}

func isValidPassword(password string) bool {
	return true
}

func PerformLogin(api router.API, username string, password string) database.Auth {
	var digest = database.GetDigest(api, username)
	var login database.Auth
	if isValidUsername(username) && isValidPassword(password) {
		if(authLogin(password, digest)) {
			log.Println("good login")
			login = database.LoginUser(api, strings.ToLower(username))
			return login;

		} else {
			log.Println("bad login")
			return database.Auth {
				Success: false,
			}
		}
	} else {
		return database.Auth {
			Success: false,
		}
	}

}