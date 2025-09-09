package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	var mutex sync.Mutex
	var count int
	for i := 0; i < 10; i++ {
		go func(mutex *sync.Mutex, wg *sync.WaitGroup, cc *int) {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				mutex.Lock()
				*cc++
				mutex.Unlock()
			}
		}(&mutex, &wg, &count)
	}
	wg.Wait()
	fmt.Println("v:", count)
}
