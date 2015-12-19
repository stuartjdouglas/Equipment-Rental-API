package sessions
import (
	"github.com/remony/Equipment-Rental-API/core/router"
	"log"
	"time"
)

type Session struct{
	Date_generated 		time.Time `json:"date_generated"`
	Date_expires		time.Time `json:"date_expires"`
	Idenf			string    `json:"idenf"`
	Active			bool      `json:"active"`
}

type session struct {
	Session		Session        `json:"session"`
	Total		int            `json:"total"`
}

type Sessions struct {
	Sessions	[]Session        `json:"session"`
	Total		int            `json:"total"`
}

type removalRes	struct {
	ID string `json:"id"`
	Message string `json:"message"`
}

func IsSessionValid(api router.API, token string) bool {
	var exist bool
	err := api.Context.Session.QueryRow("SELECT EXISTS (SELECT 1 FROM tokens WHERE token = ? AND active = 1 AND NOW() <= date_expires)", token).Scan(&exist)
	if (err != nil) {
		panic(err)
	}

	if exist {
		return true
	} else {
		DisableToken(api, token)
	}
	return false
}

func DisableToken(api router.API, token string) bool {
	stmt, err := api.Context.Session.Prepare("UPDATE tokens SET active =? WHERE idenf =? AND active =?")

	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()
	rows, err := stmt.Exec(0, token, 1)

	if err != nil {
		log.Fatal(err)
	}
	rows.RowsAffected()


	return true
}

func GetSession(api router.API, token string) []Session {
	sessions := []Session{}

	stmt, err := api.Context.Session.Prepare("SELECT date_generated, date_expires, idenf, active FROM tokens WHERE token = ?")

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
		var sess Session
		err := rows.Scan(
			&sess.Date_generated,
			&sess.Date_expires,
			&sess.Idenf,
			&sess.Active,
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



func GetSessions(api router.API, token string) Sessions {
	result := []Session{}
	userid := GetUserIdFromToken(api, token)

	stmt, err := api.Context.Session.Prepare("SELECT date_generated, date_expires, idenf, active FROM tokens WHERE user_id = ? AND active = 1")
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
		var sess Session
		err := rows.Scan(
			&sess.Date_generated,
			&sess.Date_expires,
			&sess.Idenf,
			&sess.Active,
		)

		if err != nil {
			panic(err)
		}
		result = append(result, sess)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return Sessions{Sessions:result, Total: len(result)}
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

