package main 

import "fmt"

func add(a int, arg ...int) int {
	var sum int = a 
	for i := 0; i< len(arg); i++ {
		sum += arg[i]
	}

	return sum 
}

// ...string 可变参数，可以给参数，也可以不给，可以给多个参数
func  concat(a string, arg ...string)(result string){

	result = a 
	for i := 0; i < len(arg); i++ {
		result += arg[i]
	}

	return  
}

func main(){
	sum := add(10, 3, 3, 3,3)
	fmt.Println(sum)

	res := concat("hello", " ", "world")
	fmt.Println(res)
}