package database

import (
	"fmt"
	"github.com/remony/Equipment-Rental-API/core/router"
	"log"
)

func GetHello(api router.API, token string) hello {

	author := GetUsername(api, GetUserIdFromToken(api, token))
	message := fmt.Sprintf("こんにちは, %s!", author)

	return hello{Message:message}

}

func LikeProduct(api router.API, pid string, token string) bool {

	stmt, err := api.Context.Session.Prepare("CALL `like`(?, ?)")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(token, pid)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	return true
}

func UnLikeProduct(api router.API, pid string, token string) bool {

	stmt, err := api.Context.Session.Prepare("CALL `unLike`(?, ?)")
	if err != nil {
		log.Println(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(token, pid)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	return true
}

