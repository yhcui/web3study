package main

import "fmt"

func main() {
	var num int = 10
	add(&num)
	fmt.Println(num)
}
func add(num *int) {
	*num += 10
}
