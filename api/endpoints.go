package main

import (
	"log"
	"net/http"
	"time"
	"weather/common/responses"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.PathPrefix("/api/v1")
	r.HandleFunc("/past-weather", PastWeatherHandler).Queries("location", "location").Methods("GET")
	r.HandleFunc("/current-weather", CurrentWeatherHandler).Queries("location", "location").Methods("GET")
	r.HandleFunc("/future-weather", FutureWeatherHandler).Queries("location", "location").Methods("GET")

	srv := &http.Server{
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		Handler:      r,
	}

	log.SetPrefix("/api/v1")
	log.Fatal(srv.ListenAndServe())
}

// weather : weather data
type weather struct {
	Location    string `json:"location"`
	Temparature int    `json:"temparature"`
}

//PastWeatherHandler : handler function for retrieving past weather
func PastWeatherHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Printf("/locations/%v", params["location"])

	wth := weather{
		Location:    params["location"],
		Temparature: 28,
	}

	responses.RespondWithData(w, wth)
}


//CurrentWeatherHandler : handler function for retrieving the current weather
func CurrentWeatherHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Printf("/locations/%v", params["location"])

	wth := weather{
		Location:    params["location"],
		Temparature: 28,
	}

	responses.RespondWithData(w, wth)
}


//FutureWeatherHandler : handler function for retrieving future weather forecasts
func FutureWeatherHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	log.Printf("/locations/%v", params["location"])

	wth := weather{
		Location:    params["location"],
		Temparature: 28,
	}

	responses.RespondWithData(w, wth)
}
