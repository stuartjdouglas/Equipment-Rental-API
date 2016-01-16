package setup

import (
	"log"
	"github.com/remony/Equipment-Rental-API/core/config"
)

func Start(context config.Context) {
	log.Println("Setting up database")
	setupdb(context)
}


func setupdb(db config.Context) {

}