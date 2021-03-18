package main

import "fmt"

func main() {
	var a interface{}
	var b int

	a = b
	c := a.(int)
	fmt.Printf("%d %T\n", c, c)
}