package main

import (
	"log"
	"os"
	"testing"

	pb "github.com/mungujn/weather-server/weather/services"
	"google.golang.org/grpc"
)

func TestMain(m *testing.M) {
	log.Println("Running before")
	retCode := m.Run()
	log.Println("Running aftter")
	os.Exit(retCode)
}

// TestGetWeather tests the weather data backend
func TestGetWeather(t *testing.T) {
	connection, err := getRPCConnection()
	defer connection.Close()
	setUpDBClient(connection)
	
	expecting := &pb.Weather{Location: "Kampala", Date: "Today", Temperature: 28}
	weather, err := getWeather("kampala", "today")

	if err != nil || !weatherEquals(weather, expecting) {
		t.Errorf("Failed to get weather, expecting: %v, got: %v, error: %v", expecting, weather, err)
	}
	t.Log("Weather retrieved")
}

// weatherEquals tests for weather equality
func weatherEquals(w1, w2 *pb.Weather) bool {
	if w1.Location == w2.Location && w1.Date == w2.Date && w1.Temperature == w2.Temperature {
		return true
	}
	return false
}

func setUpTests(t *testing.T) *grpc.ClientConn {

	return connection
}
