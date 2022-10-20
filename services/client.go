package services

import (
	"context"
	"fmt"
	"go-echo/configs"
	"go-echo/models"
	"go-echo/printer"
	"go-echo/responses"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "client")
var validate = validator.New()

func CreateClient(c echo.Context) (responses.ClientResponse, responses.ClientErrorResponse) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var client models.Client
	defer cancel()

	//validate the request body
	if err := c.Bind(&client); err != nil {
		printer.Error(err.Error())
		return responses.ClientResponse{}, responses.ClientErrorResponse{Code: 400, Message: err.Error()}
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&client); validationErr != nil {
		printer.Error(validationErr.Error())
		return responses.ClientResponse{}, responses.ClientErrorResponse{Code: 400, Message: validationErr.Error()}
	}

	newClient := models.Client{
		Id:    client.Id,
		Name:  client.Name,
		Email: client.Email,
	}

	result, err := userCollection.InsertOne(ctx, newClient)

	if err != nil {
		printer.Error(err.Error())
		return responses.ClientResponse{}, responses.ClientErrorResponse{Code: 500, Message: err.Error()}
	}

	fmt.Println(result)

	return responses.ClientResponse{Id: client.Id, Name: client.Name, Email: client.Email}, responses.ClientErrorResponse{}
}

func GetClientById(c echo.Context, clientId string) (responses.ClientResponse, responses.ClientErrorResponse) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var client models.Client
	defer cancel()

	err := userCollection.FindOne(ctx, bson.M{"_id": clientId}).Decode(&client)

	if err != nil {
		return responses.ClientResponse{}, responses.ClientErrorResponse{Code: 500, Message: err.Error()}
	}

	return responses.ClientResponse{Id: client.Id, Name: client.Name, Email: client.Email}, responses.ClientErrorResponse{}
}
