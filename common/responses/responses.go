package responses

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

const (
	ap = 10
)
// ok :
type message struct {
	Message string `json:"message"`
}

// RespondWithData : Respond with data
func RespondWithData(w http.ResponseWriter, data interface{}) {
	p := strings.Repeat("-", ap)
	log.Printf("%s Responding with data %s", p, p)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	responseJSON, _ := json.Marshal(data)
	w.Write(responseJSON)
}

// RespondBadRequest : Respond bad request
func RespondBadRequest(w http.ResponseWriter) {
	p := strings.Repeat("-", ap)
	log.Printf("%s Responding Bad request %s", p, p)
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")
	response := message{Message: "No location/date specified"}
	responseJSON, _ := json.Marshal(response)
	w.Write(responseJSON)
}

// RespondInternalServerError : Respond ISE
func RespondInternalServerError(w http.ResponseWriter, err error) {
	log.Println(err)
	p := strings.Repeat("-", ap)
	log.Printf("%s Responding ISE %s", p, p)
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json")
	response := message{Message: "Internal Server Error"}
	responseJSON, _ := json.Marshal(response)
	w.Write(responseJSON)
}
