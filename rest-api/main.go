package main

import (
	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()

	// GET, POST, PUT, PATCH, DELETE
	server.GET("/", getBase)

	server.GET("/ping", getPing)

	server.GET("/events", getEvents)

	server.POST("/events", createEvent)

	server.Run(":8080") // localhost:8080
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event) // works like the scan, pass a pointer
	// internally gin looks at incoming body and stores in the var
	// data must match types of the model
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	event.ID = 1
	event.UserID = 1
	event.Save()

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func getBase(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func getPing(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}