package models

import (
	"log"
	"sync"

	"gitlab.com/gowtham-munukutla/vashisht-api/config"

	"gitlab.com/gowtham-munukutla/vashisht-api/db"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Event struct {
	EventName   string `json:"name"`
	EventID     int    `bson:"id" json:"id"`
	Fee         int    `json:"fee"`
	TeamSize    int    `json:"teamsize"`
	Category    string `json:"category"`
	Day         int    `json:"day"`
	StartTime   int    `json:"start"`
	EndTime     int    `json:"end"`
	Description string `json:"description"`
}

var dbinstance *mgo.Database

var (
	eventsCollection = config.GetAppConfig().MongoConfig.Collections.EventCollectionName
	eventsMu         sync.Mutex
	size             int
	err              error
)

func init() {

	dbinstance = db.GetDbInstance()

	size, err = dbinstance.C(eventsCollection).Count()

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

// AddEvent adds a new event into the db
func AddEvent(newEvent *Event) error {

	incrementSize()
	newEvent.EventID = size

	err := dbinstance.C(eventsCollection).Insert(&newEvent)

	return err
}

// DeleteEvent deletes an events from the db
func DeleteEvent(eventID int) error {
	err := dbinstance.C(eventsCollection).Remove(bson.M{"id": eventID})

	decrementSize()

	return err
}

// UpdateEvent updates an event in the db
func UpdateEvent(updateEvent *Event) error {

	err := dbinstance.C(eventsCollection).Update(bson.M{"id": updateEvent.EventID}, &updateEvent)

	return err
}

// FindEventByID finds an event given its id
func FindEventByID(id int) (Event, error) {
	var event Event
	err := dbinstance.C(eventsCollection).Find(bson.M{"id": id}).One(&event)

	return event, err
}

// FindAllEvents returns all events in the fest db
func FindAllEvents() ([]Event, error) {
	var events []Event
	err := dbinstance.C(eventsCollection).Find(bson.M{}).All(&events)

	return events, err
}
