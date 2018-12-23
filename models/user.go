package models

import (
	"errors"
	"log"
	"sync"

	"gopkg.in/mgo.v2/bson"
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

// GetUserByID returns the users given the userid
func GetUserByID(userID int) (User, error) {
	var user User
	log.Println(userID)
	err := dbinstance.C(userCollection).Find(bson.M{"userid": userID}).One(&user)

	return user, err
}

// CheckUserHash checks if the given combination of email and password exists in the db
func CheckUserHash(email, password string) (bool, error) {
	count, err := dbinstance.C(userCollection).Find(bson.M{"email": email, "pwd": password}).Count()

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func init() {
	userID = countUsers()
}

func countUsers() int {
	count, _ := dbinstance.C(userCollection).Count()

	return count
}
