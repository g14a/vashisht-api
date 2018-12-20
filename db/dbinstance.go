package db

import (
	"log"
	"sync"

	mgo "gopkg.in/mgo.v2"
)

var (
	dbInstance *mgo.Database
	once       sync.Once
	dbname     = "vashisht"
)

func GetDbInstance() *mgo.Database {
	once.Do(func() {
		connectDB()
	})

	return dbInstance
}

func connectDB() {
	session, err := mgo.Dial("localhost:27017")

	if err != nil {
		log.Fatal(err)
	}

	dbInstance = session.DB(dbname)
}
