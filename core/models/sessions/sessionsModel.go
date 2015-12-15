package sessions
import (
	"github.com/remony/Equipment-Rental-API/core/router"
	"log"
	"time"
)

type session struct{
	Date_generated 		time.Time `json:"date_generated"`
	Date_expires		time.Time `json:"date_expires"`
	Ip_address		string    `json:"ip_address"`
}

func IsSessionValid(api router.API, token string) bool {
	var exist bool
	err := api.Context.Session.QueryRow("SELECT EXISTS (SELECT 1 FROM tokens WHERE token = ?)", token).Scan(&exist)
	if (err != nil) {
		panic(err)
	}

	if exist {
		return true
	}
	return false
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



func GetSessions(api router.API, token string) []session {
	sessions := []session{}
	userid := GetUserIdFromToken(api, token)

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

func GetUserIdFromToken(api router.API, token string) int {
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

