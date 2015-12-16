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
	Type 	string 	`json:"type"`
	Title 	string 	`json:"title"`
	Port 	int 	`json:"port"`
	DbUrl 	string 	`json:"dburl"`
}


// Generates a template JSON file, writes it to file and returns the struct
func GenConfig(path string) Config{
	log.Println("Configuration file is missing; generating....")
	// Create default configuration struct
	config:= Config{
		Development: Properties{
			Title: "Default Title",
			DbUrl: "root:l3mon@tcp(lemondev.xyz:3306)/honoursproject?parseTime=true?clientMultiResults=true?clientMultiResults=true",
			Port: 3000,
		},
		Production: Properties{
			Title: "Default Title",
			DbUrl: "root:l3mon@tcp(lemondev.xyz:3306)/honoursproject?parseTime=true?clientMultiResults=true?clientMultiResults=true",
			Port: 80,
		},
	}

	// Parse the json and format in pretty print format
	str, err := json.MarshalIndent(config, "", "    ")
	if err != nil {
		// If something goes wrong log it
		log.Fatal(err)
	}

	// Write to file
	e := ioutil.WriteFile(path, str, 0644)
	if e != nil {
		log.Fatal("Unable to write to filesystem")
	}
	// Generated config will return even if write failed
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

	//	If the operator has defined to use devMode use development values otherwise set to Production values
	if devMode {
		return Properties{
			Type:	"Development",
			Title: 	config.Development.Title,
			DbUrl: 	config.Development.DbUrl,
			Port: 	config.Development.Port,
		}
	}

	//	Always fall back to production values
	return Properties{
		Type:	"Production",
		Title: config.Development.Title,
		DbUrl: config.Development.DbUrl,
		Port: config.Development.Port,
	}
}