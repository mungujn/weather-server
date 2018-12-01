package main

import (
	"log"
	"net"

	pb "github.com/mungujn/weather-server/database/services"
	"github.com/mungujn/weather-server/database/backend"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

const (
	port        = ":8081"
	certificate = "server.crt"
	privateKey  = "server.key"
	ap          = 10
)

func main() {
	// Create the channel to listen on
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listed: %v", err)
	}
	log.Println("Database RPC server starting...")
	backend.SetUpDb("host", "port")
	log.Printf("TCP Listening on port: %v", port)

	// Add TLS credentials
	creds, err := credentials.NewServerTLSFromFile(certificate, privateKey)
	if err != nil {
		log.Fatalf("Failed to load TLS keys: %v", err)
	}

	s := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterDatabaseServiceServer(s, &backend.Server{})

	// Register reflection service on gRPC server
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
