package routes

import (
	"example.com/rest-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// GET, POST, PUT, PATCH, DELETE
	server.GET("/", getBase)
	server.GET("/ping", getPing)

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventById)
	server.POST("/events/:id/register", register)
	server.DELETE("/events/:id/register", deleteEvent)

	server.POST("/login", login)
	server.POST("/signup", signup)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
}
