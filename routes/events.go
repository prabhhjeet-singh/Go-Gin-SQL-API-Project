package routes

import (
	"net/http"
	"strconv"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEvent(c *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch, try again later"})
	}

	c.JSON(http.StatusOK, events)
}

func postEvent(c *gin.Context) {
	var newEvent models.Event

	if err := c.BindJSON(&newEvent); err != nil {
		return
	}

	userId := c.GetInt("userId")
	newEvent.UserId = userId

	models.Save(newEvent)
	c.JSON(http.StatusCreated, newEvent)

}

func getEventByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse ID"})
		return
	}

	event, err := models.GetEventByID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not find"})
		return
	}

	c.JSON(http.StatusFound, event)

}

func updateEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse ID"})
		return
	}

	_, err = models.GetEventByID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		return
	}

	var updatedEvent models.Event

	if err := c.BindJSON(&updatedEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request"})
		return
	}

	updatedEvent.ID = id
	err = models.Update(updatedEvent)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event updated succesfully"})
}

func deleteEvent(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not parse ID"})
		return
	}

	err = models.Delete(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event deleted succesfully"})
}
