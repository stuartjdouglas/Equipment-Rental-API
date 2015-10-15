// package main is the main handler
package main

import (
	"./core/server"
	"./core/config"
	"fmt"
)

// Starts the server
func main() {
	// Create the server and give it the config values
	fmt.Println(config.LoadConfig().Title);
	settings := config.LoadConfig()
	server.Start(settings, config.Connection(settings.MongoDb));
}
