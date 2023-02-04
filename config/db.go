package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MgConnect(ac AppConfig) *mongo.Client {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	uri := fmt.Sprintf("mongodb+srv://%s:%s@bob.ynw24pz.mongodb.net/?retryWrites=true&w=majority", ac.DBUser, ac.DBPass)
	clientOptions := options.Client().
		ApplyURI(uri).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connection is good")

	return client
}

func MgCollection(coll string, client *mongo.Client) *mongo.Collection {
	return client.Database("bob").Collection(coll)
}
