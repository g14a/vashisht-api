package db

import (
	"crypto/tls"
	"net"
	"sync"

	"gitlab.com/gowtham-munukutla/vashisht-api/config"

	mgo "github.com/globalsign/mgo"
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

	tlsConfig := &tls.Config{}

	dialInfo := &mgo.DialInfo{
		Addrs:    mongoConfig.Hosts,
		Database: mongoConfig.Database,
		Username: mongoConfig.Username,
		Password: mongoConfig.Password,
		Source:   "scram-sha1",
	}

	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}

	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		panic(err.Error())
	}

	dbInstance = session.DB(mongoConfig.Database)
}
