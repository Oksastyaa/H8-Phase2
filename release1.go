package main

import (
	"fmt"
	"time"
)

// PrintNumbers  print all numbers
func PrintNumbers() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
	}
}

// PrintLetters  print all letters
func PrintLetters() {
	for ch := 'a'; ch <= 'j'; ch++ {
		fmt.Println(string(ch))
		time.Sleep(100 * time.Millisecond)
	}
}
func main() {
	go PrintNumbers()
	go PrintLetters()

	time.Sleep(2 * time.Second)
}
