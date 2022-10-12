package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Handles all request
// Specify request payload struct/standard
// switch case action sent in
//
// Call logger service endpoint to log request
// Send request to kafka/listener
// Return response

// Logs data to Mongodb

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
