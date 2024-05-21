package main

import (
	"time"
	"github.com/hkiiita/logger-service-golang/logger"
)

func main() {
	//simple logger service
	log := logger.GetLogger()
	log.InputChan <- "Message 1"
	time.Sleep(5 * time.Second)
	log.InputChan <- "Message 2"
}
