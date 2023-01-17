package main

import (
	"todo-server/configs"
	"todo-server/kafka"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	config := configs.GetConfig()
	kafka := kafka.KafkaConfig{}.Init()

	// middleware
	e.Use(config.ContextDB)
	e.Use(kafka.KafkaMiddleware)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routers := configs.Router{}.NewRouter()
	routers.Init(e)

	e.Logger.Fatal(e.Start(":1323"))
}
