package main

import (
	"log"
	"net/http"

	"github.com/mungujn/weather-server/api/backend"

	"github.com/gorilla/handlers"
)

const (
	servercrt = "/data/wapi.crt"
	serverkey = "/data/wapi.key"
)

func main() {
	log.Println("Weather REST API server initialising...")
	router := backend.CreateRouter()
	port := "8080"

	connection, err := backend.GetRPCConnection()
	defer connection.Close()

	if err != nil {
		log.Fatalf("Failed to get RPC connection to weather service")
	}

	backend.SetRPCConnection(connection)

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET"})

	log.Printf("Server listening on port %v", port)
	log.Fatal(http.ListenAndServeTLS(":"+port, servercrt, serverkey, handlers.CORS(allowedOrigins, allowedMethods)(router)))
}
