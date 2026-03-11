package main

import (
	"net/http"

	"go-rest-backend/db"
	"go-rest-backend/models"

	"github.com/gin-gonic/gin"
)

func main() {

	// Initialize database
	db.InitDB()

	// Create server
	server := gin.Default()

	// Health check route
	server.GET("/", healthCheck)

	// Event routes
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	// Start server
	server.Run(":8080")
}

// Health check
func healthCheck(c *gin.Context) {
	c.String(http.StatusOK, "Golang Backend is running!")
}

// Get all events
func getEvents(c *gin.Context) {

	events, err := models.GetAllEvents()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not retrieve events",
		})
		return
	}

	c.JSON(http.StatusOK, events)
}

// Create new event
func createEvent(c *gin.Context) {

	var event models.Event

	err := c.ShouldBindJSON(&event)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Temporary user assignment
	event.UserID = 1

	err = event.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not create event",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Event created successfully",
		"event":   event,
	})
}