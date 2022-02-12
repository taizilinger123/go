package main

import (
	"fmt"
)

type Cart1 struct {
	name string
	age  int
}

type Cart2 struct {
	name string
	age  int
}

type Train struct {
	Cart1
	Cart2
}

func main() {
	var t Train
	t.Cart1.name = "train"
	t.Cart1.age = 100

	// t.Cart.age = 200     //结构体可以不写Cart
	// t.Cart.name = "001"  //结构体可以不写Cart

	fmt.Println(t)
}
