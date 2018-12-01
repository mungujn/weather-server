package main

import (
	"log"
	"net/http"
	"github.com/mungujn/weather/common/responses"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// createRouter : create router
func createRouter() *mux.Router {
	r := mux.NewRouter()

	r.Handle("/weather", dataValidationMiddleware(http.HandlerFunc(WeatherHandler)))

	return r
}

func main() {
	router := createRouter()
	port := "8080"

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET"})
	
	log.Printf("Server listening on port %v", port)
	log.Fatal(http.ListenAndServe(":"+port, handlers.CORS(allowedOrigins, allowedMethods)(router)))
}

//WeatherHandler : handler function for retrieving past weather
func WeatherHandler(w http.ResponseWriter, r *http.Request) {
	location, date, _ := getLocationAndDateFromURL(r)
	log.Printf("---%v---", r.URL)

	wth, err := getWeather(location, date)

	if err != nil {
		responses.RespondInternalServerError(w, err)
	}
	responses.RespondWithData(w, wth)
}

//dataValidationMiddleware : validate that the location has been defined
func dataValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _, found := getLocationAndDateFromURL(r)

		if found {
			next.ServeHTTP(w, r)
		} else {
			log.Println("Location/Date missing from request")
			responses.RespondBadRequest(w)
			return
		}
	})
}

// getLocationAndDateFromURL : get Location From URL
func getLocationAndDateFromURL(r *http.Request) (string, string, bool) {
	location, locationFound := r.URL.Query()["location"]
	date, dateFound := r.URL.Query()["date"]

	if !locationFound || !dateFound || len(location) != 1 || len(date) != 1 {
		return "", "", false
	}

	return location[0], date[0], true
}
