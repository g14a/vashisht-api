package models

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Event struct {
	EventName string `json:"name"`
	EventId   string `json:"id"`
	Fee       int    `json:"fee"`
	TeamSize  int    `json:"teamsize"`
	Category  string `json:"category"`
	Day       int    `json:"day"`
	StartTime int    `json:"start"`
	EndTime   int    `json:"end"`
}

var db *mgo.Database

var (
	DATABASE   = "vashisht"
	COLLECTION = "events"
)

func AddEvent(newEvent Event, db *mgo.Database) error {
	err := db.C(COLLECTION).Insert(&newEvent)

	return err
}

func DeleteEvent(newEvent Event, db *mgo.Database) error {
	err := db.C(COLLECTION).Remove(&newEvent)

	return err
}

func UpdateEvent(updateEvent Event, db *mgo.Database) error {
	err := db.C(COLLECTION).Update(updateEvent.EventId, &updateEvent)

	return err
}

func FindEventById(id string, db *mgo.Database) (Event, error) {
	var event Event
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&event)

	return event, err
}
