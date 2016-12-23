package main

import (
	"context"
	"fmt"
	"runtime/debug"
	"time"
)

// START OMIT
type Ball struct{ hits int }

// MAINSTART OMIT
func main() {
	table := make(chan *Ball)
	ctx, cancel := context.WithCancel(context.Background()) // HL
	go Player(ctx, "ping", table)                           // HL
	go Player(ctx, "pong", table)                           // HL
	table <- new(Ball)
	time.Sleep(1 * time.Second)
	cancel() // HL
	<-table
	debug.SetTraceback("all")
	panic("stack trace")
}

// MAINEND OMIT

// PLAYERSTART OMIT
func Player(ctx context.Context, name string, table chan *Ball) {
	for {
		select { // HL
		case <-ctx.Done(): // HL
			return // HL
			// PLAYERTMPEND OMIT
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
