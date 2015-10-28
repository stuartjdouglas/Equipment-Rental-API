// package main is the main handler
package main

import (
	"./core/config"
	"./core/config/database"
	"./core/server"
	"fmt"
	"strconv"
)

const confFile = "./config.json"

// Starts the server
func main() {
	// Create the server and give it the config values
	settings := config.LoadConfig(confFile, true)

	fmt.Println(">>>>>>>>>" + settings.Title + ": " + settings.DbUrl + ": " + strconv.Itoa(settings.Port))

	server.Start(settings, database.Connection(settings.DbUrl))
}
