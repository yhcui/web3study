package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	vv := twoSum(nums, target)
	fmt.Println(vv)
}
func twoSum(nums []int, target int) []int {
	for i, v := range nums {
		for j := 1; j < len(nums); j++ {
			if v+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
