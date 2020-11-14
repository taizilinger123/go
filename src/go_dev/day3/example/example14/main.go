package main  

import "fmt"

type  op_func func(int, int) int  

func add(a, b int) int {
	return a + b 
}

func sub(a, b int) int{
	return  a - b 
} 

//op是参数，参数类型是函数op_func跟a,b是参数类型是int一样
func operator(op  op_func, a, b int) int {
	 return  op(a, b)
}

func  main(){
	var a, b  int 
	add(a, b)
	//c := add 
	//sum := operator(c, 100, 200)
	var c op_func 
	c = add
	fmt.Println(add)
	fmt.Println(c)
	sum := operator(c, 100, 200)
	fmt.Println(sum)
}

//注意，map,slice,chan,指针，interface默认以引用的方式传递
