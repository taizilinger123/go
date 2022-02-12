package main

import (
	"bufio"
	"fmt"
	"os"
)

func count(str string) (worldCount, spaceCount, numberCount, otherCount int) {
	t := []rune(str)
	for _, v := range t {
		switch {
		case v >= 'a' && v <= 'z':
			fallthrough
		case v >= 'A' && v <= 'Z':
			worldCount++
		case v == ' ':
			spaceCount++
		case v >= '0' && v <= '9':
			numberCount++
		default:
			otherCount++

		}
	}
	return
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	result, _, err := reader.ReadLine()
	if err != nil {
		fmt.Println("read from console err:", err)
		return
	}
	wc, sc, nc, oc := count(string(result))
	fmt.Printf("world count:%d\n space count:%d\n number count:%d\n others count:%d\n", wc, sc, nc, oc)
}

// PS D:\project> go run "d:\project\src\go_dev\day4\work\example4\main.go"
// abc上海自来水来自海上cba
// world count:6
//  space count:0
//  number count:0
//  others count:9
