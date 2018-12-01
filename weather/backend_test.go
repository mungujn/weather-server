package main

import (
	"log"
	"os"
	"testing"

	"github.com/mungujn/weather-server/common/utils"

	"google.golang.org/grpc"
)

func TestMain(m *testing.M) {
	log.Println("Running test setup")
	connection, _ := getRPCConnection()
	setUpDBClient(connection)

	retCode := m.Run()
	log.Println("Running test cleanup")
	connection.Close()
	os.Exit(retCode)
}

// TestGetWeather tests the weather data backend
func TestGetWeather(t *testing.T) {
	expecting := make(map[string]string)
	expecting["location"] = "Kampala"
	expecting["temperature"] = "28"
	expecting["date"] = "Today"

	weather, err := getWeather("kampala", "today")

	if err != nil || !utils.MapsEqualish(expecting, weather.Data) {
		t.Errorf("Failed to get weather, expecting: %v, got: %v, error: %v", expecting, weather, err)
	} else {
		t.Log("Weather retrieved")
	}
}

func setUpTests(t *testing.T) *grpc.ClientConn {
	return connection
}
