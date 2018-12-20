package handlers

import (
	"log"

	mgo "gopkg.in/mgo.v2"
)

var db *mgo.Database

var (
	DATABASE   = "vashisht"
	COLLECTION = "events"
)

func Connect() {
	session, err := mgo.Dial("localhost:27017")

	if err != nil {
		log.Fatal(err)
	}

	db = session.DB(DATABASE)
}
