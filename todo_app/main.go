package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main(){
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	fmt.Println("Server is running on : ", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}