package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

type Event struct {
	URL     string
	Topic   string
	GroupID string
	Ctx     context.Context
	mongodb MongoClient
}

type Payload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

func NewEvent(ctx context.Context, topic string) *Event {
	tp := os.Getenv("topic")
	if topic != "" {
		tp = topic
	}
	return &Event{
		URL:     os.Getenv("kafkaURL"),
		Topic:   tp,
		Ctx:     ctx,
		GroupID: os.Getenv("groupID"),
		mongodb: *NewMongoClient(),
	}
}

func (e *Event) Writer() *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(e.URL),
		Topic:    e.Topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func (e *Event) Reader() *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{e.URL},
		GroupID:  e.GroupID,
		Topic:    e.Topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
}

func (e *Event) Set(data Payload) {
	w := e.Writer()
	j, err := json.Marshal(data)
	if err != nil {
		log.Fatalln("unable to marshal payload:", err)
	}
	msg := kafka.Message{
		Key:   []byte(data.Name),
		Value: []byte(j),
	}
	err = w.WriteMessages(e.Ctx, msg)

	if err != nil {
		log.Fatalln(err)
	}
}

func (e *Event) Get() {
	reader := e.Reader()
	ctx := context.Background()
	for {
		msg, err := reader.ReadMessage(ctx)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Data received: ", string(msg.Value))

		var payload Payload
		err = json.Unmarshal(msg.Value, &payload)
		if err != nil {
			log.Fatal(err)
		}
		err = handlePayload(payload)
		if err != nil {
			log.Fatal(err)
		}
		// data, err := e.mongodb.Save(payload)
		// if err != nil {
		// 	log.Fatal("Unable:: ", err)
		// }
		// fmt.Println("Inserted a single document: ", data)
	}
}
