package main

import "runtime/debug"

// START OMIT
func main() {
	ch := make(chan int)
	go func() { ch <- 0 }()
	<-ch

	debug.SetTraceback("all")
	panic("stack trace")
}

// END OMIT
