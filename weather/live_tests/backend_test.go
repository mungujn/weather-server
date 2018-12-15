package live_tests

import (
	"log"
	"os"
	"testing"

	"github.com/mungujn/weather-server/common/utils"
	"github.com/mungujn/weather-server/weather/backend"
)

func TestMain(m *testing.M) {
	log.Println("Running test setup")
	backend.SetUpRedis()

	retCode := m.Run()
	log.Println("Running test cleanup")
	os.Exit(retCode)
}

// TestGetWeather tests the weather data backend
func TestGetWeather(t *testing.T) {
	expecting := make(map[string]string)
	expecting["location"] = "Kampala"
	expecting["temperature"] = "28"
	expecting["date"] = "Today"

	weather, err := backend.GetWeather("kampala", "today")

	if err != nil || !utils.MapsEqualish(expecting, weather.Data) {
		t.Errorf("Failed to get weather, expecting: %v, got: %v, error: %v", expecting, weather, err)
	} else {
		t.Log("Weather retrieved")
	}
}
