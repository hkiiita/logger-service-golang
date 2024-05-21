package main

import (
	"time"
	"github.com/hkiiita/logger-service-golang/logger"
)

func main() {
	//simple logger service
	log := logger.GetLogger()

	log.Log("Message 11")
	time.Sleep(2*time.Second)
	log.Log("Message 22")
	time.Sleep(2*time.Second)

}
