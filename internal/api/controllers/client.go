package controllers

import (
	"go-echo/internal/domain/services"
	"go-echo/pkg/printer"
	"net/http"

	"github.com/labstack/echo"
)

var (
	ClientController = &ClientControllerImpl{}
)

type ClientControllerImpl struct {
	Service services.IClientService
}

func NewClientController() IClientController {
	service := services.NewClientService()
	return &ClientControllerImpl{service}
}

func (cl *ClientControllerImpl) CreateClient(c echo.Context) error {
	clientResponse, err := cl.Service.CreateClient(c)
	if err.Message != "" {
		if err.Code == 400 {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, clientResponse)
}

func (cl *ClientControllerImpl) GetClient(c echo.Context) error {
	clientId := c.Param("id")
	clientResponse, err := cl.Service.GetClientById(clientId)
	if err.Message != "" {
		return c.JSON(http.StatusInternalServerError, err)
	}

	printer.Infof("Return client %s", clientResponse.ToString())

	return c.JSON(http.StatusOK, clientResponse)
}
