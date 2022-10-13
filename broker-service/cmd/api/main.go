package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://*", "https://*"},
		AllowMethods: []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// routes
	e.GET("/health-check", HandleHealthTest)
	e.POST("/", HandleRequest)

	e.Logger.Fatal(e.Start(":1323"))
}
