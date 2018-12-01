package backend

import (
	"log"
	"net/http"
	"strings"

	"github.com/mungujn/weather-server/common/responses"

	"github.com/gorilla/mux"
)

// CreateRouter creates the mux router
func CreateRouter() *mux.Router {
	r := mux.NewRouter()

	r.Handle("/weather", dataValidationMiddleware(http.HandlerFunc(weatherHandler)))

	return r
}

//weatherHandler : handler function for retrieving past weather
func weatherHandler(w http.ResponseWriter, r *http.Request) {
	location, date, _ := getLocationAndDateFromURL(r)
	p := strings.Repeat("-", 10)
	log.Printf("%s %v %s", p, r.URL, p)

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
