package controllers

import (
	"go-echo/printer"
	"go-echo/responses"
	"go-echo/services"
	"net/http"

	"github.com/labstack/echo"
)

func CreateClient(c echo.Context) error {
	client, err := services.CreateClient(c)
	if err.Message != "" {
		if err.Code == 400 {
			return c.JSON(http.StatusBadRequest, err)
		}
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, responses.ClientResponse{Id: client.Id, Name: client.Name, Email: client.Email})
}

func GetClient(c echo.Context) error {
	clientId := c.Param("id")
	client, err := services.GetClientById(c, clientId)
	if err.Message != "" {
		return c.JSON(http.StatusInternalServerError, err)
	}

	printer.Infof("Return client %s", client.ToString())

	return c.JSON(http.StatusOK, responses.ClientResponse{Id: client.Id, Name: client.Name, Email: client.Email})
}
