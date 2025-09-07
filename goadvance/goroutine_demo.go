package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go printA(&wg)
	go printB(&wg)
	wg.Wait()
	fmt.Println("end")
}
func printA(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i < 10; i++ {
		if i%2 == 1 {
			fmt.Printf("%d ", i)
		}
	}
}

func printB(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i < 10; i = i + 2 {
		if i%2 == 0 {
			fmt.Printf("%d ", i)
		}
	}
}
