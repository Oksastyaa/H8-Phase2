package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumbers(wg *sync.WaitGroup, start chan bool) {
	defer wg.Done()
	<-start
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters(wg *sync.WaitGroup, start chan bool) {
	defer wg.Done()
	for ch := 'a'; ch <= 'j'; ch++ {
		fmt.Println(string(ch))
		time.Sleep(100 * time.Millisecond)
	}
	close(start)
}

func main() {
	var wg sync.WaitGroup
	start := make(chan bool)
	wg.Add(2)
	go printNumbers(&wg, start)
	go printLetters(&wg, start)
	wg.Wait()
}
