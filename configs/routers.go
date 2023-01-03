package configs

import (
	"todo-server/handlers"

	"github.com/labstack/echo/v4"
)

type Router struct {
}

func (Router) NewRouter() *Router {
	return &Router{}
}

func (router *Router) Init(e *echo.Echo) {
	homeHandler := handlers.HomeHandler{Path: "/"}
	eventHandler := handlers.EventHandler{Path: "/events"}

	handlers := []handlers.Handler{&homeHandler, &eventHandler}
	for _, handler := range handlers {
		handler.Init(e.Group(handler.GetPath()))
	}
}
