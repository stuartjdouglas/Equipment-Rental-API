package config

import (
	"log"
 "database/sql"
_ "github.com/go-sql-driver/mysql"
	"fmt"
)

type Context struct {
	Session *sql.DB
	Err error
}

type Person struct {
	Name string
	Email string
}


func Connection(url string) Context{
	fmt.Println(url);
	db, err := sql.Open("mysql", url)
	if err != nil {
		fmt.Println(err)
		panic(err)
//		return Context{Session:nil, Err: err}
	} else {
		log.Printf("Connected to " + url)
	}
	context := Context{}
	context.Session = db
	context.Err = err
	return context
}