package main

import (
	"fmt"
	"sort"
)

func testIntSort() {
	var a = [...]int{1, 8, 30, 2, 348, 484}
	sort.Ints(a[:])

	fmt.Println(a)
}

func testStrings() {
	var a = [...]string{"abc", "cfg", "b", "A", "eeee"}
	sort.Strings(a[:])

	fmt.Println(a)
}

func testFloat() {
	var a = [...]float64{2.3, 0.8, 20.2, 392342.2, 0.6}
	sort.Float64s(a[:])

	fmt.Println(a)
}

func testIntSearch() {
	var a = [...]int{1, 8, 30, 2, 348, 484}
	sort.Ints(a[:])
	index := sort.SearchInts(a[:], 348)
	fmt.Println(index)
}

func main() {
	// testIntSort()
	// testStrings()
	//testFloat()
	testIntSearch()
}

//[1 2 8 30 348 484]
