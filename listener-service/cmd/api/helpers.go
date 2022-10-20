package main

import (
	"bytes"
	"encoding/json"
	"errors"
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
	// js, err := json.Marshal(data)
	if err != nil {
		return err
	}
	// call the service
	request, err := http.NewRequest("POST", "http://logger-service/log", bytes.NewBuffer(js))
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

	if resp.StatusCode == http.StatusUnauthorized {
		return errors.New("invalid credentials")
	} else if resp.StatusCode != http.StatusAccepted {
		return errors.New("error calling auth service")
	}
	return nil
}
