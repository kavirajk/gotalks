package main

var count = 0

// START OMIT
func increment() {
	if count == 0 {
		count = count + 1
	}
}

func main() {
	increment() //g1 // HL
	increment() //g2 // HL
}

// END OMIT
