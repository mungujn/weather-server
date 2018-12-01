package main

import (
	"log"

	"github.com/mungujn/weather-server/common/utils"

	"strconv"

	pb "github.com/mungujn/weather-server/weather/services"
)

const (
	cacheTTL = 60 * 60 * 10
)

// getWeather retrieves weather
func getWeather(location, date string) (*pb.Weather, error) {
	log.Printf("Getting %s's weather in %s", date, location)
	// read weather from db
	path := location + "/" + date
	data, err := readData(path)

	// if the data doesnt exist in the db, get it fresh from the provider
	if err != nil {
		log.Println("Weather data not found in DB")
		return refreshWeatherData(location, date)
	}

	// if using db weather data check if weather it is too old
	now := utils.Timestamp()
	then, err := strconv.ParseInt(data["retrieved"], 10, 64)
	if err != nil {
		return nil, err
	}
	diff := now - then
	if diff > cacheTTL {
		return refreshWeatherData(location, date)
	}
	return &pb.Weather{
		Data: data,
	}, nil
}

func refreshWeatherData(location, date string) (*pb.Weather, error) {
	data, err := readWeatherFromProvider(location, date)
	if err != nil {
		return nil, err
	}
	return &pb.Weather{
		Data: data,
	}, nil
}

func readWeatherFromProvider(location, date string) (map[string]string, error) {
	log.Println("Reading weather data from provider")
	data := make(map[string]string)
	data["location"] = "Kampala"
	data["temperature"] = "28"
	data["date"] = "Today"

	// mark retrieval time for caching logic
	data["retrieved"] = strconv.FormatInt(utils.Timestamp(), 10)

	// cache the fresh weather data in the db
	go func() {
		log.Println("Updating db weather data in thread")
		createData(location+"/"+date, data)
	}()

	return data, nil
}
