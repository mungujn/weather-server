package tests

import (
	"testing"

	"github.com/mungujn/weather-server/common/utils"
	db "github.com/mungujn/weather-server/database/backend"
)

// TestCreate tests the create function
func TestCreate(t *testing.T) {
	db.SetUpDb("localhost", "")

	location := "kampala/today"
	data := make(map[string]string)
	data["temperature"] = "28"

	err := db.CreateData(location, data)

	if err != nil {
		t.Errorf("Failed to save %v at %s", data, location)
	}
	t.Log("Created")
}

// TestRead tests the read function
func TestRead(t *testing.T) {
	db.SetUpDb("localhost", "")

	location := "kampala/today"
	data := make(map[string]string)
	data["temperature"] = "28"

	db.CreateData(location, data)

	savedData, err := db.ReadData(location)

	if err != nil || !utils.MapsEqual(savedData, data) {
		t.Errorf("Failed to read %s, expecting: %v got: %v %v", location, data, savedData, err)
	}
	t.Log("Read")
}

// TestUpdate tests the update function
func TestUpdate(t *testing.T) {
	db.SetUpDb("localhost", "")

	location := "kampala/today"
	data := make(map[string]string)
	data["temperature"] = "28"

	db.CreateData(location, data)

	err := db.UpdateData(location, data)

	if err != nil {
		t.Errorf("Failed to update %v at %s, error: %v", data, location, err)
	}
	t.Log("Updated")
}

// TestDelete tests the delete function
func TestDelete(t *testing.T) {
	db.SetUpDb("localhost", "")

	location := "kampala/today"

	err := db.DeleteData(location)

	if err != nil {
		t.Errorf("Failed to delete %s, error: %v", location, err)
	}
	t.Log("Deleted")
}
