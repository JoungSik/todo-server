package main

import (
	"net/http"
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

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!!!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
