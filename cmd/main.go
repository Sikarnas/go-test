package main

import (
	"test/internal/api"
	"test/internal/api/event"
	"test/internal/db"
)

func main() {
	dbClient := db.New()
	apiClient := api.New()
	eventController := event.NewController(dbClient, apiClient)

	apiClient.POST("/v1/event", eventController.Create)

	apiClient.Logger.Fatal(apiClient.Start(":1323"))
}
