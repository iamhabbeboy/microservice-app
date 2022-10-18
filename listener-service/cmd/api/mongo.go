package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoClient struct {
	collection *mongo.Collection
}

type LogEntry struct {
	Name      string
	Data      string
	CreatedAt time.Time
}

func NewMongoClient() *MongoClient {
	mongoURL := os.Getenv("mongoURL")
	dbName := os.Getenv("dbName")
	collectionName := os.Getenv("collectionName")
	opt := options.Client().ApplyURI(mongoURL)
	opt.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})
	client, err := mongo.Connect(context.TODO(), opt)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB ... !!")
	if dbName == "" || collectionName == "" {
		log.Fatal("unable to get collection name")
	}

	db := client.Database(dbName)
	collection := db.Collection(collectionName)

	return &MongoClient{
		collection: collection,
	}
}

func (m *MongoClient) Save(data Payload) (any, error) {
	ctx := context.TODO()
	res, err := m.collection.InsertOne(ctx, LogEntry{
		Name:      data.Name,
		Data:      data.Data,
		CreatedAt: time.Now(),
	})
	if err != nil {
		return nil, err
	}

	return res.InsertedID, nil
}
