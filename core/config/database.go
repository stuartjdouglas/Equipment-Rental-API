package config

import (
	"log"
	"database/sql"
	_ "github.com/cxflag203/mysql"
	"fmt"
)

type Context struct {
	Session *sql.DB
	Err     error
	Debug   bool
}

type Person struct {
	Name  string
	Email string
}

func Connection(url string) Context {
	//	root:l3mon@tcp(lemondev.xyz:3306)/honoursproject?parseTime=true
	db, err := sql.Open("mysql", url)
	if err != nil {
		fmt.Println(err)
		panic(err)
		//		return Context{Session:nil, Err: err}
	} else {
		log.Println("Connected to " + url)
		//log.Println("Connected to maria");
	}
	context := Context{}
	context.Session = db
	context.Err = err
	return context
}