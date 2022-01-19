package storage

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

const (
	connectTimeout = 5
)

func GetConnection() (*mongo.Client, context.Context, context.CancelFunc) {
	log.Println(&StorageInstance.Config.ApplyURI)
	client, err := mongo.NewClient(options.Client().ApplyURI(StorageInstance.Config.ApplyURI))
	if err != nil {
		log.Println("Can't create mongodb client ", err)
	}

	if StorageInstance.Client == nil {
		StorageInstance.Client = client
	}

	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Printf("Failed to connect to cluster: %v", err)
	}
	// Force a connection to verify our connection string
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Printf("Failed to ping cluster: %v", err)
	}
	fmt.Println("Connected to MongoDB!")
	return client, ctx, cancel
}
