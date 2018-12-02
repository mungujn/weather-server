package main

import (
	"log"
	"net"

	"github.com/mungujn/weather-server/common/utils"

	"github.com/mungujn/weather-server/database/backend"
	pb "github.com/mungujn/weather-server/database/services"

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
	log.Println("Database RPC server starting...")
	utils.LoadDotEnvFile()
	utils.WriteServerKeysFromEnv()

	// Create the channel to listen on
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	backend.SetUpDb("host", "port")

	// Add TLS credentials
	creds, err := credentials.NewServerTLSFromFile(certificate, privateKey)
	if err != nil {
		log.Fatalf("Failed to load TLS keys: %v", err)
	}

	log.Printf("TCP Listening on port: %v", port)

	s := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterDatabaseServiceServer(s, &backend.Server{})

	// Register reflection service on gRPC server
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
