package models

import (
	"sync"

	"github.com/mongodb/mongo-go-driver/bson"
	"gitlab.com/gowtham-munukutla/vashisht-api/config"

	"gitlab.com/gowtham-munukutla/vashisht-api/db"
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
	ImageUrl    string `json:"url"`
}

var (
	eventsCollection = config.GetAppConfig().MongoConfig.Collections.EventCollectionName
	eventsMutex      = &sync.Mutex{}
)

// AddEvent adds a new event into the db
func AddEvent(newEvent *Event) error {
	eventCollection, ctx := db.GetMongoCollectionWithContext(eventsCollection)
	count, err := eventCollection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return err
	}
	eventsMutex.Lock()
	count = count + 1
	newEvent.EventID = int(count)
	eventsMutex.Unlock()
	_, err = eventCollection.InsertOne(ctx, newEvent)
	return err
}

// DeleteEvent deletes an events from the db
func DeleteEvent(eventID int) error {
	eventCollection, ctx := db.GetMongoCollectionWithContext(eventsCollection)
	_, err := eventCollection.DeleteOne(ctx, bson.M{"id": eventID})
	return err
}

// UpdateEvent updates an event in the db
func UpdateEvent(updateEvent *Event) error {
	eventCollection, ctx := db.GetMongoCollectionWithContext(eventsCollection)
	_, err := eventCollection.ReplaceOne(ctx, bson.M{"id": updateEvent.EventID}, &updateEvent)
	return err
}

// FindEventByID finds an event given its id
func FindEventByID(id int) (Event, error) {
	eventCollection, ctx := db.GetMongoCollectionWithContext(eventsCollection)
	var event Event
	err := eventCollection.FindOne(ctx, bson.M{"id": id}).Decode(&event)
	return event, err
}

// FindAllEvents returns all events in the fest db
func FindAllEvents() ([]Event, error) {
	eventCollection, ctx := db.GetMongoCollectionWithContext(eventsCollection)
	var events []Event
	cursor, err := eventCollection.Find(ctx, bson.M{})
	for cursor.Next(ctx) {
		var event Event
		if err := cursor.Decode(&event); err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, err
}
