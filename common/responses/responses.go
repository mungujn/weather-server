package responses

import (
	"encoding/json"
	"net/http"
)

// ok :
type message struct {
	Message string `json:"message"`
}

// RespondOk : Respond Ok
func RespondOk(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	response := message{Message: "success"}
	responseJSON, _ := json.Marshal(response)
	w.Write(responseJSON)
}

// RespondWithData : Respond with data
func RespondWithData(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	responseJSON, _ := json.Marshal(data)
	w.Write(responseJSON)
}

// RespondInternalServerError : Respond ISE
func RespondInternalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json")
	response := message{Message: "Internal Server Error"}
	responseJSON, _ := json.Marshal(response)
	w.Write(responseJSON)
}
