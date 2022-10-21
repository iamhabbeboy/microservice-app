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
		return c.String(http.StatusBadRequest, err.Error())
	}
	fmt.Println(params)
	mongo := NewMongoClient()
	_, err := mongo.Save(Payload{
		Name: params.Name,
		Data: params.Data,
	})
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	resp := JsonResponse{
		Error:   false,
		Data:    params,
		Message: "logged",
	}

	return c.JSON(http.StatusOK, resp)
}

func HandleLogs(c echo.Context) error {
	mongo := NewMongoClient()
	data, err := mongo.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadGateway, err.Error())
	}
	// j, err := json.Marshal(data)
	if err != nil {
		return c.JSON(http.StatusBadGateway, err.Error())
	}
	resp := JsonResponse{
		Error:   false,
		Data:    data,
		Message: "logs",
	}
	return c.JSON(http.StatusOK, resp)
}
