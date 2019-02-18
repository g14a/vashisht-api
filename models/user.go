package models

import (
	"errors"
	"log"
	"sync"

	"gitlab.com/gowtham-munukutla/vashisht-api/idgen"

	"github.com/mongodb/mongo-go-driver/bson"
	"gitlab.com/gowtham-munukutla/vashisht-api/config"
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
	SamID        string `bson:"samid" json:"samid"`
}

var (
	usersCollection = config.GetAppConfig().MongoConfig.Collections.UserCollectionName
	usersMutex      = &sync.Mutex{}
)

// AddUser adds a new user after registration into the db
func AddUser(u *User) error {
	log.Println("Will add user")
	usersCollection, ctx := db.GetMongoCollectionWithContext(usersCollection)

	duplicateUsers, err := usersCollection.Count(ctx, bson.M{"email": u.EmailAddress})

	if duplicateUsers > 0 {
		return errors.New("user with the given email address already exists")
	}

	count, err := usersCollection.CountDocuments(ctx, bson.M{})
	eventsMutex.Lock()
	count = count + 1
	u.UserID = int(count)
	u.SamID = idgen.GenerateSamID(u.UserID)
	eventsMutex.Unlock()
	_, err = usersCollection.InsertOne(ctx, u)

	return err
}

// GetAllUsers returns all the users registered for the fest
func GetAllUsers() ([]User, error) {
	usersCollection, ctx := db.GetMongoCollectionWithContext(usersCollection)
	var users []User
	cursor, err := usersCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var user User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, err
}

// GetUserByID returns the users given the userid
func GetUserByID(userID int) (User, error) {
	usersCollection, ctx := db.GetMongoCollectionWithContext(usersCollection)
	var user User
	err := usersCollection.FindOne(ctx, bson.M{"userid": userID}).Decode(&user)
	return user, err
}

// CheckUserHash checks if the given combination of email and password exists in the db
func CheckUserHash(email, password string) bool {
	usersCollection, ctx := db.GetMongoCollectionWithContext(usersCollection)
	count, err := usersCollection.Count(ctx, bson.M{"email": email, "pwd": password})
	if err != nil {
		return false
	}
	return count > 0
}

// Login checks if a user combination is present in the database. If yes, returns the user object, or else returns a nil
func Login(email, password string) (interface{}, error) {
	usersCollection, ctx := db.GetMongoCollectionWithContext(usersCollection)

	var err error
	if CheckUserHash(email, password) {
		var user User

		err = usersCollection.FindOne(ctx, bson.M{"email": email, "pwd": password}).Decode(&user)

		return &user, nil
	}

	return nil, err
}
