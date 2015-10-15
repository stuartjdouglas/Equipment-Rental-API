// Package config provides config values provided by the config.json file
package config

import (
	"io/ioutil"
	"encoding/json"
	"log"
)

type Config struct {
	Title string
	Port int
	MongoDb string
}


// Generates a template JSON file, writes it to file and returns the struct
func GenConfig() Config{

	log.Println("Creating config.json")

	config:= Config{}
	//Set the default values
	config.Title = "Default Title"
	config.Port = 3000
	config.MongoDb = "mongodb://remon:lemon@ds042898.mongolab.com:42898/lemon"

	// Parse the json and format in pretty print format
	str, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		// If something goes wrong log it
		log.Fatal(err)
	}

	// Write to file
	e := ioutil.WriteFile("config.json",str , 0644)
	if e != nil {
		// If file fails to write, panic
		panic(e)
	}

	// return the new config struct
	return config
}

// Loads the config from config.json, if not existing create one and return config struct
func LoadConfig() Config{
	// Read in the file
	file, e := ioutil.ReadFile("config.json")
	// Create the empty struct
	config:= Config{}
	if e != nil {
		//file does not exist
		config = GenConfig()
	} else {
		// Parse the json
		json.Unmarshal(file, &config)
	}

	// Return the config struct
	return config
}