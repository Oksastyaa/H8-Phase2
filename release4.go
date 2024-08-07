package main

import (
	"fmt"
	"sync"
	"time"
)

func Produce(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		ch <- i
		time.Sleep(100 * time.Millisecond)
	}
	close(ch)
}

func Consume(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Println(num)
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 5)
	wg.Add(2)
	go Produce(ch, &wg)
	go Consume(ch, &wg)
	wg.Wait()
}

/**
channel unbuffered func produce hanya dapat mengirim nilai ketika fungsi 'consume' siap menerimanya,
jika consume lambat , 'produce' akan diblokir sampai 'consume' membaca dari channel

channel buffered disini saya menggunakan contoh 5 buffer size , artinya func 'produce' akan mengirimkan 5 nilai tanpa menunggu
'consume' membacanya
**/
