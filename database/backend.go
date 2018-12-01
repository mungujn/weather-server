package main

import (
	"log"

	"github.com/pkg/errors"
)

var data map[string]map[string]string

func setUpDb(host, port string) {
	log.Printf("Db initialisation")
	if data == nil {
		data = make(map[string]map[string]string)
	}
}

// create data in the db
func createData(location string, newData map[string]string) error {
	log.Printf("Creating %s", location)
	data[location] = newData
	return nil
}

// create data in the db
func readData(location string) (map[string]string, error) {
	log.Printf("Reading %s", location)
	data, exists := data[location]

	if !exists {
		log.Println("Not found")
		return nil, errors.New("Not found")
	}
	return data, nil
}

// create data in the db
func updateData(location string, newData map[string]string) error {
	log.Printf("Updating %s", location)
	oldData, exists := data[location]

	if !exists {
		log.Println("Not found")
		return errors.New("Not found")
	}

	for key, value := range newData {
		oldData[key] = value
	}

	return nil
}

// create data in the db
func deleteData(location string) error {
	log.Printf("Deleting %s", location)
	delete(data, location)
	return nil
}
