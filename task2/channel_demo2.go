package main

import (
	"fmt"
	"sync"
)

func main() {

	ch := make(chan int, 30)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go pp(ch, &wg)
	go cc(ch, &wg)

	wg.Wait()
}
func pp(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 100; i++ {
		ch <- i
		fmt.Println("生产", i)
	}
	close(ch)
}
func cc(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch {
		fmt.Println("消费", v)
	}
}
