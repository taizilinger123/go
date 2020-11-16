package main

import (
	"fmt"
)

func factor(n int) int {
	if n == 1 {
		return 1
	}
	return factor(n-1) * n
}

func fab(n int) int {
	if n <= 1 {
		return 1
	}

	return fab(n-1) + fab(n-2)
}

func main() {
	// fmt.Println(factor(5))
	for i := 0; i < 10; i++ {
		fmt.Println(fab(i))
	}
}

//计算递归
//斐波那契
// 1
// 1
// 2
// 3
// 5
// 8
// 13
// 21
// 34
// 55
