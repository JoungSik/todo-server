package handlers

import (
	"net/http"
	"todo-server/handlers/dto"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type EventHandler struct {
	Path string
}

func (eventHandler *EventHandler) GetPath() string {
	return eventHandler.Path
}

func (eventHandler *EventHandler) Init(e *echo.Group) {
	e.GET("", eventHandler.Index)
	e.POST("", eventHandler.Create)
	e.GET("/:id", eventHandler.Show)
	e.PATCH("/:id", eventHandler.Update)
	e.DELETE("/:id", eventHandler.Delete)
}

func (eventHandler *EventHandler) Index(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{})
}

func (eventHandler *EventHandler) Create(ctx echo.Context) error {
	eventDto := &dto.EventDto{}
	err := ctx.Bind(eventDto)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	event := eventDto.ToModel()
	db := ctx.Get("db").(*gorm.DB)
	result := db.Create(event)
	if result.Error != nil {
		return ctx.JSON(http.StatusBadRequest, result.Error.Error())
	} else {
		return ctx.JSON(http.StatusCreated, event)
	}
}

func (eventHandler *EventHandler) Show(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{})
}

func (eventHandler *EventHandler) Update(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{})
}

func (eventHandler *EventHandler) Delete(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{})
}
