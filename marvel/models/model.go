package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Hero struct {
	ID           primitive.ObjectID
	Name         string `json:"hero_name"`
	CreationDate int    `json:"creation_date"'`
}
