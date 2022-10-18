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
}

func NewEvent(ctx context.Context, topic string) *Event {
	tp := os.Getenv("topic")
	if topic != "" {
		tp = topic
	}
	return &Event{
		URL:     os.Getenv("kafkaURL"),
		Topic:   tp,
		GroupID: os.Getenv("groupID"),
		Ctx:     ctx,
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
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Data received: ", string(msg.Value))
	}
}
