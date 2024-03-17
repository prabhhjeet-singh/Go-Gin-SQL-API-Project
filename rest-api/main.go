package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"example.com/rest-api/models"
)

func main() {
	router := gin.Default()
	router.GET("/events", getEvent)
	router.POST("/events", postEvent)

	router.Run("localhost:8081")
}

func getEvent(c *gin.Context) {
	events := models.GetAllEvents()
	c.JSON(http.StatusOK, events)
}

func postEvent(c *gin.Context) {
	var newEvent models.Event

	if err := c.BindJSON(&newEvent); err != nil {
		return
	}

	models.Save(newEvent)
	c.JSON(http.StatusCreated, newEvent)

}
