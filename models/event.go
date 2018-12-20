package models

import (
	"log"
	"sync"

	"github.com/vashisht-api/db"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Event struct {
	EventName string `json:"name"`
	EventId   int    `bson:"id" json:"id"`
	Fee       int    `json:"fee"`
	TeamSize  int    `json:"teamsize"`
	Category  string `json:"category"`
	Day       int    `json:"day"`
	StartTime int    `json:"start"`
	EndTime   int    `json:"end"`
}

var dbinstance *mgo.Database

var (
	DATABASE   = "vashisht"
	COLLECTION = "events"
	mu         sync.Mutex
	size       int
	err        error
)

func init() {

	dbinstance = db.GetDbInstance()

	size, err = dbinstance.C(COLLECTION).Count()
	if err != nil {
		log.Fatal(err)
	}
}

func incrementSize() {
	mu.Lock()
	size++
	mu.Unlock()
}

func decrementSize() {
	mu.Lock()
	size--
	mu.Unlock()
}

func AddEvent(newEvent Event, db *mgo.Database) error {

	incrementSize()
	newEvent.EventId = size

	err := db.C(COLLECTION).Insert(&newEvent)

	return err
}

func DeleteEvent(newEvent Event, db *mgo.Database) error {
	err := db.C(COLLECTION).Remove(&newEvent)

	decrementSize()
	return err
}

func UpdateEvent(updateEvent Event, db *mgo.Database) error {
	err := db.C(COLLECTION).Update(updateEvent.EventId, &updateEvent)

	return err
}

func FindEventById(id string, db *mgo.Database) (Event, error) {
	var event Event
	err := db.C(COLLECTION).Find(bson.M{"id": id}).One(&event)

	return event, err
}

func FindAllEvents(db *mgo.Database) ([]Event, error) {
	var events []Event
	err := db.C(COLLECTION).Find(bson.M{}).All(&events)

	return events, err
}
