package main

// Logs data to Mongodb

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	port     = "80"
	mongoURL = "mongodb://mongo:27017"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
