package api

import (
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"log"
	"marvel/handlers"
	"marvel/storage"
)

func ConfigureRouterField() {
	r := gin.Default()
	r.GET("/marvel/", handlers.HandleGetHeroes)
	//r.GET("/marvel/:id", handlers.HandleGetHero)
	//r.PUT("/marvel/", handlers.HandleUpdateHero)
	r.POST("/marvel/", handlers.HandleCreateHero)
	r.Run() // listen and serve on 0.0.0.0:8080
}
func ConfigureStorageField() {
	storageConfig := storage.NewConfig()
	_, err := toml.DecodeFile(Server.Config.ConfigPath, storageConfig)
	if err != nil {
		storageConfig.DBname = "heroes"
		storageConfig.ApplyURI = "mongodb://localhost:27017" // VERY BAD PRACTICE
		log.Println("Can't read config file, use default values", err)
	}
	StorageInstance := storage.New()
	StorageInstance.Config = storageConfig
	storage.StorageInstance = StorageInstance
}
