package models

import (
	"errors"

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
		return errors.New(`{r.UserID} already registered`)
	}

	err = dbinstance.C(eventRegCollection).Insert(&r)

	return err
}

// CancelRegistration cancels a registration of an event
func CancelRegistration(r Registration) error {
	err := dbinstance.C(eventRegCollection).Remove(bson.M{"eventid": r.EventID, "userid": r.UserID})

	return err
}

// users/{id}/events/{eventid}/register
