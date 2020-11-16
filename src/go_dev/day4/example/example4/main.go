package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello")
	time.Sleep(time.Second)
	main()
}

//递归举例
