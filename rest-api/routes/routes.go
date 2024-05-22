package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	// GET, POST, PUT, PATCH, DELETE
	server.GET("/", getBase)

	server.GET("/ping", getPing)

	server.GET("/events", getEvents)

	server.GET("/events/:id", getEventById)

	server.POST("/events", createEvent)

	server.PUT("/events/:id", updateEvent)
}
