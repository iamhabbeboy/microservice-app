package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/segmentio/kafka-go"
)

const (
	host    = "localhost:9092"
	topic   = "my-topic"
	groupID = "demo-group"
)

// type Payload struct {
// 	Name string `json:"name"`
// 	Data any    `json:"data"`
// }

func Consumer(ctx context.Context) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{host},
		Topic:   topic,
		GroupID: groupID,
	})
	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		// after receiving the message, log its value
		// fmt.Println(string(msg.))
		fmt.Println("received: ", string(msg.Value))
		var payload Payload
		_ = json.Unmarshal(msg.Value, &payload)
		handlePayload(payload)
	}
}

func handlePayload(data Payload) {
	switch data.Name {
	case "log":
	}
}
