package controllers

import (
	"go-echo/internal/api/requests"
	"go-echo/internal/api/responses"
	"go-echo/internal/domain/models"
	"go-echo/internal/domain/services"
	"go-echo/internal/domain/types"
	"go-echo/pkg/printer"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

var validate = validator.New()

type ClientControllerImpl struct {
	Service services.IClientService
}

func NewClientController() IClientController {
	service := services.NewClientService()
	return &ClientControllerImpl{service}
}

func (cl *ClientControllerImpl) CreateClient(c echo.Context) error {
	var clientRequest requests.ClientRequest

	// validate the requests body
	if err := c.Bind(&clientRequest); err != nil {
		printer.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, responses.ClientErrorResponse{Code: 400, Message: err.Error()})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&clientRequest); validationErr != nil {
		printer.Error(validationErr.Error())
		return c.JSON(http.StatusInternalServerError, responses.ClientErrorResponse{Code: 400, Message: validationErr.Error()})
	}

	client := models.Client{Id: types.NewID(), Name: clientRequest.Name, Email: clientRequest.Email, Status: types.Status("ACTIVE")}

	clientResponse, err := cl.Service.CreateClient(client)
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
		return c.JSON(http.StatusNotFound, err)
	}

	printer.Infof("Return client %s", clientResponse.ToString())

	return c.JSON(http.StatusOK, clientResponse)
}
