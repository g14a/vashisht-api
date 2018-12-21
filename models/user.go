package models

import (
	"errors"
	"log"
	"sync"

	"gopkg.in/mgo.v2/bson"

	"gitlab.com/gowtham-munukutla/vashisht-api/db"
)

// User is the structure of how a User looks
type User struct {
	Name         string `bson:"name" json:"name"`
	Password     string `bson:"pwd" json:"pwd"`
	PhoneNumber  string `bson:"number" json:"number"`
	EmailAddress string `bson:"email" json:"email"`
	CollegeName  string `bson:"college" json:"college"`
	UserID       int    `bson:"userid" json:"userid"`
}

var (
	userCollection = "users"
	userID         int
	usersMu        sync.Mutex
)

func init() {

	dbinstance = db.GetDbInstance()

	size, err = dbinstance.C(userCollection).Count()

	if err != nil {
		log.Fatal(err)
	}
}

func incrementUserID() {
	usersMu.Lock()
	userID++
	usersMu.Unlock()
}

func decrementUserID() {
	usersMu.Lock()
	userID--
	usersMu.Unlock()
}

// AddUser adds a new user after registration into the db
func AddUser(u *User) error {
	incrementUserID()
	u.UserID = userID

	count, err := dbinstance.C(userCollection).Find(bson.M{"email": u.EmailAddress}).Count()

	if count > 0 {
		return errors.New("User already exists, try logging in")
	}

	err = dbinstance.C(userCollection).Insert(&u)

	return err
}

// GetAllUsers returns all the users registered for the fest
func GetAllUsers() ([]User, error) {
	var users []User

	err := dbinstance.C(userCollection).Find(bson.M{}).All(&users)

	return users, err
}

// CheckUserHash checks if the given combination of email and password exists in the db
func CheckUserHash(email, password string) (bool, error) {
	count, err := dbinstance.C(userCollection).Find(bson.M{"email": email, "pwd": password}).Count()

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
