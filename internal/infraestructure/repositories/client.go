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

func (r *ClientRepository) FindAll() ([]models.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	cursor, err := r.Collection.Find(ctx, bson.D{{}})

	var clients []models.Client

	for cursor.Next(context.TODO()) {
		var elem models.Client

		err := cursor.Decode(&elem)
		if err != nil {
			fmt.Println(err)
		}

		clients = append(clients, elem)

	}

	return clients, err
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

	_, err := r.Collection.InsertOne(ctx, client)

	return client, err
}

func (r *ClientRepository) DeleteClient(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	_, err := r.Collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		fmt.Println(err)
		return err
	}

	return err
}
