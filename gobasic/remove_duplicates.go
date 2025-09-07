package main

import "fmt"

func main() {
	nums := []int{1, 2, 2, 3, 4, 5, 5, 5, 6}
	result := removeDuplicates(nums)
	fmt.Println(result)
}
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	i, j := 0, 0
	for j < len(nums) {
		if nums[i] == nums[j] {
			j++
		} else {
			i++
			j++
		}
	}
	return i
}
