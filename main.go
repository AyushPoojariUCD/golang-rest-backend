package main

import (
	"net/http"
	"go-rest-backend/db"
	"go-rest-backend/models"
	"github.com/gin-gonic/gin"
)

// Main Function
func main() {
	
	// Initialize the database
	db.InitDB()

	// Server
	server := gin.Default()

	// GET: / - Health Check Backend
	server.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Golang Backend is running!")
	})

	// GET: /events - Get all events
	server.GET("/events", getEvents)

	// POST: /events - Create a new event
	server.POST("/events", createEvents)

	// Run the server
	server.Run(":8080")
}

// Function: Get all events
func getEvents(c *gin.Context) {
	events := models.GetAllEvents()
	c.JSON(http.StatusOK, events)
}

// Function: Create a new event
func createEvents(c *gin.Context) {
	var event models.Event

	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event.ID = 1
	event.UserID = 1
	event.Save()

	c.JSON(http.StatusCreated, gin.H{
		"message": "Event created successfully",
		"event":   event,
	})
}