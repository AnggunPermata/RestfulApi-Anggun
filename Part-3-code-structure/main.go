package main

import (
	"project/anggun/config"
	// "project/anggun/middlewares"
	"project/anggun/routes"

	"github.com/labstack/echo"
)

func main() {
	config.InitDB()
	e := routes.New()
	// middlewares.LogMiddleWares(e)
	echo.New().Logger.Fatal(e.Start(":8080"))
}
