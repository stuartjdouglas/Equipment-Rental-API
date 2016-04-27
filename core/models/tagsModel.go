package models

import (
	"strings"
	"github.com/remony/Equipment-Rental-API/core/database"
	"github.com/remony/Equipment-Rental-API/core/router"
	"strconv"
	"log"
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
	for i := 0; i < len(tags); i++ {
		if len(tags[i].Title) > 0 {

			database.AddTagToProduct(api, pid, tags[i].Title)
		}
	}
	return true
}

func RemoveTag(api router.API, pid string, tag string, token string) bool {
	if IsOwner(api, token, pid) {
		database.RemoveTag(api, pid, tag)
		return true
	}
	return false
}

func AddTag(api router.API, pid string, tag string, token string) bool {
	if IsOwner(api, token, pid) {
		UploadTags(api, pid, tag)
		return true
	}
	return false
}

func GetTags(api router.API, pid string) []database.Tag {
	return database.GetTags(api, pid)
}

func parseStringToInt(value string) int {
	result, err := strconv.Atoi(value)
	if err != nil {
		log.Println(err)
	}
	return result
}

func parseCount(count int) int {
	if count < 6 {
		return 6
	}
	return count
}

func GetProductsOfTag(api router.API, tag string, start string, count string) database.Items {
	return database.GetProductsWithTag(api, tag, parseStringToInt(start), parseCount(parseStringToInt(count)))
}

func GetTagsMostUsed(api router.API, step int, count int, token string, order bool) []database.MostUsedTag {
	return database.GetTagsMostUsed(api, step, count, token, order)
}