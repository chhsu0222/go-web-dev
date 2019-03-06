package config

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

// Books is the collection stores all the book documents.
var Books *mgo.Collection

func init() {
	// get a mongo session
	s, err := mgo.Dial("mongodb://localhost/bookstore")
	if err != nil {
		panic(err)
	}

	if err = s.Ping(); err != nil {
		panic(err)
	}

	Books = s.DB("bookstore").C("books")

	fmt.Println("You connected to your mongo database.")
}
