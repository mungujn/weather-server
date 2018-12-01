package main

import (
	"log"
	"time"
	pb "github.com/mungujn/weather-server/weather/services"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	address     = "localhost:8081"
	name        = "weather_service"
	certificate = "server.crt"
)

var connection *grpc.ClientConn

//getWeather : get the weather
func getWeather(location, date string) (*pb.Weather, error) {
	log.Println("Getting weather from weather server")
	connection, err := getRPCConnection()
	defer connection.Close()

	if err != nil {
		log.Println("Failed to get RPC connection")
		return nil, err
	}

	c := pb.NewWeatherServiceClient(connection)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	weather, err := c.GetWeather(ctx, &pb.LocationAndDate{Location: location, Date: date})

	if err != nil {
		return nil, err
	}
	return weather, nil
}

//getRPCConnection returns an rpc connection
func getRPCConnection() (*grpc.ClientConn, error) {
	log.Println("Getting connection to RPC server")
	if connection == nil {
		log.Println("Creating new connection")

		creds, err := credentials.NewClientTLSFromFile(certificate, "")

		if err != nil {
			log.Printf("Failed to load tls certificate: %v", err)
			return nil, err
		}

		connection, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
		if err != nil {
			log.Printf("Failed to connect: %v", err)
			return nil, err
		}
		return connection, nil
	}
	log.Println("Reusing RPC connection")
	return connection, nil
}
