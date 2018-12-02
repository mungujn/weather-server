package backend

import (
	"log"

	"github.com/pkg/errors"
)

var data map[string]map[string]string

// SetUpDb sets up the database
func SetUpDb(host, port string) {
	log.Printf("Initialising database engine")
	if data == nil {
		data = make(map[string]map[string]string)
	}
	log.Printf("Done")
}

// CreateData creates data in the db
func CreateData(location string, newData map[string]string) error {
	log.Printf("Creating %s", location)
	data[location] = newData
	return nil
}

// ReadData reads data in the db
func ReadData(location string) (map[string]string, error) {
	log.Printf("Reading %s", location)
	data, exists := data[location]

	if !exists {
		log.Println("Not found")
		return nil, errors.New("Not found")
	}
	return data, nil
}

// UpdateData updates data in the db
func UpdateData(location string, newData map[string]string) error {
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

// DeleteData deletes data in the db
func DeleteData(location string) error {
	log.Printf("Deleting %s", location)
	delete(data, location)
	return nil
}
