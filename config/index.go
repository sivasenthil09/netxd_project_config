package config

import (
	"context"
	"log"
	"github.com/sivasenthil09/netxd_project_config/constants"
	"time"
)

var MongoClient *mongo.Client

func ConnectDatabase() (*mongo.Client, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	mongoconnection := options.Client().ApplyURI(constants.ConnectionString)

	MongoClient, err := mongo.Connect(ctx, mongoconnection)

	if err != nil {
		log.Fatal(err.Error())
		return nil, err
	}
	// if err := MongoClient.Ping(ctx); err != nil {
	// 	log.Fatal(err.Error())
	// 	return nil, err
	// }

	return MongoClient, nil
}

func GetCollection(client *mongo.Client, dbName string, collectionName string) *mongo.Collection {
	collection := client.Database(dbName).Collection(collectionName)
	return collection
}
