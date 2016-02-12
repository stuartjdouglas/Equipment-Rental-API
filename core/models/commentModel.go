package models

import (
	"github.com/remony/Equipment-Rental-API/core/router"
	"github.com/remony/Equipment-Rental-API/core/database"
	"github.com/minimaxir/big-list-of-naughty-strings/naughtystrings"
	"log"
	"gitlab.com/remon/lemon-swear-detector"
)


func checkIfNaughtyWord(word string) bool {
	log.Println("using word " + word)
	for _, element := range naughtystrings.Unencoded() {
		//log.Println(element)
		if word == element {
			return true
		}
	}
	return false;
}

func AddComment(api router.API, token string, pid string, comment string) bool {
	if IsSessionValid(api, token) {
		//if (checkIfNaughtyWord(comment))
		if lemon_swear_detector.CheckSentence(comment) {
			//log.Println(checkIfNaughtyWord("0x0"))
			database.AddComment(api, token, pid, comment, true)
			return true
		} else {
			database.AddComment(api, token, pid, comment, false)
			return true
		}
	}
	return false
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
