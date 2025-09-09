package main

import (
	"fmt"
	"time"
)

func main() {

	cc := make(chan int)
	go func(cc chan<- int) {
		for i := 1; i < 11; i++ {
			cc <- i
			fmt.Printf("发送%d\n", i)
		}
	}(cc)

	go func(cc <-chan int) {
		for a := range cc {
			fmt.Println("接收到:", a)
		}
	}(cc)

	fmt.Println("end")
	timeout := time.After(2 * time.Second)
	for {
		select {
		case v, ok := <-cc:
			if !ok {
				fmt.Println("通道关闭")
				return
			}
			fmt.Printf("主线程接受到%d\n", v)
		case <-timeout:
			fmt.Println("timeout")
			return
		default:
			fmt.Println("没有数据")
			time.Sleep(1 * time.Second)
		}

	}
}
