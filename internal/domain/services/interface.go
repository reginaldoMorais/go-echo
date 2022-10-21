package services

import (
	"go-echo/internal/api/responses"

	"github.com/labstack/echo"
)

type IClientService interface {
	CreateClient(c echo.Context) (responses.ClientResponse, responses.ClientErrorResponse)
	GetClientById(clientId string) (responses.ClientResponse, responses.ClientErrorResponse)
}
