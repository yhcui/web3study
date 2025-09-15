package main

import (
	"fmt"
	"slices"
)

func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	result := merge1(intervals)
	fmt.Println(result)
}

func merge2(intervals [][]int) (ans [][]int) {
	slices.SortFunc(intervals, func(p, q []int) int { return p[0] - q[0] }) // 按照左端点从小到大排序
	for _, p := range intervals {
		m := len(ans)
		if m > 0 && p[0] <= ans[m-1][1] { // 可以合并
			ans[m-1][1] = max(ans[m-1][1], p[1]) // 更新右端点最大值
		} else { // 不相交，无法合并
			ans = append(ans, p) // 新的合并区间
		}
	}
	return
}

func merge1(intervals [][]int) (ans [][]int) {
	slices.SortFunc(intervals, func(e, e2 []int) int {
		return e[0] - e2[0]
	})

	for _, p := range intervals {
		m := len(ans)
		if m > 0 && p[0] <= ans[m-1][1] {
			ans[m-1][1] = max(p[1], ans[m-1][1])
		} else {
			ans = append(ans, p)
		}
	}
	return
}
