package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

func main()  {
	randomString :=  uuid.New().String()
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		fmt.Printf("%s: %s\n", time.Now().UTC().Format(time.RFC3339Nano), randomString);
		
		<- ticker.C
	}
}