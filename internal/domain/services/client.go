package services

import (
	"go-echo/internal/api/responses"
	"go-echo/internal/domain/models"
	"go-echo/internal/domain/repositories"
	repositoriesImpl "go-echo/internal/infraestructure/repositories"
	"go-echo/pkg/printer"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

var validate = validator.New()
var (
	ClientService = &ClientServiceImpl{}
)

type ClientServiceImpl struct {
	Repository repositories.IClientRepository
}

func NewClientService() IClientService {
	repository := repositoriesImpl.NewClientRepository()
	return &ClientServiceImpl{repository}
}

func (s *ClientServiceImpl) CreateClient(c echo.Context) (responses.ClientResponse, responses.ClientErrorResponse) {
	var client models.Client

	// validate the request body
	if err := c.Bind(&client); err != nil {
		printer.Error(err.Error())
		return responses.ClientResponse{}, responses.ClientErrorResponse{Code: 400, Message: err.Error()}
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&client); validationErr != nil {
		printer.Error(validationErr.Error())
		return responses.ClientResponse{}, responses.ClientErrorResponse{Code: 400, Message: validationErr.Error()}
	}

	printer.Infof("Creating client: %s", client.ToString())

	newClient, err := s.Repository.Create(client)
	if err != nil {
		printer.Error(err.Error())
		return responses.ClientResponse{}, responses.ClientErrorResponse{Code: 500, Message: err.Error()}
	}

	return responses.ClientResponse{Id: newClient.Id, Name: newClient.Name, Email: newClient.Email, Status: client.Status}, responses.ClientErrorResponse{}
}

func (s *ClientServiceImpl) GetClientById(clientId string) (responses.ClientResponse, responses.ClientErrorResponse) {
	client, err := s.Repository.FindById(clientId)
	if err != nil {
		return responses.ClientResponse{}, responses.ClientErrorResponse{Code: 500, Message: err.Error()}
	}

	return responses.ClientResponse{Id: client.Id, Name: client.Name, Email: client.Email, Status: client.Status}, responses.ClientErrorResponse{}
}
