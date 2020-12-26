package controllers

import (
	"net/http"

	"go.maxstanley.uk/finance/models"

	"github.com/labstack/echo/v4"
)

func RegisterTransactionRoutes(group *echo.Group) {
	group.GET("/", getTransaction)
}

func getTransaction(c echo.Context) error {
	t := models.Transaction{}
	models.Database.First(&t)

	return c.JSON(http.StatusOK, t)
}

