package main

import(
	"fmt"
)

func modify1(a *int){
	*a = 10
}

func modify(a int){
	a = 10
	return
}

func  main(){
	a := 5
	b := make(chan int, 1)

	fmt.Println("a=", a)
	fmt.Println("b=", b)

	modify(a)
	fmt.Println("a=",a)

	
	modify1(&a)
	fmt.Println("a=",a)
}