// package main is the main handler
package main

import (
	"./core/config"
	"./core/config/database"
	"./core/server"
	"fmt"
	"strconv"
	"os"
	"os/exec"
	"log"
)

const confFile = "./config.json"
var clear map[string]func()

// Starts the server
func main() {
	// Create the server and give it the config values
	settings := config.LoadConfig(confFile, true)

	fmt.Println(">>>>>>>>>" + settings.Title + ": " + settings.DbUrl + ": " + strconv.Itoa(settings.Port))
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
		server.Start(settings, database.Connection(settings.DbUrl))

	}
}
