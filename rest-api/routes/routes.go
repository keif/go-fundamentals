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

	server.POST("/login", login)
	server.POST("/signup", signup)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistrationForEvent)

	// list registrations?
}
