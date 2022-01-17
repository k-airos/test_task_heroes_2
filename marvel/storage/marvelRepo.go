package storage

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"marvel/models"
)

// GetAllHeroes Retrives all heroes from the db
func GetAllHeroes() ([]*models.Hero, error) {
	var heroes []*models.Hero

	client, ctx, cancel := GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	db := client.Database(StorageInstance.Config.DBname)
	collection := db.Collection(StorageInstance.Config.CollectionName)
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	err = cursor.All(ctx, &heroes)
	if err != nil {
		log.Printf("Failed marshalling %v", err)
		return nil, err
	}
	return heroes, nil
}

//Create creating a hero in a mongo
func Create(hero *models.Hero) (primitive.ObjectID, error) {
	client, ctx, cancel := GetConnection()
	defer cancel()
	defer client.Disconnect(ctx)
	hero.ID = primitive.NewObjectID()

	result, err := client.Database(StorageInstance.Config.DBname).Collection(StorageInstance.Config.CollectionName).InsertOne(ctx, hero)
	if err != nil {
		log.Printf("Could not create Hero: %v", err)
		return primitive.NilObjectID, err
	}
	oid := result.InsertedID.(primitive.ObjectID)
	return oid, nil
}
