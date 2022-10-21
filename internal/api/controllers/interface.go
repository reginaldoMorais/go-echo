package controllers

import (
	"github.com/labstack/echo"
)

type IClientController interface {
	CreateClient(c echo.Context) error
	GetClient(c echo.Context) error
}
