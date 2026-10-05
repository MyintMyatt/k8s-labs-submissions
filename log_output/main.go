package main

import (
	"fmt"
	"time"
	"net/http"

	"github.com/google/uuid"
)

func main()  {
	randomString :=  uuid.New().String()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(randomString))
	})
	http.ListenAndServe(":6969", nil)

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		fmt.Printf("%s: %s\n", time.Now().UTC().Format(time.RFC3339Nano), randomString);
		<- ticker.C
	}

}