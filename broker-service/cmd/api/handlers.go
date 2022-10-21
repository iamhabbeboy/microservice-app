package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RequestPayload struct {
	Action string `json:"action"`
	Auth   AuthPayload
	Log    LoggerPayload
}

// payload for auth
type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoggerPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

type JsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `string:"data,omitempty"`
}

type LogEntry struct {
	Data any
}

func HandleHealthTest(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

func HandleRequest(c echo.Context) error {
	params := new(RequestPayload)
	if err := c.Bind(params); err != nil {
		return c.JSON(http.StatusOK, statusResponse(true, err.Error()))
	}

	switch params.Action {
	case "auth":
		return c.JSON(http.StatusAccepted, authentication(params.Auth))
	case "log":
		return c.JSON(http.StatusAccepted, logger(params.Log))
	case "logs":
		return c.JSON(http.StatusAccepted, getLogs())
	case "broker":
		return c.JSON(http.StatusOK, statusResponse(false, "This is from broker"))
	default:
		return c.JSON(http.StatusOK, statusResponse(true, errors.New("unknown action").Error()))
	}
}

func getLogs() JsonResponse {
	// call the service
	req, err := http.Get("http://logger-service:3500/logs")
	if err != nil {
		return statusResponse(true, err.Error())
	}
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return statusResponse(true, err.Error())
	}
	var log LogEntry
	err = json.Unmarshal(body, &log)
	if err != nil {
		return statusResponse(true, err.Error())
	}
	var p JsonResponse
	p.Error = false
	p.Message = "logged"
	p.Data = log

	return p
}

func authentication(data AuthPayload) JsonResponse {
	js, err := json.MarshalIndent(data, "", "\t")
	// js, err := json.Marshal(data)
	if err != nil {
		return statusResponse(true, err.Error())
	}
	// call the service
	request, err := http.NewRequest("POST", "http://auth-service/api/", bytes.NewBuffer(js))
	request.Header.Set("content-type", "application/json")
	if err != nil {
		return statusResponse(true, err.Error())
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return statusResponse(true, err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return statusResponse(true, "invalid credentials")
	} else if resp.StatusCode != http.StatusAccepted {
		return statusResponse(true, "error calling auth service: %s")
	}

	var jsonResponse JsonResponse
	err = json.NewDecoder(resp.Body).Decode(&jsonResponse)
	if err != nil {
		return statusResponse(true, err.Error())
	}

	if jsonResponse.Error {
		return statusResponse(true, string(http.StatusUnauthorized))
	}

	var p JsonResponse
	p.Error = false
	p.Message = "Authenticated!:"
	p.Data = jsonResponse.Data

	return p
}

func logger(log LoggerPayload) JsonResponse {
	evt := NewEvent(context.Background(), "")
	data := Payload{
		Name: log.Name,
		Data: log.Data,
	}
	evt.Set(data)

	var p JsonResponse
	p.Error = false
	p.Message = log.Name
	p.Data = log.Data

	return p
}
