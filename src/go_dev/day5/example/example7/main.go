package main

import (
	"fmt"
	"time"
)

type Cart struct {
	name string
	age  int
}

type Train struct {
	Cart
	int
	start time.Time
}

func main() {
	var t Train
	t.name = "train"
	t.age = 100

	// t.Cart.age = 200     //结构体可以不写Cart
	// t.Cart.name = "001"  //结构体可以不写Cart
	t.int = 200

	fmt.Println(t)
}
