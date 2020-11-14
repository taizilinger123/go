package main

import(
	"fmt"
)

// var a int8 = 100
// var b uint8 = 200
// var c float32
// var d float64

func swap(a *int, b *int){
	tmp := *a
	*a = *b
	*b = tmp
	return 
}

func swap1(a int, b int)(int, int){
	return b, a
}

func test(){
	var a = 100
	fmt.Println(a)
    var b int
	for i := 0;i < 100; i++{
		b = i * 2
		fmt.Println(b)
	}

	if a > 0 {
		var c int = 100
		fmt.Println(c)
	}
}

func test2(){
	var a int8 = 100
	var b int16 = int16(a)
	fmt.Printf("a=%d b=%d\n", a, b)
}

func main(){
	first := 100
	second := 200
	//swap(&first, &second)
	//first, second = swap1(first, second)
	first, second = second, first
	fmt.Println("first=", first)
	fmt.Println("second=", second)
	test2()
}

/*
G:\project>go  build  go_dev/day2/example9/main
G:\project>main.exe
first= 100
second= 200
*/