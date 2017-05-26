package main

var count = 0

func increment() {
	if count == 0 {
		count = count + 1
	}
}

func main() {
	go increment()
	go increment()
}
