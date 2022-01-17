package handlers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"marvel/models"
	"marvel/storage"
	"net/http"
)

func HandleGetHeroes(c *gin.Context) {
	var loadedTasks, err = storage.GetAllHeroes()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"tasks": loadedTasks})
}

func HandleCreateHero(c *gin.Context) {
	var hero models.Hero
	if err := c.ShouldBindJSON(&hero); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	id, err := storage.Create(&hero)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(201, gin.H{"id": id})
}

func HandleGetHero(c *gin.Context) {
	//var hero models.Hero
	orderID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(orderID)

	//if err := c.BindUri(&hero); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"msg": err})
	//	return
	//}
	//var loadedHero, err = storage.GetHeroByID(hero.ID)
	var loadedHero, err = storage.GetHeroByID(docID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": err, "msg2": "no hero in db"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ID": loadedHero.ID, "Name": loadedHero.Name})
}

func HandleUpdateHero(c *gin.Context) {
	var hero models.Hero
	if err := c.ShouldBindJSON(&hero); err != nil {
		log.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}
	savedHero, err := storage.Update(&hero)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"hero": savedHero})
}
func HandleDeleteHero(c *gin.Context) {
	orderID := c.Params.ByName("id")
	docID, _ := primitive.ObjectIDFromHex(orderID)

	var err = storage.DeleteHeroByID(docID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"msg": err, "msg2": "no hero in db"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"msg": "Successfully delete hero"})
}
