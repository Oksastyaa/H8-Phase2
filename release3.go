package main

import (
	"fmt"
	"sync"
	"time"
)

func produce(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 10; i++ {
		ch <- i
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
}

func consume(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Println(num)
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(2)

	go produce(ch, &wg)
	go consume(ch, &wg)
	wg.Wait()
	fmt.Println("All Goroutines finished")
}
