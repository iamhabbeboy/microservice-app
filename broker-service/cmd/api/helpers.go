package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"
)

type Payload struct {
	Name string `json:"name"`
	Data any    `json:"data"`
}

func statusResponse(status bool, err string) JsonResponse {
	return JsonResponse{
		Error:   status,
		Message: err,
	}
}

func Producer(ctx context.Context) {
	fmt.Println("initializing....")
	// intialize the writer with the broker addresses, and the topic
	data := Payload{
		Name: "log",
		Data: "This is another test here",
	}
	j, _ := json.Marshal(&data)

	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "my-topic",
	})

	err := w.WriteMessages(ctx, kafka.Message{
		Key: []byte("this is a key"),
		// create an arbitrary message payload for the value
		Value: []byte(j),
	})
	if err != nil {
		panic("could not write message " + err.Error())
	}
}
