package controllers

import (
	"github.com/labstack/echo"
)

type IClientController interface {
	GetClients(c echo.Context) error
	GetClient(c echo.Context) error
	CreateClient(c echo.Context) error
	DeleteClient(c echo.Context) error
}
