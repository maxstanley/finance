package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterHelloRoutes(group *echo.Group) {
	group.GET("/", get)
}

func get(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!")
}

