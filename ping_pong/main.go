package main

import (
	"fmt"
	"os"
	"net/http"
)

var counter int

func pingPongHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "pong %d\n", counter)
	counter++
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/pingpong", pingPongHandler)

	fmt.Printf("Server started in port %s\n", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}
