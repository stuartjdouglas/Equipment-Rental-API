package config

import (

//	"fmt"
//	"os"
	"gopkg.in/mgo.v2"
//	"log"
//	"gopkg.in/mgo.v2/bson"
	"log"
	"gopkg.in/mgo.v2/bson"
	"fmt"
)

type Context struct {
	Session *mgo.Session
	Err error
}

type Person struct {
	Name string
	Email string
}


func Connection(url string) Context{
//	If url = nil return nil session

	sess, e := mgo.Dial(url)

//	fmt.Printf("MONGO URL: " + url)

	if e != nil {
		return Context{Session:nil, Err: e}
	} else {
		log.Printf("Connected to " + url)
	}

	context := Context{}
	context.Session = sess
	context.Err = e

	return context
}

func PersonGet(context Context) {
//	sess, err := connection();


	collection := context.Session.DB("lemon").C("remon")
	err:= context.Err

	err = collection.Insert(&Person{"Stefan Klaste", "klaste@posteo.de"},
		&Person{"Nishant Modak", "modak.nishant@gmail.com"},
		&Person{"Prathamesh Sonpatki", "csonpatki@gmail.com"},
		&Person{"murtuza kutub", "murtuzafirst@gmail.com"},
		&Person{"aniket joshi", "joshianiket22@gmail.com"},
		&Person{"Michael de Silva", "michael@mwdesilva.com"},
		&Person{"Alejandro Cespedes Vicente", "cesal_vizar@hotmail.com"})
	if err != nil {
		log.Fatal("Problem inserting data: ", err)
		return
	}

	result := Person{}
	err = collection.Find(bson.M{"name": "Prathamesh Sonpatki"}).One(&result)
	if err != nil {
		log.Fatal("Error finding record: ", err)
		return
	}

	fmt.Println("Email Id:", result.Email)
}