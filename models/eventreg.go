package models

import (
	"errors"
	"log"

	"gopkg.in/mgo.v2/bson"
)

// Registration contains the record of user registered for a particular event
type Registration struct {
	EventID int    `bson:"eventid" json:"eventid"`
	UserID  int    `bson:"userid" json:"userid"`
	RegID   string `bson:"regid" json:"regid"`
}

var (
	eventRegCollection = "eventreg"
)

// AddRegistration adds a registration into the db
func AddRegistration(r *Registration) error {

	count, err := dbinstance.C(eventRegCollection).Find(bson.M{"userid": r.UserID, "eventid": r.EventID}).Count()

	if count > 0 {
		return errors.New("User already registered")
	}

	err = dbinstance.C(eventRegCollection).Insert(&r)

	return err
}

// CancelRegistration cancels a registration of an event
func CancelRegistration(r Registration) error {
	err := dbinstance.C(eventRegCollection).Remove(bson.M{"eventid": r.EventID, "userid": r.UserID})

	return err
}

// GetEventsOfUser returns all the events registered by a user
func GetEventsOfUser(userID int) ([]Event, error) {
	var registrations []Registration
	err := dbinstance.C(eventRegCollection).Find(bson.M{"userid": userID}).All(&registrations)

	if err != nil {
		return nil, err
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

func GetUsersForEvent(eventID int) ([]User, error) {
	var registrations []Registration
	err := dbinstance.C(eventRegCollection).Find(bson.M{"eventid": eventID}).All(&registrations)

	if err != nil {
		return nil, err
	}

	users := make([]User, 0)

	for _, reg := range registrations {
		log.Println(reg.RegID)
		userID := reg.UserID

		user, _ := GetUserByID(userID)

		users = append(users, user)
	}

	return users, nil
}

// get all events registered by a user
// get all users registered for an event
