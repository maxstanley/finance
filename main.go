package main

import (
	"fmt"

	"go.maxstanley.uk/finance/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main()  {
	fmt.Println("Starting!")

	e := echo.New()
	e.HideBanner = true
	e.Pre(middleware.AddTrailingSlash())

	helloGroup := e.Group("/hello")
	controllers.RegisterHelloRoutes(helloGroup)

	e.Logger.Fatal(e.Start(":3000"))
}

