package main

// START OMIT
// shared variable
var count = 0

func increment() {
	if count == 0 { // HL
		count = count + 1 // HL
	}
}

func main() {
	go increment() //g1
	go increment() //g2
}

// END OMIT
