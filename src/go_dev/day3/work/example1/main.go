package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++{
        if n%i == 0 {
           return false
		}
	}
	return true
}

// func modify(n *int){
// 	fmt.Println("in  modify function,", n)
// 	*n = 100
// }

func main(){
	var n int 
	var m int 
	fmt.Scanf("%d%d", &n, &m)
	
	for i := n; i < m; i++ {
		if isPrime(i) == true {
			fmt.Printf("%d\n", i)
			continue
		}
	}

	// fmt.Println("out modify function,", &n)
	// modify(&n)
	// fmt.Println(n)
}