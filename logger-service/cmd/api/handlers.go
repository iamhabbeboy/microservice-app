package main

import (
	"fmt"
	"main/data"
	"net/http"
	"time"

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

	event := data.LogEntry{
		Name:      params.Name,
		Data:      params.Data,
		CreatedAt: time.Now(),
	}
	model := data.NewModel()
	err := model.Create(event)
	if err != nil {
		resp := JsonResponse{
			Error:   true,
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, resp)
	}
	r, err := model.Read()
	if err != nil {
		resp := JsonResponse{
			Error:   true,
			Message: err.Error(),
		}
		return c.JSON(http.StatusBadRequest, resp)
	}

	resp := JsonResponse{
		Error:   false,
		Data:    string(r),
		Message: "logged",
	}

	return c.JSON(http.StatusAccepted, resp)
}
