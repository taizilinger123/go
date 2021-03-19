package main

import "fmt"

func Test(a interface{}) {
	b, ok := a.(int)
	if ok == false {
		fmt.Println("convert failed")
		return
	}
	b += 3
	fmt.Println(b)
}

func main() {

	var b int
	Test(b)
}
