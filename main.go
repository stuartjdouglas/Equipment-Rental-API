// package main is the main handler
package main

import (
	"./core/config"
	"./core/server"
	"fmt"
)

// Starts the server
func main() {
	// Create the server and give it the config values
	fmt.Println(config.LoadConfig().Title);
	settings := config.LoadConfig()
	server.Start(settings, config.Connection(settings.MongoDb));
}
