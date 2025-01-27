package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/REST/models"
	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to the Event Management API",
		"routes": gin.H{
			"events":       "/events (GET, POST, PUT, DELETE)",
			"registration": "/events/:id/register (POST, DELETE)",
			"user":         "/signup, /login",
		},
	})
}
func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch events"})
		return
	}
	context.JSON(http.StatusOK, events)

}
func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil || eventId <= 0 {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	event, err := models.GetEventByID(eventId)
	if err != nil {
		if err.Error() == fmt.Sprintf("event with ID %d not found", eventId) {
			context.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
		}
		return
	}

	context.JSON(http.StatusOK, event)
}
func createEvent(context *gin.Context) {

	var event models.Event
	if err := context.ShouldBindJSON(&event); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}
	// userId := context.GetInt64("userId")

	if err := event.Save(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil || eventId <= 0 {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	// Debug log
	fmt.Printf("Attempting to delete event with ID: %d\n", eventId)

	err = models.DeleteEventByID(eventId)
	if err != nil {
		if err.Error() == fmt.Sprintf("event with ID %d not found", eventId) {
			context.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete event"})
		}
		return
	}

	// Debug log
	fmt.Printf("Event with ID: %d deleted successfully\n", eventId)

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully"})
}
func updateEvent(context *gin.Context) {
	// Parse the event ID from the URL parameter
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	// Retrieve the user ID from the context
	userId := context.GetInt64("userId")

	// Declare the variable 'event' before using it
	var event *models.Event
	event, err = models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
		return
	}

	// Check if the user is authorized to update the event
	if event.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update event"})
		return
	}

	// Bind the request body to the Event struct
	var updatedEvent models.Event
	if err = context.ShouldBindJSON(&updatedEvent); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}

	// Update the event with the parsed ID
	updatedEvent.ID = eventId
	if err = updatedEvent.Update(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update event"})
		return
	}

	// Respond with success
	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully", "event": updatedEvent})
}
