package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	// GET, POST, PUT, PATCH, DELETE
	server.GET("/", getBase)

	server.GET("/ping", getPing)

	server.GET("/events", getEvents)

	server.POST("/events", createEvent)

	server.DELETE("/events/:id", deleteEvent)

	server.GET("/events/:id", getEventById)

	server.PUT("/events/:id", updateEvent)

	server.POST("/signup", signup)
}
