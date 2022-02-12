package main

import (
	"fmt"
    "math/rand"
)

func main() {

	// var a int = 10

	// switch a {
	// case 0, 1, 2, 3, 4, 5:
	// 	fmt.Println("a is equal 0")
	// 	fmt.Println("yes ")
	// case 10:
	// 	fmt.Println("a is equal 10")
	// 	fallthrough
	// default:
	// 	fmt.Println("a is equal default")
	// }
	var  n  int  
	n = rand.Intn(100)
    for {
		var input int 
		fmt.Scanf("%d\n", &input)
		flag := false 
		switch {
		case  input == n:
			fmt.Println("you  are  right")
			flag = true
		case  input > n:
			fmt.Println("bigger")
		case  input < n:
			fmt.Println("less")
		}
		if flag {
			break
		}
	}

}
