package core

import (
	"context"
	"fmt"
	"go-echo/pkg/color"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(color.Fatal("Error to get client connected to MongoDB!"))
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println(color.Fatal("Error to health check MongoDB!"))
		log.Fatal(err)
	}

	fmt.Println(color.Green("Connected to MongoDB!"))

	return client
}

// Client instance
var DB *mongo.Client = ConnectDB()

// getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database(EnvMongoDatabase()).Collection(collectionName)
	return collection
}
