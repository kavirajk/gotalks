package main

import (
	"fmt"
	"runtime/debug"
	"time"
)

// START OMIT
type Ball struct{ hits int }

var done = make(chan struct{}) // HL
// END OMIT

// MAINSTART OMIT
func main() {
	table := make(chan *Ball)
	go Player("ping", table)
	go Player("pong", table)
	table <- new(Ball)
	time.Sleep(1 * time.Second)
	close(done) // HL
	<-table
	debug.SetTraceback("all")
	panic("stack trace")
}

// MAINEND OMIT

// PLAYERSTART OMIT
func Player(name string, table chan *Ball) {
	for {
		select { // HL
		case <-done: // HL
			return // HL
		default:
			ball := <-table
			ball.hits++
			fmt.Printf("%s: %d\n", name, ball.hits)
			time.Sleep(200 * time.Millisecond)
			table <- ball
		}
	}
}

// PLAYEREND OMIT
