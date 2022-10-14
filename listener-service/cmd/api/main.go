package main

import (
	"context"
	"fmt"
	"time"
)

// connect to kafka
func main() {
	// r := kafka.NewReader(kafka.ReaderConfig{
	// 	Brokers:  []string{"localhost:9092"},
	// 	Topic:    "my-topic",
	// 	GroupID:  "my-group",
	// 	MinBytes: 5,
	// 	MaxBytes: 1e6,
	// 	// wait for at most 3 seconds before receiving new data
	// 	MaxWait: 3 * time.Second,
	// })

	fmt.Println("Kafka has been started")
	// create a new context
	ctx := context.Background()
	// produce messages in a new go routine, since
	// both the produce and consume functions are
	// blocking
	Consumer(ctx)
	time.Sleep(50)
}
