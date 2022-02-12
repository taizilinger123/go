package main

import (
	"fmt"
)

func process(str string) bool {
	t := []rune(str)
	length := len(t)
	for i, _ := range t {
		if i == length/2 {
			break
		}
		last := length - i - 1
		//fmt.Printf("%c----%c, %d---%d, %v",str[i], str[last], i, last, v)*/
		//fmt.Printf("%v----%v----%d\n", i, v, len(string(v)))

		if t[i] != t[last] {
			return false
		}
	}

	return true
}

func main() {
	var str string
	fmt.Scanf("%sd", &str)
	if process(str) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}

// PS D:\project> go run "d:\project\src\go_dev\day4\work\example3\main.go"
// 上海自来水来自海上
// 0----19978----3
// 1----28023----3
// 2----33258----3
// 3----26469----3
// 4----27700----3
// 5----26469----3
// 6----33258----3
// 7----28023----3
// 8----19978----3
// yes
//abc上海自来水来自海上cba
//1上海自来水来自海上1
//没有被使用的变量用_代替
