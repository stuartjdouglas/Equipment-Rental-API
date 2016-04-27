package models

import (
	"github.com/remony/Equipment-Rental-API/core/router"
	"github.com/remony/Equipment-Rental-API/core/database"
	"github.com/minimaxir/big-list-of-naughty-strings/naughtystrings"
	"gitlab.com/remon/lemon-swear-detector"
	"github.com/remony/Equipment-Rental-API/core/utils/email"
	"strconv"
)

func checkIfNaughtyWord(word string) bool {
	for _, element := range naughtystrings.Unencoded() {
		if word == element {
			return true
		}
	}
	return false;
}

func EditComment(api router.API, token string, cid string, comment string, rating int) database.Comment {
	if (IsSessionValid(api, token)) {
		if (rating < 0) {
			rating = 0;
		} else if (rating > 5) {
			rating = 5
		}

		cid := database.EditComment(api, token, cid, comment, rating)
		return database.GetComment(api, cid)
	}

	return database.Comment{}
}

func AddComment(api router.API, token string, pid string, comment string, rating int) database.Comment {
	// Rating cannot be more than 5 or less than 0

	if (rating > 5) {
		rating = 5
	} else if (rating < 0) {
		rating = 0
	}

	if IsSessionValid(api, token) {
		//if (checkIfNaughtyWord(comment))
		if lemon_swear_detector.CheckSentence(comment) {
			cid := database.AddComment(api, token, pid, comment, true, rating)
			username := database.GetUserNameFromToken(api, token)
			userdata := database.GetUserDetails(api, username)
			owner := database.GetProductFromID(api, pid, token)
			ownerDetails := database.GetUserDetails(api, owner.Items[0].Owner.Username)
			//SendEmail(api router.API, sender string, receipt string, subject string, body string)
			email.SendEmail(
				api,
				ownerDetails.Username,
				ownerDetails.Email,
				"Someone has reviewed " + owner.Items[0].Product_name,
				userdata.Username + " wrote review on owner.Items[0].Product_name: \n" + comment + "\n\nRating: " + strconv.Itoa(rating) + "/5",
			)
			return database.GetComment(api, cid)
		} else {
			cid := database.AddComment(api, token, pid, comment, false, rating)
			username := database.GetUserNameFromToken(api, token)
			userdata := database.GetUserDetails(api, username)
			owner := database.GetProductFromID(api, pid, token)

			ownerDetails := database.GetUserDetails(api, owner.Items[0].Owner.Username)
			email.SendEmail(api,
				ownerDetails.Username,
				ownerDetails.Email,
				"Someone has reviewed " + owner.Items[0].Product_name,
				userdata.Username + " wrote review on " + owner.Items[0].Product_name + ": <br><br>" + comment + "<br><br><b>Rating</b>: " + strconv.Itoa(rating) + "/5",
			)
			return database.GetComment(api, cid)
		}
	}
	return database.Comment{}
}

func DeleteComment(api router.API, pid string, cid string, token string) {
	if IsSessionValid(api, token) {
		database.DeleteComment(api, pid, cid, token)
	}
}

func EnableComments(api router.API, pid string, token string) {
	database.EnableComments(api, pid)
}
func DisableComments(api router.API, pid string, token string) {
	database.DisableComments(api, pid)
}

func ApproveComment(api router.API, pid string, cid string) {
	database.ApproveComment(api, pid, cid)
}
