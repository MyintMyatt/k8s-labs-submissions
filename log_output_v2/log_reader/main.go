package main

import (
	"log"
	"net/http"
	"os"
)

const DEFAULT_PATH = "../log.txt"

func main() {
	log.Println("Log reader is running .................")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		content, err := readFile()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		w.Write(content)
	})
	http.ListenAndServe(":6969", nil)

}

func readFile() ([]byte, error) {
	filePath := os.Getenv("LOG_DIR")

	if filePath == "" {
		filePath = DEFAULT_PATH
	}
	return os.ReadFile(filePath)
}
