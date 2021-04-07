package main

import (
	"fmt"
)

func calc(taskChan chan int, resChan chan int, exitChan chan bool) {
	for v := range taskChan {
		flag := true
		//判断是不是质数，除了1和它自身外,不能被其他自然数整除的数叫做质数
		for i := 2; i < v; i++ {
			if v%i == 0 {
				flag = false
				break
			}
		}

		if flag {
			resChan <- v
		}
	}

	fmt.Println("exit")
	exitChan <- true
}

func main() {
	intChan := make(chan int, 1000)
	resultChan := make(chan int, 1000)
	exitChan := make(chan bool, 8)

	go func() {
		for i := 0; i < 10000; i++ {
			intChan <- i
		}

		close(intChan)
	}()

	for i := 0; i < 8; i++ {
		go calc(intChan, resultChan, exitChan)
	}

	//等待所有计算的goroutine全部退出,下面是一个携程go func(){}()方式
	go func() {
		for i := 0; i < 8; i++ {
			//<-exitChan //只是取出来不关注它的值
			a := <-exitChan
			fmt.Println(a)
		}
		close(resultChan)
	}()

	for v := range resultChan {
		fmt.Println(v)
	}
}
