package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"go-rest-backend/models"
)

// Functions => Handlers for each route

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

// Get event by ID
func getEventByID(c *gin.Context) {

	id := c.Param("id")	

	event, err := models.GetEventByID(id)

	if err != nil {	
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Event not found",
		})
		return
	}
	c.JSON(http.StatusOK, event)
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

// Update Event
func updateEvent(c *gin.Context) {

	id := c.Param("id")

	// find existing event
	event, err := models.GetEventByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Event not found",
		})
		return
	}

	var updatedEvent models.Event

	err = c.ShouldBindJSON(&updatedEvent)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// keep same ID
	updatedEvent.ID = event.ID

	err = updatedEvent.Update()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not update event",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Event updated successfully",
		"event":   updatedEvent,
	})
}