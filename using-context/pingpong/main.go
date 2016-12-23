package main

import (
	"fmt"
	"time"
)

// START OMIT
type Ball struct{ hits int }

func main() {
	table := make(chan *Ball)
	go Player("ping", table)
	go Player("pong", table)
	table <- new(Ball)
	time.Sleep(1 * time.Second)
	<-table
}

func Player(name string, table chan *Ball) {
	for {
		ball := <-table
		ball.hits++
		fmt.Printf("%s: %d\n", name, ball.hits)
		time.Sleep(200 * time.Millisecond)
		table <- ball
	}
}

// END OMIT
