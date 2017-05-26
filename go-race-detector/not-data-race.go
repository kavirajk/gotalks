package main

import "sync"

func main() {
	mutex := sync.Mutex

	// START1 OMIT
	// goroutine1
	mutex.Lock()
	counter = 4
	mutex.UnLock()
	// END1 OMIT

	// START2 OMIT
	// goroutine2
	mutex.Lock()
	counter = 5
	mutex.Unlock()
	// END2 OMIT
}
