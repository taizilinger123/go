package main

import (
	"fmt"
)

var sum int

func test_goroute(a int, b int) {
	sum := a + b
	fmt.Println(sum)
}

func test_goroute2(a int, b int) {
	fmt.Println(sum)
}

//gofmt  -w  test.go
func test_print(a int) {
	for {

	}
}
