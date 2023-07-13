package main

import (
	"prakerja6/config"
	"prakerja6/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	config.ConnectDatabase()
	e := echo.New()
	e = routes.InitRoute(e)
	e.Start(":8000")
}
