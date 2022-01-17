package handlers

import (
	"github.com/gin-gonic/gin"
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
	c.JSON(http.StatusOK, gin.H{"id": id})
}
