package setup

import (
	"log"
	"github.com/remony/Equipment-Rental-API/core/config/database"
)

func Start(context database.Context) {
	log.Println("Setting up database")
	setupdb(context)
}


func setupdb(db database.Context) {

}