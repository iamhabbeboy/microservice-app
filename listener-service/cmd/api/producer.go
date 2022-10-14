package main

import (
	"context"

	"github.com/segmentio/kafka-go"
)

type Payload struct {
	Name string `json:"name"`
	Data any    `json:"data"`
}

func Producer(ctx context.Context) {
	// intialize the writer with the broker addresses, and the topic
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "my-topic",
	})

	err := w.WriteMessages(ctx, kafka.Message{
		Key: []byte("this is a key"),
		// create an arbitrary message payload for the value
		Value: []byte("this is message"),
	})
	if err != nil {
		panic("could not write message " + err.Error())
	}
}
