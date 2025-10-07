package main

import (
	"log"

	"sensors-telemetry-service/db"
	"sensors-telemetry-service/handlers"

	"github.com/gin-gonic/gin"
)

const PORT = ":8082"

func main() {
	// Set up database connection
	dbURL := "port=5432 user=postgres password=postgres dbname=smarthome sslmode=disable host=postgres"
	db.New(dbURL)

	router := gin.Default()

	// API routes
	router.GET("/api/v2/telemetry", func(c *gin.Context) {
		c.JSON(200, gin.H{"telemetry status": "ok"})
	})

	router.PATCH("api/v2/sensors/:id/value", handlers.UpdateSensorValue)

	router.Run(PORT)

	log.Println("Server starting on :8082")
}
