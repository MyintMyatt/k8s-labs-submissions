package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	
	"todo_app/handlers"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	fmt.Println("Server is running on : ", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.LandingPageHandler(w, *r)
	})

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
