package main

import (
	"fmt"
	"log"
	"net/http"

	"temperature-api/data"

	"github.com/gin-gonic/gin"
)

const PORT = ":8081"

func main() {

	log.Println("Start temperature-api . . .")

	// Initialize router
	router := gin.Default()

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// Get temperature by location
	// http://localhost:8081/temperature?location="Living Rooom"
	router.GET("/temperature", func(c *gin.Context) {
		location := c.Query("location")
		// Generate random temperature data based on location
		data := data.GetTemperatureByLocation(location)

		c.JSON(http.StatusOK, data)
	})

	router.GET("/sensorId", func(c *gin.Context) {
		sensorId := c.Query("sensorId")
		// Generate random temperature data based on sensorId
		data := data.GetTemperatureBySensorId(sensorId)

		c.JSON(http.StatusOK, data)
	})

	router.GET("/temperature/:id", func(c *gin.Context) {
		sensorId := c.Query("sensorId")
		// Generate random temperature data based on sensorId
		data := data.GetTemperatureBySensorId(sensorId)

		c.JSON(http.StatusOK, data)
	})

	log.Println("Visit http://localhost:8081/health to see if it works.")
	log.Println("Visit http://localhost:8081/temperature?location=Kitchen to get temperature from the kitchen.")
	log.Println("Visit http://localhost:8081/temperature?sensorId=2 to get temperature on the sensor 2.")

	// Start the server in a goroutine
	router.Run(PORT)

	log.Println("Server exited properly")
	fmt.Scanln()
}
