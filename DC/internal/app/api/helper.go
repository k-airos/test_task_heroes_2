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
	r.GET("/dc/", handlers.HandleGetHeroes)
	r.GET("/dc/:id", handlers.HandleGetHero)
	r.PUT("/dc/", handlers.HandleUpdateHero)
	r.POST("/dc/", handlers.HandleCreateHero)
	r.DELETE("/dc/:id", handlers.HandleDeleteHero)
	r.Run(Server.Config.Port) // listen and serve on 0.0.0.0:8081
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
