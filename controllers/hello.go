package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterHelloRoutes(group *echo.Group) {
	group.GET("/", getHello)
}

func getHello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!")
}

