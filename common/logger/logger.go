package logger

import (
	"log"
)

// SetUp : Set up the logger
func SetUp(){
	log.SetPrefix("App")
}

//Info : Logs an info message
func Info(message string){
	log.Println(message)
}