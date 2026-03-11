package routes

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"go-rest-backend/models"
	"go-rest-backend/utils"
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

	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Authorization header missing",
		})
		return
	}

	// Split "Bearer <token>"
	tokenParts := strings.Split(authHeader, " ")

	if len(tokenParts) != 2 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid authorization format",
		})
		return
	}

	tokenString := tokenParts[1]

	parsedToken, err := utils.VerifyToken(tokenString)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token",
		})
		return
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token claims",
		})
		return
	}

	userId := int(claims["userId"].(float64))

	var event models.Event

	err = c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	event.UserID = userId

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

// Delete Event
func deleteEvent(c *gin.Context) {

	id := c.Param("id")

	err := models.DeleteEvent(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not delete event",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Event deleted successfully",
	})
}