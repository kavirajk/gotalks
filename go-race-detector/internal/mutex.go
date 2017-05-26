package main

import "sync"

var count = 0

var mu sync.Mutex // HL

// START OMIT
func increment() {
	mu.Lock() // HL
	if count == 0 {
		count = count + 1
	}
	mu.Unlock() // HL
}
func main() {
	go increment() //g1
	go increment() //g2
}

// END OMIT
