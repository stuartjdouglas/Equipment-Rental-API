package database

import (
	"github.com/remony/Equipment-Rental-API/core/router"
	"log"
	"strings"
)

type PNCreate struct {
	Success bool `json:"success"`
}

func PushNotificationAddRegID(api router.API, regid string, deviceType string, token string) PNCreate {
	var success PNCreate
	stmt, err := api.Context.Session.Prepare("Call addPushNotificationRegID(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(token, regid, deviceType)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&success.Success,
		)

		if err != nil {
			log.Println("Getting paged results error scanning")
			panic(err)
		}
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return success
}

func parseids(data string) []string {

	return strings.Split(data, ", ");
}

type RequestID struct {
	ID string `json:"id"`
}

type PushNotificationReqID struct {
	Username string `json:"username"`
	RequestIDs []string `json:"requestids"`
	Type string `json:"type"`
}

func GetPushNotificationProductregid(api router.API, pid string) PushNotificationReqID {
	var content PushNotificationReqID
	stmt, err := api.Context.Session.Prepare("Call GetPushNotificationIDsOfProduct(?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(pid)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var ids string
		err := rows.Scan(
			&content.Username,
			&ids,
			&content.Type,
		)

		content.RequestIDs = parseids(ids)

		if err != nil {
			log.Println("Getting paged results error scanning")
			panic(err)
		}
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return content
}
func GetPushNotificationUserRegID(api router.API, username string) PushNotificationReqID {
	var content PushNotificationReqID
	stmt, err := api.Context.Session.Prepare("Call GetPushNotificationIDsOfUser(?)")
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
		var ids string
		err := rows.Scan(
			&content.Username,
			&ids,
			&content.Type,
		)

		content.RequestIDs = parseids(ids)

		if err != nil {
			log.Println("Getting paged results error scanning")
			panic(err)
		}
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return content
}
