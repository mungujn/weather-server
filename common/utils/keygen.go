package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	servercrt = "SERVER_CRT"
	serverkey = "SERVER_KEY"
)

// LoadDotEnvFile reads the .env file and sets the env variables it finds there
func LoadDotEnvFile() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

// WriteServerKeysFromEnv writes server keys from .env
func WriteServerKeysFromEnv() {
	log.Println("Writing server keys from env")
	key := os.Getenv(servercrt)
	filename := "server.crt"
	writeToFile(key, filename)

	key = os.Getenv(serverkey)
	filename = "server.key"
	writeToFile(key, filename)
	log.Println("Done")
}

// WriteClientKeysFromEnv writes client keys from .env
func WriteClientKeysFromEnv() {
	log.Println("Writing client keys from env")
	key := os.Getenv(servercrt)
	filename := "server.crt"
	writeToFile(key, filename)
	log.Println("Done")
}

// writeToFile writes a string to a file
func writeToFile(str, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Failed to create key file: %v", err)
	}
	defer f.Close()

	f.WriteString(str)
}
