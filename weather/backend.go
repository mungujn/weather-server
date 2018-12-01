package main

import (
	pb "github.com/mungujn/weather-server/weather/services"
)

// getWeather retrieves weather from an API
func getWeather(location, date string) (*pb.Weather, error) {
	return &pb.Weather{
		Location: "Kampala",
		Date: "Today",
		Temperature: 28,
	}, nil
}