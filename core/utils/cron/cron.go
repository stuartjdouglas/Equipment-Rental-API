package cron

import (
	"log"
	"github.com/jasonlvhit/gocron"
	"github.com/remony/Equipment-Rental-API/core/router"
	"github.com/remony/Equipment-Rental-API/core/database"
	"github.com/remony/Equipment-Rental-API/core/models"
	"time"
	"strconv"
)

func getDays(date time.Time) int {
	now := time.Now()

	diff := now.Sub(date)

	return int(-diff.Hours() / 24)
}

func getHours(date time.Time) int {
	now := time.Now()

	diff := now.Sub(date)

	return int(-diff.Hours())
}


func task(api router.API) {
	log.Println("Performing CRON Task Notification")
	// Get all rentals that are due in 3 days from database.
	Products := database.GetRentalsDueLessThreeDays(api)

	for _, element := range Products.Items {
		days := getDays(element.Due)
		hours := getHours(element.Due)
		if (days < 3) {
			if (hours < 24) {
				models.SendNotificationToUser(
					api,
					element.Owner.Username,
					models.Notification{
						Title: "Due in " + strconv.Itoa(hours) + " hours.",
						Message: element.Title + " is due",
					},
				);
			} else {
				models.SendNotificationToUser(
					api,
					element.Owner.Username,
					models.Notification{
						Title: "Due in " + strconv.Itoa(days) + " days.",
						Message: element.Title + " is due soon",
					},
				);
			}
		}

	}
	//log.Println(Products)
}

func Cron(api router.API) {
	s := gocron.NewScheduler()
	//s.Every(1).Minute().Do(task, api)
	s.Every(13).Hours().Do(task, api)
	<-s.Start()
}
