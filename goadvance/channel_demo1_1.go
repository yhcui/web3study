package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go produce(ch, &wg)
	go consumer(ch, &wg)
	wg.Wait()
}

func produce(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i < 11; i++ {
		ch <- i
		fmt.Printf("生产%d\n", i)
	}
	close(ch)
}

func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch {
		fmt.Printf("消费%d\n", v)
	}
}
