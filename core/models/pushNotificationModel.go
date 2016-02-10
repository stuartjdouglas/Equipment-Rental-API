package models

import (
	"log"
	"fmt"
	"github.com/alexjlockwood/gcm"
	"github.com/remony/Equipment-Rental-API/core/router"
	"github.com/remony/Equipment-Rental-API/core/database"
)


type Notification struct {
	Title string `json:"title"`
	Message string `json:"message"`
}

func PushNotificationAddRegID(api router.API, regid string, deviceType string, token string) database.PNCreate {

	if len(regid) > 0 && len(deviceType) > 0 && len(token) > 0 {
		return database.PushNotificationAddRegID(api, regid, deviceType, token)
	}

	return database.PNCreate{Success:false}
}

func SendNotificationProduct(api router.API, token string, pid string, message Notification) bool {
	if len(pid) > 0 && len(token) > 0 {
		if IsOwner(api, token, pid) {
			data := database.GetPushNotificationProductregid(api, pid)
			sendNotification(data, message)
			return true
		}
	}
	return false
}

func SendNotificationToUser(api router.API, username string, message Notification) bool {
	if (len(username) > 0) {
		data := database.GetPushNotificationUserRegID(api, username)
		sendNotification(data, message)
		return true
	}
	return false
}

func sendNotification(user database.PushNotificationReqID, message Notification) {
	// Create the message to be sent.
	data := map[string]interface{}{
		"title": message.Title,
		"message": message.Message,
	}
	//regIDs := []string{"APA91bEmTcK5SsI4Isj89UUyHtaFJQbRtxrEZxkDHbEebaNqf-qfdu2kgLqjErm1tX1TmtP-v8NeLEha1J2KQJRAP6CDacdkmTzQsOUuEMJwsY156zV6iDC217GUnI8mLp4bRuUSiuFb"}
	msg := gcm.NewMessage(data, user.RequestIDs...)

	// Create a Sender to send the message.
	sender := &gcm.Sender{ApiKey: "AIzaSyB1Y4aWqsxsvKDQl87PLq62ppGSEpC3Ozs"}

	// Send the message and receive the response after at most two retries.
	response, err := sender.Send(msg, 2)
	if err != nil {
		fmt.Println("Failed to send message:", err)
		return
	}
	log.Println(response)
}

func TestNotification() {
	log.Println("test notification")



	// Create the message to be sent.
	data := map[string]interface{}{
		"title": "Hello",
		"message": "Look a real notification",
	}
	regIDs := []string{"APA91bEmTcK5SsI4Isj89UUyHtaFJQbRtxrEZxkDHbEebaNqf-qfdu2kgLqjErm1tX1TmtP-v8NeLEha1J2KQJRAP6CDacdkmTzQsOUuEMJwsY156zV6iDC217GUnI8mLp4bRuUSiuFb"}
	msg := gcm.NewMessage(data, regIDs...)

	// Create a Sender to send the message.
	sender := &gcm.Sender{ApiKey: "AIzaSyB1Y4aWqsxsvKDQl87PLq62ppGSEpC3Ozs"}

	// Send the message and receive the response after at most two retries.
	response, err := sender.Send(msg, 2)
	if err != nil {
		fmt.Println("Failed to send message:", err)
		return
	}
	log.Println(response)
}
