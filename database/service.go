package main

import (
	"log"
	"net"
	pb "weather/common/services"

	"golang.org/x/net/context"
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
	pb.RegisterDatabaseServiceServer(s, &server{})

	// Register reflection service on gRPC server
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

// server is used to implement pb.DatabaseService
type server struct{}

// Create implements pb.DatabaseService.Create
func (s *server) Create(ctx context.Context, in *pb.Data) (*pb.Result, error) {
	log.Println("Db create rpc call started")

	err := createData(in.Location, in.Values)

	if err != nil {
		return getError(err)
	}

	return &pb.Result{
		Status:  true,
		Message: "Created",
	}, nil
}

// Create implements pb.DatabaseService.Read
func (s *server) Read(ctx context.Context, in *pb.Data) (*pb.Data, error) {
	log.Println("Db read rpc call started")

	data, err := readData(in.Location)
	if err != nil {
		return nil, err
	}

	return &pb.Data{
		Values: data,
	}, nil
}

// Create implements pb.DatabaseService.Create
func (s *server) Update(ctx context.Context, in *pb.Data) (*pb.Result, error) {
	log.Println("Db update rpc call started")

	err := updateData(in.Location, in.Values)

	if err != nil {
		return getError(err)
	}

	return &pb.Result{
		Status:  true,
		Message: "Data updated",
	}, nil
}

// Create implements pb.DatabaseService.Create
func (s *server) Delete(ctx context.Context, in *pb.Data) (*pb.Result, error) {
	log.Println("Db delete rpc call started")

	err := deleteData(in.Location)

	if err != nil {
		return getError(err)
	}

	return &pb.Result{
		Status:  false,
		Message: "Functionality Not implemented",
	}, nil
}

func getError(err error) (*pb.Result, error) {
	return &pb.Result{
		Status:  false,
		Message: err.Error(),
	}, err
}
