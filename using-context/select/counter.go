package main

import (
	"fmt"
	"os"
	"time"
)

// START OMIT
func main() {
	abort := make(chan struct{})
	fmt.Println("Commencing countdown. Press return to abort")
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	tick := time.Tick(1 * time.Second)
	for countdown := 5; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select { // HL
		case <-tick: // HL
		case <-abort: // HL
			fmt.Println("Launch aborted.") // HL
			return                         // HL
		} // HL
	}
	launch()
}

// END OMIT
func launch() {
	fmt.Println("BAAAAAAAAAAAAAANG!!!!!")
}
