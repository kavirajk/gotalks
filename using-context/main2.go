package main

import (
	"context"
	"fmt"
	"time"
)

type Ball struct {
	hits int
}

func main() {
	ball := new(Ball)
	table := make(chan *Ball)
	ctx, cancel := context.WithCancel(context.Background())
	go Player(ctx, "ping", table)
	go Player(ctx, "pong", table)
	table <- ball
	time.Sleep(2 * time.Second)
	cancel()
	<-table
	panic("stack trace")
}

func Player(ctx context.Context, name string, table chan *Ball) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			ball := <-table
			ball.hits++
			fmt.Printf("%s: %d\n", name, ball.hits)
			time.Sleep(200 * time.Millisecond)
			table <- ball
		}
	}
}
