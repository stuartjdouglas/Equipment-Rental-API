// package main is the main handler
package main

import (
	"./core/config"
	"./core/server"
	"fmt"
	"strconv"
)

const CONF_FILE = "./config.json"

// Starts the server
func main() {
	// Create the server and give it the config values
	settings := config.LoadConfig(CONF_FILE, true)

	fmt.Println(settings.Title + ": " + settings.DbUrl + ": " + strconv.Itoa(settings.Port))

	server.Start(settings, config.Connection(settings.DbUrl));
}
