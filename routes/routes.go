package routes

import (	
	"github.com/gin-gonic/gin"
)

// Routes registration : Endpoint definitions
func RegisterRoutes(server *gin.Engine) {
	// Health check route
	server.GET("/", healthCheck)

	// Event routes
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventByID)

	// Create event route
	server.POST("/events", createEvent)
	// Update event route
	server.PUT("/events/:id", updateEvent)
}

