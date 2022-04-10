package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func connectToDatabase(uri string) *mongo.Database {
	clientOptions := options.Client().ApplyURI(uri) // Set client options
	client, err := mongo.Connect(context.TODO(), clientOptions) // Connect to the Database
	if err != nil {
		fmt.Printf("Unknown error while connecting to the MongoDB database on URL: %s\n", uri)
		os.Exit(1)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Printf("Unknown error while checking the connection with MongoDB database on URL: %s\n", uri)
		os.Exit(1)
	}

	fmt.Printf("Successfully connected to the MongoDB database!!!\n")

	database := client.Database("Syntil")
	return database
}