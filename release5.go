package main

import (
	"fmt"
	"sync"
)

func sendNumbers(evenCh, oddCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			evenCh <- i
		} else {
			oddCh <- i
		}
	}
	close(evenCh)
	close(oddCh)
}

func main() {
	evenCh := make(chan int)
	oddCh := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go sendNumbers(evenCh, oddCh, &wg)

	done := make(chan bool)
	go func() {
		wg.Wait()
		done <- true
	}()

	for {
		select {
		case even, ok := <-evenCh:
			if ok {
				fmt.Printf("Received even number : %d\n", even)
			}
		case odd, ok := <-oddCh:
			if ok {
				fmt.Printf("Received odd number : %d\n", odd)

			}
		case <-done:
			fmt.Println("All numbers have been process")
			return
		}

	}

}
