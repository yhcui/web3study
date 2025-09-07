package main

import (
	"errors"
	"fmt"
)

func main() {
	nums := []int{7, 7, 8, 8, 9, 9, 0, 1, 23, 0, 1}
	r, err := quchong(nums)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("result:", r)
	}
}
func quchong(nums []int) (int, error) {
	m := make(map[int]int)
	for _, num := range nums {
		fmt.Println("---", num, ":", m[num])
		m[num] += 1
	}
	for k, v := range m {
		fmt.Println(k, "---", v)
		if v == 1 {

			return k, nil
		}
	}
	return 0, errors.New("not found")
}
