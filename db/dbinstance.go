package db

import (
	"fmt"
	"log"
	"sync"

	"gitlab.com/gowtham-munukutla/vashisht-api/config"

	mgo "gopkg.in/mgo.v2"
)

var (
	dbInstance *mgo.Database
	once       sync.Once
)

func GetDbInstance() *mgo.Database {
	once.Do(func() {
		connectDB()
	})

	return dbInstance
}

func connectDB() {
	config := config.GetAppConfig()
	mongoConfig := config.MongoConfig

	session, err := mgo.Dial(fmt.Sprintf("%s:%d", mongoConfig.Host, mongoConfig.Port))

	if err != nil {
		log.Fatal(err)
	}

	dbInstance = session.DB(mongoConfig.Database)
}
