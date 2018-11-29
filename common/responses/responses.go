package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// ok :
type message struct {
	Message string `json:"message"`
}

// RespondWithData : Respond with data
func RespondWithData(w http.ResponseWriter, data interface{}) {
	log.Println("---Responding with data---")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	responseJSON, _ := json.Marshal(data)
	w.Write(responseJSON)
}

// RespondBadRequest : Respond bad request
func RespondBadRequest(w http.ResponseWriter) {
	log.Println("---Responding Bad request---")
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")
	response := message{Message: "No location/date specified"}
	responseJSON, _ := json.Marshal(response)
	w.Write(responseJSON)
}

// RespondInternalServerError : Respond ISE
func RespondInternalServerError(w http.ResponseWriter) {
	log.Println("---Responding ISE---")
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json")
	response := message{Message: "Internal Server Error"}
	responseJSON, _ := json.Marshal(response)
	w.Write(responseJSON)
}
