package models

import (
	"strings"
	"log"
	"github.com/remony/Equipment-Rental-API/core/database"
	"github.com/remony/Equipment-Rental-API/core/router"
)

type Tag struct {
	Title string `json:"title"`
}

func removeWhitespace(value string) string {
	return strings.TrimSpace(value)
}

func parseJSArrayTags(json string) []Tag {
	var tags []Tag;
	new := strings.Split(json, ",")
	for i := 0; i < len(new); i++ {
		var tag = Tag{
			Title:  removeWhitespace(new[i]),
		}
		tags = append(tags, tag)
	}
	return tags;
}

func UploadTags(api router.API, pid string, data string) bool {
	tags := parseJSArrayTags(data)
	log.Println(len(tags))
	for i := 0; i < len(tags); i++ {
		if len(tags[i].Title) > 0 {

			database.AddTagToProduct(api, pid, tags[i].Title)
		}
	}
	return true
}