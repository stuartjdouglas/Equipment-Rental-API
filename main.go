// package main is the main handler
package main

import (
	"./core/config"
	"./core/server"
	"fmt"
)

const CONF_FILE = "./config.json"

// Starts the server
func main() {
	// Create the server and give it the config values
	fmt.Println(config.LoadConfig(CONF_FILE).Title);
	settings := config.LoadConfig(CONF_FILE)
	server.Start(settings, config.Connection(settings.MongoDb));
}
