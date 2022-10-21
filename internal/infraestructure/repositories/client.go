package repositories

import (
	"context"
	"fmt"
	"go-echo/internal/core"
	"go-echo/internal/domain/models"
	"go-echo/internal/domain/repositories"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ClientRepository struct {
	Collection *mongo.Collection
}

func NewClientRepository() repositories.IClientRepository {
	collection := core.GetCollection(core.DB, "client")
	return &ClientRepository{collection}
}

func (r *ClientRepository) FindById(id string) (models.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	client := models.Client{}
	err := r.Collection.FindOne(ctx, bson.M{"_id": id}).Decode(&client)

	return client, err
}

func (r *ClientRepository) Create(client models.Client) (models.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	result, err := r.Collection.InsertOne(ctx, client)

	fmt.Println(result)

	return client, err
}
