package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://*", "https://*"},
		AllowMethods: []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusAccepted, "OK")
	})
	e.POST("/log", HandleRequest)
	e.Logger.Fatal(e.Start(":1323"))
}
