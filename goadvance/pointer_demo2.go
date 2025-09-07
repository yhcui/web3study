package main

import "fmt"

func main() {
	nums := []int{2, 3, 4, 5}
	allAdd(&nums)
	fmt.Println(nums)
}
func allAdd(num *[]int) {
	for i := range *num {
		(*num)[i] = (*num)[i] * 2
	}
}
