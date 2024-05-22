package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	server := gin.Default()

	// GET, POST, PUT, PATCH, DELETE
	server.GET("/", getBase)

	server.GET("/ping", getPing)

	server.GET("/events", getEvents)

	server.Run(":8080") // localhost:8080
}

func getBase(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func getPing(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
