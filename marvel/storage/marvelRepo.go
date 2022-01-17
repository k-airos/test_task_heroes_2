package storage

import (
	"errors"
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

// GetHeroByID Retrives a hero by its id from the db
func GetHeroByID(id primitive.ObjectID) (*models.Hero, error) {
	var hero *models.Hero

	client, ctx, cancel := GetConnection()

	defer cancel()
	defer client.Disconnect(ctx)

	result := client.Database(StorageInstance.Config.DBname).Collection(StorageInstance.Config.CollectionName).FindOne(ctx, bson.M{"id": id})
	if result == nil {
		return nil, errors.New("Could not find a hero")
	}
	err := result.Decode(&hero)

	if err != nil {
		log.Printf("Failed marshalling %v", err)
		return nil, err
	}
	log.Printf("HERO: %v", hero)
	return hero, nil
}
