package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/REST/models"
	"github.com/gin-gonic/gin"
)

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
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}
	event.ID = 1
	event.UserID = 1
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusFailedDependency, gin.H{"message": "could not fetch events"})
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})

}
func updateEvent(context *gin.Context) {
	// Parse the event ID from the URL parameter
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}

	// Check if the event exists
	_, err = models.GetEventByID(eventId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Event not found"})
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
	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully"})
	updatedEvent.ID = eventId
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
