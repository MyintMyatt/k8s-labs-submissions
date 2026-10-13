package main

import (
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

const DEFAULT_PATH = "../log.txt"

func main() {
	log.Println("Log writer is running......")

	filePath := os.Getenv("LOG_DIR")

	if filePath == "" {
		filePath = DEFAULT_PATH
	}

	var eror error
	var logFile *os.File
	logFile, eror = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0664)
	if eror != nil {
		log.Fatalf("Failed to open or create file: %s \n", eror)
	}

	defer logFile.Close()

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {

		randomStr, err := uuid.NewRandom()
		if err != nil {
			log.Fatalf("File writing error : %v \n", err)
		}

		bytesWritten, err := logFile.WriteString(randomStr.String() + "\n")
		if err != nil {
			log.Fatalf("Failed to write to file: %s", err)
		}

		log.Printf("Successfully appended %d bytes to the file!\n", bytesWritten)

		<-ticker.C
	}
}
