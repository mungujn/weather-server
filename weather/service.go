package main 

import (
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc"
	"net"
	"log"
	"golang.org/x/net/context"
	pb "weather/common/services"
)

const (
	port = ":8081"
	certificate = "server.crt"
	privateKey = "server.key"
)

// server is used to implement the pb.GetWeather service
type server struct{}

// ReturnWeather implements pb.GetWeather
func (s *server) GetWeather(ctx context.Context, in *pb.LocationAndDate) (*pb.Weather, error){
	log.Println("Weather request received")
	return &pb.Weather{
		Location: "Kampala",
		Date: "Today",
		Temperature: 28,
	}, nil
}

func main(){
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
	pb.RegisterWeatherServiceServer(s, &server{})

	// Register reflection service on gRPC server
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}