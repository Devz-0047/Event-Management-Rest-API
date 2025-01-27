package routes

import (
	"example.com/REST/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// Event-related routes
	server.GET("/", Home)
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.POST("/events", middleware.Authenticate, createEvent)
	server.PUT("/events/:id", middleware.Authenticate, updateEvent)
	server.DELETE("/events/:id", middleware.Authenticate, deleteEvent)
	server.POST("/events/:id/register", middleware.Authenticate, registerForEvent) // Fixed route for registration
	server.DELETE("/events/:id/register", middleware.Authenticate, cancelRegistration)

	// User-related routes
	server.POST("/signup", signup)
	server.POST("/login", login)
}
