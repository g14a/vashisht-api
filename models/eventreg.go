package models

import (
	"errors"

	"github.com/mongodb/mongo-go-driver/bson"
	"gitlab.com/gowtham-munukutla/vashisht-api/config"
	"gitlab.com/gowtham-munukutla/vashisht-api/db"
)

// Registration contains the record of user registered for a particular event
type Registration struct {
	EventID int    `bson:"eventid" json:"eventid"`
	UserID  int    `bson:"userid" json:"userid"`
	RegID   string `bson:"regid" json:"regid"`
}

var (
	eventRegCollectionName = config.GetAppConfig().MongoConfig.Collections.RegistrationCollectionName
)

// AddRegistration adds a registration into the db
func AddRegistration(r *Registration) error {
	eventRegCollection, ctx := db.GetMongoCollectionWithContext(eventRegCollectionName)
	count, err := eventRegCollection.Count(ctx, bson.M{"userid": r.UserID, "eventid": r.EventID})
	if count > 0 {
		return errors.New("User already registered")
	}
	_, err = eventRegCollection.InsertOne(ctx, r)
	return err
}

// CancelRegistration cancels a registration of an event
func CancelRegistration(r Registration) error {
	eventRegCollection, ctx := db.GetMongoCollectionWithContext(eventRegCollectionName)
	_, err := eventRegCollection.DeleteOne(ctx, bson.M{"eventid": r.EventID, "userid": r.UserID})
	return err
}

// GetEventsOfUser returns all the events registered by a user
func GetEventsOfUser(userID int) ([]Event, error) {
	eventRegCollection, ctx := db.GetMongoCollectionWithContext(eventRegCollectionName)
	var registrations []Registration
	cursor, err := eventRegCollection.Find(ctx, bson.M{"userid": userID})
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var registration Registration
		if err := cursor.Decode(&registration); err != nil {
			return nil, err
		} else {
			registrations = append(registrations, registration)
		}
	}

	events := make([]Event, 0)

	for _, reg := range registrations {
		eventID := reg.EventID
		event, err := FindEventByID(eventID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

// GetUsersForEvent returns all the users registered for a single event
func GetUsersForEvent(eventID int) ([]User, error) {
	eventRegCollection, ctx := db.GetMongoCollectionWithContext(eventRegCollectionName)
	var registrations []Registration
	cursor, err := eventRegCollection.Find(ctx, bson.M{"eventid": eventID})
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var registration Registration
		if err := cursor.Decode(&registration); err != nil {
			return nil, err
		}
		registrations = append(registrations, registration)
	}

	users := make([]User, 0)
	for _, reg := range registrations {
		userID := reg.UserID
		user, _ := GetUserByID(userID)
		users = append(users, user)
	}

	return users, nil
}

// CheckIfUserRegisteredForEvent checks if the user registered for a particular event
func CheckIfUserRegisteredForEvent(userID, eventID int) (bool, error) {
	eventRegCollection, ctx := db.GetMongoCollectionWithContext(eventRegCollectionName)
	count, err := eventRegCollection.Count(ctx, bson.M{"userid": userID, "eventid": eventID})
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// CheckIfUserRegisteredForEventByMongoID checks if the user registered for the event using Mongo Object ID
func CheckIfUserRegisteredForEventByMongoID(mongoID string, eventID int) (bool, error) {
	eventRegCollection, ctx := db.GetMongoCollectionWithContext(eventRegCollectionName)
	count, err := eventRegCollection.Count(ctx, bson.M{"_id": mongoID, "eventid": eventID})
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
