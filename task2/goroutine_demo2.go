package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	inters := []func(){func() {
		time.Sleep(2 * time.Second)
		fmt.Println("task1 completed")
	}, func() {
		time.Sleep(2 * time.Second)
		fmt.Println("task2 completed")
	}, func() {
		time.Sleep(2 * time.Second)
		fmt.Println("task3 completed")
	}, func() {
		time.Sleep(2 * time.Second)
		fmt.Println("task4 completed")
	}, func() {
		time.Sleep(2 * time.Second)
		fmt.Println("task5 completed")
	}}
	times := schedue(inters)
	fmt.Println("\nTask Execution Times:")
	for i, d := range times {
		fmt.Printf("Task %d: %v\n", i, d)
	}
}

func schedue(inters []func()) map[int]time.Duration {
	if len(inters) == 0 {
		return make(map[int]time.Duration)
	}
	results := sync.Map{}
	var wg sync.WaitGroup
	wg.Add(len(inters))

	for i, inter := range inters {
		index := i
		go func() {
			defer wg.Done()
			start := time.Now()
			inter()
			duration := time.Since(start)
			results.Store(index, duration)
		}()

	}
	wg.Wait()

	resultMap := make(map[int]time.Duration)
	results.Range(func(key, value interface{}) bool {
		resultMap[key.(int)] = value.(time.Duration)
		return true
	})
	return resultMap
}
