package main

import "fmt"

type Carer interface {
	GetName() string
	Run()
	DiDi()
}

func main() {
	var a interface{}
	var b int
	var c float32
	//空接口可以保存任何类型的变量，反过来不行b = a
	a = b
	a = c
	fmt.Printf("type of a %T\n", a)
}
