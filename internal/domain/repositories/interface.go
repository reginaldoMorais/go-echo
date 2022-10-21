package repositories

import "go-echo/internal/domain/models"

type IClientRepository interface {
	//FindAll() ([]models.Client, error)
	FindById(id string) (models.Client, error)
	Create(client models.Client) (models.Client, error)
}
