package models

import (
	"fmt"
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
	database   = "vashisht"
	collection = "events"
	eventsMu   sync.Mutex
	size       int
	err        error
)

func init() {

	dbinstance = db.GetDbInstance()

	size, err = dbinstance.C(collection).Count()

	if err != nil {
		log.Fatal(err)
	}
}

func incrementSize() {
	eventsMu.Lock()
	size++
	eventsMu.Unlock()
}

func decrementSize() {
	eventsMu.Lock()
	size--
	eventsMu.Unlock()
}

func AddEvent(newEvent *Event) error {

	incrementSize()
	newEvent.EventId = size

	err := dbinstance.C(collection).Insert(&newEvent)

	return err
}

func DeleteEvent(eventID int) error {
	err := dbinstance.C(collection).Remove(bson.M{"id": eventID})

	decrementSize()

	return err
}

func UpdateEvent(updateEvent *Event) error {
	fmt.Println(updateEvent.EventId)

	err := dbinstance.C(collection).Update(bson.M{"id": updateEvent.EventId}, &updateEvent)

	return err
}

func FindEventById(id string) (Event, error) {
	var event Event
	err := dbinstance.C(collection).Find(bson.M{"id": id}).One(&event)

	return event, err
}

func FindAllEvents() ([]Event, error) {
	var events []Event
	err := dbinstance.C(collection).Find(bson.M{}).All(&events)

	return events, err
}
