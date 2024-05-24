package main

import (
	"strconv"
	"time"

	"github.com/hkiiita/logger-service-golang/logger"
)

func main() {
	// Initializing the logger
	log := logger.GetLogger()
	// Write an INFO message
	log.Log("[INFO]: Starting the application...")
	// Simulate a task
	for i := 0; i <= 3; i++ {
		log.Log("[INFO]: Executing task ... " + strconv.Itoa(i))
		time.Sleep(2 * time.Second)
	}
	// Write a WARNING message
	log.Log("[WARNING]: Configuration file missing... Generating the default configuration")
	// Simulate another task
	for z := 0; z <= 3; z++ {
		log.Log("[INFO]: Generating file number : " + strconv.Itoa(z))
		time.Sleep(2 * time.Second)
	}
	// Critical error
	log.Log("[ERROR]: A critical error has occurred, application stopped ")
	// End program
	log.Log("[ERROR]: Application stopped")
}
