package main

import (
	"errors"
	"fmt"
	"time"
)

func initConfig() (err error) {
	return errors.New("init config failed")
}

func test() {
	/*
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()*/

	err := initConfig()
	if err != nil {
		panic(err)
	}

	b := 0
	a := 100 / b
	fmt.Println(a)
	return
}

func main() {
	for {
		test()
		time.Sleep(time.Second)
	}

	var a []int
	a = append(a, 10, 20, 383)
	a = append(a, a...)
	fmt.Println(a)
}

// PS D:\project> go run "d:\project\src\go_dev\day4\example\example2\main.go"
// [10 20 383 10 20 383]
//new:用来分配内存，主要用来分配值类型，比如int、struct。返回的是指针
//make:用来分配内存，主要用来分配引用类型，比如chan、map、slice
//append:用来追加元素到数组、slice中
//panic和recover:用来做错误处理
//close:主要用来关闭channel
//len:用来求长度，比如string、array、slice、map、channel
