// Package config provides config values provided by the config.json file
package config

import (
	"io/ioutil"
	"encoding/json"
	"log"
)

type Config struct {
	Development Properties `json:"development"`
	Production Properties `json:"production"`
}

type Properties struct {
	Title string `json:"title"`
	Port int `json:"port"`
	DbUrl string `json:"dburl"`
}


// Generates a template JSON file, writes it to file and returns the struct
func GenConfig(path string) Config{

	log.Println("Creating config.json")

	config:= Config{}
	//Set the default values
	config.Development.Title = "Default Title"
	config.Development.Port = 3000
	config.Development.DbUrl = "mongodb://remon:lemon@ds042898.mongolab.com:42898/lemon"

	config.Production.Title = "Default Title"
	config.Production.Port = 3000
	config.Production.DbUrl = "mongodb://remon:lemon@ds042898.mongolab.com:42898/lemon"

	// Parse the json and format in pretty print format
	str, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		// If something goes wrong log it
		log.Fatal(err)
	}

	// Write to file
	e := ioutil.WriteFile(path, str, 0644)
	if e != nil {
		// If file fails to write, panic
		panic(e)
	}

	// return the new config struct
	return config
}



// Loads the config from config.json, if not existing create one and return config struct
func LoadConfig(path string, devMode bool) Properties{
	// Read in the file
	file, e := ioutil.ReadFile(path)
	// Create the empty struct
	config:= Config{}
	if e != nil {
		//file does not exist
		config = GenConfig(path)
	} else {
		// Parse the json
		json.Unmarshal(file, &config)
	}

//	Create the properties object which will contain the used settings of mode defined by user
	var settings Properties

//	If the operator has defined to use devMode use development values otherwise set to Production values
	if devMode {
		settings.Title = config.Development.Title
		settings.DbUrl = config.Development.DbUrl
		settings.Port  = config.Development.Port
	} else {
//		Always fall back to production values
		settings.Title = config.Production.Title
		settings.DbUrl = config.Production.DbUrl
		settings.Port  = config.Production.Port
	}

	// Return the config struct
	return settings
}