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

// Calculate the number of days from the date given from now
func getDays(date time.Time) int {
	now := time.Now()
	diff := now.Sub(date)
	return int(-diff.Hours() / 24)
}

// Calculate the number of hours from the date given from now
func getHours(date time.Time) int {
	now := time.Now()
	diff := now.Sub(date)
	return int(-diff.Hours())
}

// Task: This is the task we perform an a regular basis;
// Communicate with the database for upcoming rentals
func task(api router.API) {
	log.Println("Performing CRON Task Notification")
	// Get all rentals that are due in 3 days from database.
	Products := database.GetRentalsDueLessThreeDays(api)
	// For each product
	for _, element := range Products.Items {
		// Get the days and hours
		days := getDays(element.Due)
		hours := getHours(element.Due)
		// If it is less than 3 days
		log.Println(days)
		if (days < 3 && hours > 0) {
			// If it is less than 24 hours
			if (hours < 24) {
				// Send a push notification to the user inform it is due in hours
				models.SendNotificationToUser(
					api,
					element.Owner.Username,
					models.Notification{
						Title: "Due in " + strconv.Itoa(hours) + " hours.",
						Message: element.Title + " is due",
					},
				);
			} else {
				// Other wise send a notification to the user informing that it is due in less than 3 days
				models.SendNotificationToUser(
					api,
					element.Owner.Username,
					models.Notification{
						Title: "Due in " + strconv.Itoa(days) + " days.",
						Message: element.Title + " is due soon",
					},
				);
			}
		} else {
			// If the days is negative, product is overdue
			models.SendNotificationToUser(
				api,
				element.Owner.Username,
				models.Notification{
					Title: "Overdue " + strconv.Itoa(-days) + " days.",
					Message: element.Title + " is overdue",
				},
			);
		}

	}
	//log.Println(Products)
}

// Cron: This will call the task every 13 hours
func Cron(api router.API) {
	// Create the custom scheduler
	s := gocron.NewScheduler()
	//s.Every(1).Minute().Do(task, api)
	// Assign the task to be performed every 13 hours
	s.Every(13).Hours().Do(task, api)
	// Star the scheduler
	<-s.Start()
}
