package routes

import (
	"example.com/rest-api/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func registerForEvent(c *gin.Context) {
	var err error
	userId := c.GetInt64("userId")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Could not fetch event ID: %v", eventId)})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Could not fetch event ID: %v", eventId)})
		return
	}

	err = event.Register(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Could not register for event: %v", eventId)})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Registered for the event!", "event": event})
}

func cancelRegistrationForEvent(c *gin.Context) {
	var err error
	var event models.Event
	userId := c.GetInt64("userId")
	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Could not fetch event ID: %v", eventId)})
		return
	}

	event.ID = eventId

	err = event.CancelRegistration(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Could not cancel registration for event: %v", eventId)})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Cancelled registration for the event!", "event": event})
}
