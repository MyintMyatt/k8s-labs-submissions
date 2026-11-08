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

	randomStr, err := uuid.NewRandom()
	if err != nil {
		log.Fatalf("UUID generation error: %v", err)
	}

	logFile, err := os.OpenFile(
		filePath,
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		0664,
	)
	if err != nil {
		log.Fatalf("Failed to open or create file: %v", err)
	}
	defer logFile.Close()

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C

		timestamp := time.Now().Format(time.RFC3339)
		line := timestamp + " " + randomStr.String() + "\n"

		bytesWritten, err := logFile.WriteString(line)
		if err != nil {
			log.Fatalf("Failed to write to file: %v", err)
		}

		log.Printf("Successfully appended %d bytes to the file!", bytesWritten)
	}
}
