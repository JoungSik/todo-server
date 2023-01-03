package handlers

import "github.com/labstack/echo/v4"

type Handler interface {
	Init(e *echo.Group)
	GetPath() string
	Index(ctx echo.Context) error
	Create(ctx echo.Context) error
	Show(ctx echo.Context) error
	Update(ctx echo.Context) error
	Delete(ctx echo.Context) error
}
