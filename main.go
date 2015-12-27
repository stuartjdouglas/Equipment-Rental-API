// package main is the main handler
package main

import (
	"fmt"
	"os"
	"os/exec"
	"log"
	"github.com/remony/Equipment-Rental-API/core/config"
	"github.com/remony/Equipment-Rental-API/core/server"
	"github.com/remony/Equipment-Rental-API/core/config/database"
)

const confFile = "./config.json"
var clear map[string]func()

// Starts the server
func main() {
	// Create the server and give it the config values
	settings := config.LoadConfig(confFile, true)

	args := os.Args



	if len(args) > 1 {
		if args[1] == "--install" {

			fmt.Println("INSTALLING!!!!")
			log.Println("Installing Bower components")
			cmd:= exec.Command("cls")
			cmd.Stdout = os.Stdout
			cmd.Run()
		}

	} else {

		server.Start(settings, database.Connection(settings.Production.DbUrl))

	}
}
