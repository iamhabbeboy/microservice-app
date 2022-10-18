package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RequestPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

type JsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func HandleRequest(c echo.Context) error {
	params := new(RequestPayload)
	if err := c.Bind(params); err != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("Unable to process data: %v", err.Error()))
	}

	mongo := NewMongoClient()
	mongo.Save(Payload{
		Name: params.Name,
		Data: params.Data,
	})

	// return c.JSON(http.StatusBadRequest, resp)

	resp := JsonResponse{
		Error:   false,
		Data:    "sdfdf",
		Message: "logged",
	}

	return c.JSON(http.StatusAccepted, resp)
}
