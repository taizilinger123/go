package main

import (
	"fmt"
)

func main() {
	var i int
	fmt.Println(i)

	j := new(int) //new返回的是地址
	*j = 100
	fmt.Println(*j)
}
