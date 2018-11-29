package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"weather/common/logger"
	"weather/common/responses"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.PathPrefix("/api/v1")
	r.HandleFunc("/locations/{location}", LocationHandler).Methods("GET")

	srv := &http.Server{
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		Handler:      r,
	}

	log.Fatal(srv.ListenAndServe())
}

//LocationHandler :
func LocationHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	logger.Info(fmt.Sprintf("/location/%v", vars["category"]))
	responses.RespondOk(w)
}
