package db

import (
	"gitlab.com/gowtham-munukutla/vashisht-api/config"
	"context"
	"log"
	"sync"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
)

var (
	mongoClient *mongo.Client
	once        sync.Once
)

func getMongoClient() *mongo.Client {
	once.Do(func() {
		connectDBOfficial()
	})

	return mongoClient
}

func GetMongoCollectionWithContext(collectionName string) (*mongo.Collection, context.Context) {
	mongoClient = getMongoClient()
	collection := mongoClient.Database("vashisht").Collection(collectionName)
	ctx, err := context.WithTimeout(context.Background(), 10*time.Second)
	if err != nil {
		log.Println(err)
	}
	return collection, ctx
}

func connectDBOfficial() {
	appConfigInstance := config.GetAppConfig()

	mClient, err := mongo.NewClient(appConfigInstance.MongoConfig.Hosts)
	if err != nil {
		panic(err.Error())
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = mClient.Connect(ctx)
	if err != nil {
		panic(err.Error())
	}
	mongoClient = mClient
}
