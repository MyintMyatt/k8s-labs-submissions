package main

import (
	"errors"
	"log"
	"os"

	"github.com/google/uuid"
)

const FILE_PATH = "./log.txt"

func main() {
	log.Println("Log writer is running......")

	filePath := os.Getenv("LOG_DIR")

	if filePath == "" {
		filePath = FILE_PATH
	}

	randomStr, err := uuid.NewRandom()
	if err != nil {
		log.Fatalf("File writing error : %v \n", err)
	}

	var logFile *os.File
	if !checkFileExists(filePath) {
		logFile, err = os.Create(filePath)
		if err != nil {
			log.Fatalf("File creation error : %v \n", err)
		}
	} else {
		logFile, err = os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalf("File opening error : %v \n", err)
		}
	}
	defer logFile.Close()
	// os.WriteFile(filePath, []byte(randomStr.String()), 0644)
	logFile.WriteString(randomStr.String() + "\n")
}

func checkFileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err == nil {
		return true
	}

	if errors.Is(err, os.ErrNotExist) {
		return false
	}

	return false
}
