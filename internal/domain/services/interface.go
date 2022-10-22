package services

import (
	"go-echo/internal/api/responses"
	"go-echo/internal/domain/models"
)

type IClientService interface {
	CreateClient(client models.Client) (responses.ClientResponse, responses.ClientErrorResponse)
	GetClientById(clientId string) (responses.ClientResponse, responses.ClientErrorResponse)
}
