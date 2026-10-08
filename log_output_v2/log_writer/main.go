package main

import (
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

	logFile, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("File writing error : %v \n", err)
	}

	// os.WriteFile(filePath, []byte(randomStr.String()), 0644)
	logFile.WriteString(randomStr.String())
}
