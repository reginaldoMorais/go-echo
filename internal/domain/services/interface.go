package services

import (
	"go-echo/internal/api/responses"
	"go-echo/internal/domain/models"
)

type IClientService interface {
	GetClients() ([]responses.ClientResponse, responses.ClientErrorResponse)
	GetClientById(id string) (responses.ClientResponse, responses.ClientErrorResponse)
	CreateClient(client models.Client) (responses.ClientResponse, responses.ClientErrorResponse)
	DeleteClient(id string) responses.ClientErrorResponse
}
