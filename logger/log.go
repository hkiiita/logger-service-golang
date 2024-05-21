package logger

import (
	"fmt"
	"os"
	"time"
)

type Logger struct {
	InputChan chan string
	LogFile   *os.File
}

var logService *Logger

func init() {
	// Initializes the logger service
	println("Initializing logger.....")
	inputFile, err := os.OpenFile("dummyAccess.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	writeChannel := make(chan string)
	logService = &Logger{
		InputChan: writeChannel,
		LogFile:   inputFile,
	}
	go func(logService *Logger) {
		for {
			select {
			case message := <-logService.InputChan:
				//defer logService.LogFile.Close()
				_, err := logService.LogFile.WriteString(message + "\n")
				if err != nil {
					fmt.Printf("Error writing to file %+v", err)
				}
			default:
			}
		}
	}(logService)
}

func (logService *Logger)Log(message string){
	
	date := time.Now()
	message = date.Format("2006-01-02T15:04:05Z07:00") + " " + message + "\n"
	logService.InputChan <- message	
	
}

func GetLogger() *Logger {
	if logService != nil {
		return logService
	}
	return nil
}
