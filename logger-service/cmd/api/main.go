package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://*", "https://*"},
		AllowMethods:     []string{"POST", "GET", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderXCSRFToken, echo.HeaderAuthorization},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Link"},
		MaxAge:           300,
	}))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusAccepted, "OK")
	})
	e.POST("/log", HandleRequest)

	e.GET("/logs", HandleLogs)
	e.Logger.Fatal(e.Start(":3500"))
}
