package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(10)

	var counter int32

	for i := 0; i < 10; i++ {
		go func(cc *int32) {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				atomic.AddInt32(cc, 1)
			}
		}(&counter)
	}
	wg.Wait()

	fmt.Println("end", counter)
}
