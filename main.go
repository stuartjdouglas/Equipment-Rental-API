// package main is the main handler
package main

import (
	"fmt"
	"os"
	"github.com/remony/Equipment-Rental-API/core/config"
	"github.com/remony/Equipment-Rental-API/core/server"
	"github.com/remony/Equipment-Rental-API/core/utils/setup"
	"log"
)

const confFile = "./config.json"

var clear map[string]func()

// Starts the server
func main() {
	// Create the server and give it the config values
	settings := config.LoadConfig(confFile, true)
	args := os.Args
	log.Println(args)
	if len(args) > 1 {
		if args[1] == "--setup" {
			fmt.Println("INSTALLING!!!!")
			setup.Start(config.Connection(settings.Production.DbUrl))
		} else if (args[1] == "--dev") {
			start(1, settings);
		}
	} else {
		start(0, settings);
	}
}

func start(mode int, settings config.Config) {
	if (mode == 1) {
		server.Start(settings, config.Connection(settings.Development.DbUrl), 1)
	} else {
		server.Start(settings, config.Connection(settings.Production.DbUrl), 0)
	}
}
