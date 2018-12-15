package main

import (
	"log"
	"net"

	"github.com/mungujn/weather-server/weather/backend"
	pb "github.com/mungujn/weather-server/weather/services"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

const (
	port        = ":8081"
	certificate = "/data/weather.crt"
	privateKey  = "/data/weather.key"
)

func main() {
	log.Println("Weather RPC server initialising...")

	// set up the db client that this server needs
	backend.SetUpRedis()

	// Create the channel to listen on
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Add TLS credentials
	creds, err := credentials.NewServerTLSFromFile(certificate, privateKey)
	if err != nil {
		log.Fatalf("Failed to load TLS keys: %v", err)
	}

	log.Printf("TCP Listening on port: %v", port)

	s := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterWeatherServiceServer(s, &backend.Server{})

	// Register reflection service on gRPC server
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
