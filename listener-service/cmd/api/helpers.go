package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func handlePayload(msg Payload) error {
	switch msg.Name {
	case "log":
		return callLoggerService(msg)
	case "auth":
		// ---- other operation here
	default:
	}
	return nil
}

func callLoggerService(data Payload) error {
	js, err := json.MarshalIndent(data, "", "\t")
	fmt.Println(data)
	fmt.Println(string(js))
	// js, err := json.Marshal(data)
	if err != nil {
		return err
	}
	// call the service
	request, err := http.NewRequest("POST", "http://logger-service:3500/log", bytes.NewBuffer(js))
	request.Header.Set("content-type", "application/json")
	if err != nil {
		return err
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusBadRequest {
		return errors.New("error occured")
	}
	return nil
}
