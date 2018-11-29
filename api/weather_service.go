package main

// weather : weather data
type weather struct {
	Location    string `json:"location"`
	Date        string `json:"date"`
	Temparature int    `json:"temparature"`
}

//getWeather : get the weather
func getWeather(location, date string) weather {
	wth := weather{
		Location:    location,
		Date:        date,
		Temparature: 28,
	}
	return wth
}
