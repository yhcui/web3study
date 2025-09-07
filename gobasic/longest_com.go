package main

import "fmt"

func main() {
	// 测试用例
	testCases := [][]string{
		{"flower", "flow", "flight"},                   // 预期: "fl"
		{"dog", "racecar", "car"},                      // 预期: ""
		{"interspecies", "interstellar", "interstate"}, // 预期: "inter"
		{},         // 预期: ""
		{"prefix"}, // 预期: "prefix"
	}
	for _, strs := range testCases {
		fmt.Printf("Input: %v, Output: %q\n", strs, longestCommonPrefix(strs))
	}
}

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}
	an := strs[0]
	for i := 1; i < len(strs); i++ {
		for j := 0; j < len(strs[i]) && j < len(an); j++ {
			if an[j] == strs[i][j] {
				continue
			} else {
				an = an[:j]
				break
			}
			if an == "" {
				return an
			}
		}
	}
	return an
}
