package data

import (
	"math/rand"
)

type TemperatureDto struct {
	Value    float64 `json:"value"`
	SensorId string  `json:"sensor_id"`
	Location string  `json:"location"`
}

func GetTemperature(sensorId string, location string) TemperatureDto {

	// If no sensor ID is provided, generate one based on location
	if sensorId == "" {
		switch location {
		case "Living Room":
			sensorId = "1"
		case "Bedroom":
			sensorId = "2"
		case "Kitchen":
			sensorId = "3"
		default:
			sensorId = "0"
		}
	} else if location == "" {
		switch sensorId {
		case "1":
			location = "Living Room"
		case "2":
			location = "Bedroom"
		case "3":
			location = "Kitchen"
		default:
			location = "Unknown"
		}
	} else {
		sensorId = "0"
		location = "Unknown"
	}

	return TemperatureDto{
		Value:    10 + float64(rand.Intn(20)),
		SensorId: sensorId,
		Location: location,
	}
}

func GetTemperatureByLocation(location string) TemperatureDto {
	sensorId := "-1"

	switch location {
	case "Living Room":
		sensorId = "1"
	case "Bedroom":
		sensorId = "2"
	case "Kitchen":
		sensorId = "3"
	default:
		sensorId = "0"
	}

	return TemperatureDto{
		Value:    10 + float64(rand.Intn(20)),
		SensorId: sensorId,
		Location: location,
	}
}

func GetTemperatureBySensorId(sensorId string) TemperatureDto {
	location := ""

	switch sensorId {
	case "1":
		location = "Living Room"
	case "2":
		location = "Bedroom"
	case "3":
		location = "Kitchen"
	default:
		location = "Unknown"
	}

	return TemperatureDto{
		Value:    10 + float64(rand.Intn(20)),
		SensorId: sensorId,
		Location: location,
	}
}
