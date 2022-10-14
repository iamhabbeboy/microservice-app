package main

import (
	"bytes"
	"encoding/json"
	"errors"
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

func HandleHealthTest(c echo.Context) error {
	return c.String(http.StatusOK, "All is well from Broker")
}

func HandleRequest(c echo.Context) error {
	params := new(RequestPayload)
	if err := c.Bind(params); err != nil {
		return c.JSON(http.StatusOK, statusResponse(false, err.Error()))
	}

	switch params.Action {
	case "auth":
		return c.JSON(http.StatusAccepted, authentication(params.Auth))
	case "log":
		return c.JSON(http.StatusAccepted, logger(params.Log))
	default:
		return c.JSON(http.StatusOK, statusResponse(false, errors.New("unknown action").Error()))
	}

	return c.JSON(http.StatusOK, statusResponse(false, "This is from broker"))
}

func authentication(data AuthPayload) JsonResponse {
	js, err := json.MarshalIndent(data, "", "\t")
	// js, err := json.Marshal(data)
	if err != nil {
		return statusResponse(false, err.Error())
	}
	// call the service
	request, err := http.NewRequest("POST", "http://auth-service/api/", bytes.NewBuffer(js))
	request.Header.Set("content-type", "application/json")
	if err != nil {
		return statusResponse(false, err.Error())
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return statusResponse(false, err.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		return statusResponse(false, errors.New("invalid credentials").Error())
	} else if resp.StatusCode != http.StatusAccepted {
		return statusResponse(false, errors.New("error calling auth service").Error())
	}

	var jsonResponse JsonResponse
	err = json.NewDecoder(resp.Body).Decode(&jsonResponse)
	if err != nil {
		return statusResponse(false, err.Error())
	}

	if jsonResponse.Error {
		return statusResponse(false, string(http.StatusUnauthorized))
	}

	var p JsonResponse
	p.Error = false
	p.Message = "Authenticated!:"
	p.Data = jsonResponse.Data

	return p
}

func logger(log LoggerPayload) JsonResponse {
	j, err := json.Marshal(log)
	if err != nil {
		return statusResponse(false, errors.New("Unable to marshal data").Error())
	}
	uri := "http://logger-service:1323/log"
	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(j))
	req.Header.Set("content-type", "application/json")
	if err != nil {
		return statusResponse(false, err.Error())
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return statusResponse(false, err.Error())
	}
	defer resp.Body.Close()

	var jsonResp JsonResponse
	err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	if err != nil {
		return statusResponse(false, errors.New("unable to decode response").Error())
	}

	if resp.StatusCode != http.StatusAccepted {
		return statusResponse(false, jsonResp.Message)
	}

	return JsonResponse{
		Error:   false,
		Data:    jsonResp.Data,
		Message: jsonResp.Message,
	}
}
