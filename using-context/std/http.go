package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"context"
)

// MAINSTART OMIT

func main() {
	client := http.Client{}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	// ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(1*time.Millisecond))
	// ctx, cancel := context.WithCancel(context.Background())

	req, _ := http.NewRequest("GET", "http://google.com", nil)
	req = req.WithContext(ctx)

	_, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("done")
	cancel() // it is good practice to call its cancelation function in any case
}

// MAINEND OMIT
