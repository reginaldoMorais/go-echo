package services

import (
	"go-echo/internal/api/responses"
	"go-echo/internal/domain/models"
	"go-echo/internal/domain/repositories"
	repositoriesImpl "go-echo/internal/infraestructure/repositories"
	"go-echo/pkg/printer"
)

type ClientServiceImpl struct {
	Repository repositories.IClientRepository
}

func NewClientService() IClientService {
	repository := repositoriesImpl.NewClientRepository()
	return &ClientServiceImpl{repository}
}

func (s *ClientServiceImpl) GetClients() ([]responses.ClientResponse, responses.ClientErrorResponse) {
	var clients []models.Client
	clients, err := s.Repository.FindAll()
	if err != nil {
		printer.Errorf("Error %s", err.Error())
		return []responses.ClientResponse{}, responses.ClientErrorResponse{}
	}

	var clientResponses []responses.ClientResponse
	for _, c := range clients {
		clientResponses = append(clientResponses, responses.ClientResponse{Id: c.Id, Name: c.Name, Email: c.Email, Status: c.Status})
	}

	return clientResponses, responses.ClientErrorResponse{}
}

func (s *ClientServiceImpl) GetClientById(id string) (responses.ClientResponse, responses.ClientErrorResponse) {
	client, err := s.Repository.FindById(id)
	if err != nil {
		printer.Errorf("Error when find for client: %s", id)
		printer.Errorf("Error %s", err.Error())
		return responses.ClientResponse{}, responses.ClientErrorResponse{Code: 404, Message: "Client not found"}
	}

	return responses.ClientResponse{Id: client.Id, Name: client.Name, Email: client.Email, Status: client.Status}, responses.ClientErrorResponse{}
}

func (s *ClientServiceImpl) CreateClient(client models.Client) (responses.ClientResponse, responses.ClientErrorResponse) {
	printer.Infof("Creating client: %s", client.ToString())

	newClient, err := s.Repository.Create(client)
	if err != nil {
		printer.Errorf("Error %s", err.Error())
		return responses.ClientResponse{}, responses.ClientErrorResponse{Code: 500, Message: err.Error()}
	}

	return responses.ClientResponse{Id: newClient.Id, Name: newClient.Name, Email: newClient.Email, Status: client.Status}, responses.ClientErrorResponse{}
}

func (s *ClientServiceImpl) DeleteClient(id string) responses.ClientErrorResponse {
	err := s.Repository.DeleteClient(id)
	if err != nil {
		printer.Errorf("Error when find for client: %s", id)
		printer.Errorf("Error %s", err.Error())
		return responses.ClientErrorResponse{Code: 404, Message: "Client not found"}
	}

	return responses.ClientErrorResponse{}
}
