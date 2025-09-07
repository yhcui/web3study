package main

import "fmt"

// plusOne 对表示大整数的数组加一
func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		// 如果等于 9，重置为 0
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			// 不为 9，直接加一返回
			digits[i]++
			return digits
		}
	}
	// 所有位都是 9，创建新数组，长度加 1，首位设为 1
	result := make([]int, len(digits)+1)
	result[0] = 1
	return result
}

func main() {
	// 测试用例
	testCases := [][]int{
		{1, 2, 3},    // 预期: [1,2,4]
		{9, 9},       // 预期: [1,0,0]
		{4, 3, 2, 1}, // 预期: [4,3,2,2]
		{9},          // 预期: [1,0]
	}
	for _, digits := range testCases {
		fmt.Printf("Input: %v, Output: %v\n", digits, plusOne(digits))
	}
}
