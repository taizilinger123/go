package main

import (
	"fmt"
)

type slice struct {
	ptr *[100]int
	len int
	cap int
}

func make1(s slice, cap int) slice {
	s.ptr = new([100]int)
	s.cap = cap
	s.len = 0
	return s
}

func modify(s slice) {
	s.ptr[1] = 1000
}

func testSlice2() {
	var s1 slice
	s1 = make1(s1, 10)

	s1.ptr[0] = 100
	modify(s1)
	fmt.Println(s1.ptr)
}

func testSlice() {
	var slice []int
	var arr [5]int = [...]int{1, 2, 3, 4, 5}

	// slice = arr[2:5]
	//slice = arr[:] //包含整个数组的长度[1 2 3 4 5]
	slice = arr[1:]              //包含下标是1开始的元素[2 3 4 5]
	slice = slice[:len(slice)-1] //[2 3 4]
	//slice = arr[:3] //包含从0开始到3结束的元素[1 2 3]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice = slice[0:1]      //对上面的slice=[3 4 5]再取值[3]包含左边，不包含右边
	fmt.Println(len(slice)) //长度
	fmt.Println(cap(slice)) //容量
}

func modify1(a []int) {
	a[1] = 1000
}

func testSlice3() {
	var b []int = []int{1, 2, 3, 4}
	modify1(b)
	fmt.Println(b)
}

func testSlice4() {
	var a = [10]int{1, 2, 3, 4}

	b := a[1:5]
	fmt.Printf("%p\n", b)
	fmt.Printf("%p\n", &a[1])
}
func main() {
	//testSlice()
	// testSlice2()
	//testSlice3()
	testSlice4()
}
