package structs

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

func (e Event) AddEvent(newEvent Event) error {
	err := db.C(COLLECTION).Insert(&newEvent)

	return err
}

func (e Event) DeleteEvent(newEvent Event) error {
	err := db.C(COLLECTION).Remove(&newEvent)

	return err
}

func (e Event) UpdateEvent(updateEvent Event) error {
	err := db.C(COLLECTION).Update(updateEvent.EventId, &updateEvent)

	return err
}

func (e Event) FindEventById(id string) (Event, error) {
	var event Event
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&event)

	return event, err
}
