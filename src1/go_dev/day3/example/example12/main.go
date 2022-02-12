package main  

import  "fmt"

func Print(n int){
   
	for  i := 1; i < n+1; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("A")
		}
		fmt.Println()
	}
}

func main(){
   Print(6)
}

/*
输出
A
AA
AAA
AAAA
AAAAA
AAAAAA
*/