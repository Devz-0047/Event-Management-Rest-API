package routes

import (
	"net/http"
	"strconv"

	"example.com/REST/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	// Get the user ID from the context
	userIdValue, exists := context.Get("userId")
	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "User ID not found in context"})
		return
	}
	userId, ok := userIdValue.(int64)
	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "User ID is not valid"})
		return
	}

	// Parse the event ID from the URL parameter
	eventIdStr := context.Param("id")
	eventId, err := strconv.ParseInt(eventIdStr, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}

	// Fetch the event by ID
	event, err := models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}

	// Register the user for the event
	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Registered"})
}

func cancelRegistration(context *gin.Context) {
	// Get the user ID from the context
	userIdValue, exists := context.Get("userId")
	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "User ID not found in context"})
		return
	}
	userId, ok := userIdValue.(int64)
	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "User ID is not valid"})
		return
	}

	// Parse the event ID from the URL parameter
	eventIdStr := context.Param("id")
	eventId, err := strconv.ParseInt(eventIdStr, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID."})
		return
	}

	// Cancel registration
	var event models.Event
	event.ID = eventId
	err = event.CancelRegistration(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registration"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Registration cancelled"})
}
