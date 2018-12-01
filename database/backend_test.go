package main

import (
	"testing"
)

// TestCreate tests the create function
func TestCreate(t *testing.T) {
	setUpDb("localhost", "")

	location := "kampala/today"
	data := make(map[string]string)
	data["temperature"] = "28"

	err := createData(location, data)

	if err != nil {
		t.Errorf("Failed to save %v at %s", data, location)
	}
	t.Log("Created")
}

// TestRead tests the read function
func TestRead(t *testing.T) {
	setUpDb("localhost", "")

	location := "kampala/today"
	data := make(map[string]string)
	data["temperature"] = "28"

	createData(location, data)

	savedData, err := readData(location)

	if err != nil || !mapsEqual(savedData, data) {
		t.Errorf("Failed to read %s, expecting: %v got: %v %v", location, data, savedData, err)
	}
	t.Log("Read")
}

// TestUpdate tests the update function
func TestUpdate(t *testing.T) {
	setUpDb("localhost", "")

	location := "kampala/today"
	data := make(map[string]string)
	data["temperature"] = "28"

	createData(location, data)

	err := updateData(location, data)

	if err != nil {
		t.Errorf("Failed to update %v at %s, error: %v", data, location, err)
	}
	t.Log("Updated")
}

// TestDelete tests the delete function
func TestDelete(t *testing.T) {
	setUpDb("localhost", "")

	location := "kampala/today"

	err := deleteData(location)

	if err != nil {
		t.Errorf("Failed to delete %s, error: %v", location, err)
	}
	t.Log("Deleted")
}

// check if two string maps are equal
func mapsEqual(m1, m2 map[string]string) bool {
	l1 := 0

	for k1, v1 := range m1 {
		l1++
		v2, e := m2[k1]
		if !e || v1 != v2 {
			return false
		}
	}

	l2 := len(m2)

	if l1 != l2 {
		return false
	}

	return true
}
