package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Payload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

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

func (m *MongoClient) GetAll() ([]LogEntry, error) {
	findOptions := options.Find()
	findOptions.SetLimit(10)
	var results []LogEntry

	cur, err := m.collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		//Create a value into which the single document can be decoded
		var log LogEntry
		err := cur.Decode(&log)
		if err != nil {
			return nil, err
		}
		results = append(results, log)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	cur.Close(context.TODO())

	return results, nil
}
