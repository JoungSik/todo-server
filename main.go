package main

import (
	"todo-server/configs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	config := configs.GetConfig()

	// middleware
	e.Use(config.ContextDB)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Logger.Fatal(e.Start(":1323"))
}
