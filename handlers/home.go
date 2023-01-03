package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HomeHandler struct {
	Path string
}

func (homeHandler *HomeHandler) GetPath() string {
	return homeHandler.Path
}

func (homeHandler *HomeHandler) Init(e *echo.Group) {
	e.GET("", homeHandler.Index)
	e.POST("", homeHandler.Create)
	e.GET("/:id", homeHandler.Show)
	e.PATCH("/:id", homeHandler.Update)
	e.DELETE("/:id", homeHandler.Delete)
}

func (homeHandler *HomeHandler) Index(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{
		"message": "hello World",
	})
}

func (homeHandler *HomeHandler) Create(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{})
}

func (homeHandler *HomeHandler) Show(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{})
}

func (homeHandler *HomeHandler) Update(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{})
}

func (homeHandler *HomeHandler) Delete(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{})
}
