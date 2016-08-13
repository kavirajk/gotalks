package main

import "fmt"

func Reverse(s string) string {
	chars := []rune(s)
	n := len(chars)
	for i := 0; i < n/2; i++ {
		chars[i], chars[n-1-i] = chars[n-1-i], chars[i]
	}
	return string(chars)
}

func main() {
	fmt.Println(Reverse("🙂!puteeM gnaloG erolgnaB ot emocleW"))
}
