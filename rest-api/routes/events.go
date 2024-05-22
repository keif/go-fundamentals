package routes

import (
	"example.com/rest-api/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event) // works like the scan, pass a pointer
	// internally gin looks at incoming body and stores in the var
	// data must match types of the model
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse request data."})
	}

	event.ID = 1
	event.UserID = 1
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save event. Try again later."})
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func getBase(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func getEventById(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Could not fetch event ID: %v", eventId)})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch event. Try again later."})
	}

	context.JSON(http.StatusOK, event)
}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch events. Try again later."})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getPing(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func updateEvent(context *gin.Context) {
	var err error
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Could not fetch event ID: %v", eventId)})
		return
	}

	_, err = models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch event. Try again later."})
		return
	}

	var updatedEvent models.Event
	err := context.ShouldBindJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse request data."})
		return
	}

}