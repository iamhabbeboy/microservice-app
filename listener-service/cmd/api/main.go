package main

import (
	"context"
	"fmt"
)

// connect to kafka
func main() {
	fmt.Println("Kafka has been started")
	evt := NewEvent(context.Background(), "")
	evt.Get()
}
