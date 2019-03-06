package models

import (
	"gopkg.in/mgo.v2/bson"
)

// User is the model to be used for storing user's data.
type User struct {
	Name   string        `json:"name" bson:"name"`
	Gender string        `json:"gender" bson:"gender"`
	Age    int           `json:"age" bson:"age"`
	Id     bson.ObjectId `json:"id" bson:"_id"`
}

// Id was of type string before
