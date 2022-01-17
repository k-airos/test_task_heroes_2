package storage

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var StorageInstance *Storage

type Storage struct {
	Config *Config
	Client *mongo.Client
}

func New() *Storage {
	return &Storage{}
}
