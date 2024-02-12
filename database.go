// database.go

package main

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client   *mongo.Client
	database *mongo.Database
)

func ConnectDatabase() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Get MongoDB URI from environment variable
	connStr := os.Getenv("MONGODB_URI")
	if connStr == "" {
		log.Fatal("MONGODB_URI environment variable not set")
	}

	clientOptions := options.Client().ApplyURI(connStr)
	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Optionally, get the database name from an environment variable
	dbName := os.Getenv("MONGODB_DB_NAME")
	if dbName == "" {
		log.Fatal("MONGODB_DB_NAME environment variable not set")
	}
	database = client.Database(dbName)

	return nil
}

func GetDatabase() *mongo.Database {
	return database
}

func PingDatabase() error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	return client.Ping(ctx, nil)
}
