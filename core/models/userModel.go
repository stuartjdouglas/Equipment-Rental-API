
package models

import (
	"../router"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"log"
)

type user struct {
	Username 	string 	`json:"username"`
	Profile 	*Profile `json:"profile"`
}

type Profile struct {
	First_name 		string `json:"first_name"`
	Last_name 		string `json:"last_name"`
	Date_registered string `json:"date_registered"`
}

type fullUser struct {
	Username 	string `json:"username", bson:"username"`
	Email 		string `json:"email", bson:"email"`
	Password 	string `json:"password", bson:"password"`
}

type auth struct {
	Username 	string `json:"username"`
	Token 		string `json:"token"`
}


//noinspection GoUnusedFunction
func GetUser(api router.API, name string) []fullUser {
	c := api.Context.Session.DB("lemon").C("user")
	persons := []fullUser{}
	err := c.Find(bson.M{"name": name}).All(&persons)
	if err != nil {
		return persons
	}
	return persons
}

func CheckIfUserExists(api router.API, username string) bool {
	c := api.Context.Session.DB("lemon").C("users")
	result := &fullUser{}
	err := c.Find(bson.M{"username": username}).One(&result)
	fmt.Println(err)
	if err != nil {
		return false
	}
	return true
}

func RegisterUser(api router.API, username string, password string, email string) bool {
	c := api.Context.Session.DB("lemon").C("users")
	err := api.Context.Err
	err = c.Insert(&fullUser{username, password, email})
	if err != nil {
		log.Fatal("Problem inserting data: ", err)
		return false
	}
	return true
}


//noinspection GoUnusedFunction
func GetUsers(api router.API) []user{
	c := api.Context.Session.DB("lemon").C("users")
	persons := []user{}
	err := c.Find(nil).All(&persons)
	if err != nil {
		return persons
	}
	return persons
}

