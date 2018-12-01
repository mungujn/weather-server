package main

import (
	"log"
	"time"

	pb "github.com/mungujn/weather-server/database/services"

	"golang.org/x/net/context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	dbAddress = "localhost:8082"
)

var connection *grpc.ClientConn

func setUpDBClient(newConnection *grpc.ClientConn) {
	connection = newConnection
}

//createData creates data in the db
func createData(location string, data map[string]string) error {
	log.Println("Creating data")

	c := pb.NewDatabaseServiceClient(connection)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	result, err := c.Create(ctx, &pb.Data{Location: location, Values: data})

	if err != nil {
		return err
	}
	log.Println("Data created: ", result)
	return nil
}

// readData reads data from the db
func readData(location string) (map[string]string, error) {
	log.Println("Reading data")

	c := pb.NewDatabaseServiceClient(connection)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	result, err := c.Read(ctx, &pb.Data{Location: location, Values: nil})

	if err != nil {
		return nil, err
	}

	log.Println("Data read: ", result)
	return result.Values, nil
}

// updateData updates data
func updateData(location string, data map[string]string) error {
	log.Println("Updating data")

	c := pb.NewDatabaseServiceClient(connection)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	result, err := c.Update(ctx, &pb.Data{Location: location, Values: data})

	if err != nil {
		return err
	}

	log.Println("Updated: ", result)
	return nil
}

// deleteData deletes data
func deleteData(location string) error {
	log.Println("Deleting data")

	c := pb.NewDatabaseServiceClient(connection)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	result, err := c.Delete(ctx, &pb.Data{Location: location, Values: nil})

	if err != nil {
		return err
	}
	log.Println("Deleted: ", result)
	return nil
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

		connection, err := grpc.Dial(dbAddress, grpc.WithTransportCredentials(creds))
		if err != nil {
			log.Printf("Failed to connect: %v", err)
			return nil, err
		}
		return connection, nil
	}
	log.Println("Reusing RPC connection")
	return connection, nil
}
