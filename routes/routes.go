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

	server.POST("/events", createEvent)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)

	// Users routes
	server.GET("/users", getUsers)
	server.GET("/users/:id", getUserByID)
	server.POST("/signup", signUp)
	server.POST("/login", login)
	server.PUT("/users/:id", updateUser)
	server.DELETE("/users/:id", deleteUser)
}