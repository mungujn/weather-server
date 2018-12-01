package main

import (
	"log"
	"net"

	pb "github.com/mungujn/weather-server/weather/services"
	"github.com/mungujn/weather-server/weather/backend"
	
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

const (
	port        = ":8082"
	certificate = "server.crt"
	privateKey  = "server.key"
)

func main() {
	log.Println("Weather RPC server initialising...")
	// set up the db rpc client that this server needs
	connection, err := backend.GetRPCConnection()
	defer connection.Close()

	if err != nil {
		log.Fatalf("Failed to get RPC connection to database service")
	}

	backend.SetRPCConnection(connection)

	// Create the channel to listen on
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listed: %v", err)
	}
	log.Printf("TCP Listening on port: %v", port)

	// Add TLS credentials
	creds, err := credentials.NewServerTLSFromFile(certificate, privateKey)
	if err != nil {
		log.Fatalf("Failed to load TLS keys: %v", err)
	}

	s := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterWeatherServiceServer(s, &backend.Server{})

	// Register reflection service on gRPC server
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}