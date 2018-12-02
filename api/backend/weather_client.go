package backend

import (
	"log"
	"time"

	pb "github.com/mungujn/weather-server/weather/services"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	serviceAddress = "localhost:8082"
	name           = "weather_service"
	certificate    = "server.crt"
)

var connection *grpc.ClientConn

// SetRPCConnection sets the active rpc connection
func SetRPCConnection(newConnection *grpc.ClientConn) {
	connection = newConnection
}

//getWeather : get the weather
func getWeather(location, date string) (*pb.Weather, error) {
	log.Println("Getting weather from weather server")

	c := pb.NewWeatherServiceClient(connection)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	weather, err := c.GetWeather(ctx, &pb.LocationAndDate{Location: location, Date: date})

	if err != nil {
		return nil, err
	}
	return weather, nil
}

//GetRPCConnection returns an rpc connection
func GetRPCConnection() (*grpc.ClientConn, error) {
	log.Println("Getting connection to RPC server")
	if connection == nil {
		log.Println("Creating new connection")

		creds, err := credentials.NewClientTLSFromFile(certificate, "")

		if err != nil {
			log.Printf("Failed to load tls certificate: %v", err)
			return nil, err
		}

		connection, err := grpc.Dial(serviceAddress, grpc.WithTransportCredentials(creds))
		if err != nil {
			log.Printf("Failed to connect: %v", err)
			return nil, err
		}
		return connection, nil
	}
	log.Println("Reusing RPC connection")
	return connection, nil
}
