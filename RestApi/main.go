package main

import (
	"net/http"

	"travel_api/RestApi/database"
	"travel_api/RestApi/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to DB and pass to models
	database.Connect()
	models.DB = database.DB

	server := gin.Default()
	server.GET("/events", getEvents)
	server.POST("/events", CreateEvents)
	server.Run(":8080")
}
func getEvents(context *gin.Context) {
	events, err := models.GetEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to fetch events"})
		return
	}
	context.JSON(http.StatusOK, events) // it will convert the data in the JSON formate
}
func CreateEvents(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to parse the request "})
		return
	}
	event.UserId = 1
	if err := event.Save(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not save event"})
		return

	}
	context.JSON(http.StatusCreated, gin.H{"message": "successfully created the event", "event": event})

}
