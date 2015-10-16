
package models

import (
	"../router"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type person struct {
	Name 	string `json:"name"`
	Email 	string `json:"email"`
}

type user struct {
	Username 	string 	`json:"username"`
	Profile 	*Profile `json:"profile"`
}

type Profile struct {
	First_name 		string `json:"first_name"`
	Last_name 		string `json:"last_name"`
	Date_registered string `json:"date_registered"`
}

type fulluser struct {
	Username 	string `json:"username", bson:"username"`
	First_name 	string `json:"first_name"`
	Last_name 	string `json:"last_name"`
	Email 		string `json:"email"`
	Password 	string `json:"password"`
}

type auth struct {
	Username 	string `json:"username"`
	Token 		string `json:"token"`
}


//noinspection GoUnusedFunction
func GetUser(api router.API, name string) []person {
	c := api.Context.Session.DB("lemon").C("user")

	persons := []person{}
	err := c.Find(bson.M{"name": name}).All(&persons)
	if err != nil {
		return persons
	}

	return persons
}



func registerUser(api router.API) {
	c := api.Context.Session.DB("lemon").C("users")
	err := api.Context.Err
	err = c.Insert(&fulluser{"Remon", "Remon", "Yamano", "remonasebi@gmail.com", "lemon"})
	if err != nil {
		log.Fatal("Problem inserting data: ", err)
		return
	}

}


//noinspection GoUnusedFunction
func GetUsers(api router.API) []person{
	c := api.Context.Session.DB("lemon").C("user")
//	err:= api.Context.Err

	persons := []person{}
	err := c.Find(nil).All(&persons)
	if err != nil {
		return persons
	}

	return persons



}

