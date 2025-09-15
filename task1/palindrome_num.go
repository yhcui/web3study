package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 121321
	b := huiwen(n)
	fmt.Println(b)
}
func huiwen(num int) bool {
	if num < 0 {
		return false
	}
	bb := strconv.Itoa(num)
	l := len(bb)
	//i := 0
	//j := l
	//for i <= j {
	//	fmt.Println("方式1，第", i+1, "次循环")
	//}
	i := 0
	j := l - 1
	for ; i <= j; i, j = i+1, j-1 {
		if bb[i] == bb[j] {
			continue
		} else {
			return false
		}
	}
	return true
}
