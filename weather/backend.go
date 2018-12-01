package main

import (
	pb "github.com/mungujn/weather-server/weather/services"
)

const (
	cacheTTL = 60 * 60 * 10
)

// getWeather retrieves weather
func getWeather(location, date string) (*pb.Weather, error) {
	data, err := readData(location + "/" + date)
	if err != nil {
		return nil, err
	}

	return &pb.Weather{
		Location:    data["location"],
		Date:        data["date"],
		Temperature: data["temperature"],
	}, nil
}

func readWeatherFromDb(location, date string) (*pb.Weather, error) {

}
