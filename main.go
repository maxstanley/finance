package main

import (
	"fmt"
	"log"

	"go.maxstanley.uk/finance/controllers"
	"go.maxstanley.uk/finance/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main()  {
	fmt.Println("Starting!")

	err := models.InitialiseDatabase("./finance.db")
	if err != nil {
		log.Fatal(err)
	}

	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Recover())
	e.Pre(middleware.AddTrailingSlash())

	helloGroup := e.Group("/hello")
	controllers.RegisterHelloRoutes(helloGroup)

	statementGroup := e.Group("/statement")
	controllers.RegisterStatementRoutes(statementGroup)

	transactionGroup := e.Group("/transaction")
	controllers.RegisterTransactionRoutes(transactionGroup)

	e.Logger.Fatal(e.Start(":3000"))
}

