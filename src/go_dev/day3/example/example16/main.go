package main  

import "fmt"

// func read(){
// 	lock.lock()

// 	defer lock.Unlock()
// 	if i == 5 {
// 		return 
// 	}

// 	if i == 10 {
// 		return 
// 	}

// 	func read1() {
// 		file := open(filename)
// 		defer file.Close()  
// 	}

// 	func read2() {
// 		mc.Lock()
// 		defer mc.Unlock()
// 	}

// 	func  read3(){
// 		conn, err := openConn()
// 		defer func() {
// 			if err != nil {
// 				conn.Close()
// 			}
// 	}()
// }

var (
	result = func(a1 int, b1 int) int {
		return a1 + b1 
	}
)

//匿名函数
func test(a, b int) int {
	result := func(a1 int, b1 int) int {
		return  a1 + b1 
	}(a, b)

	return result
}

func main(){
    fmt.Println(result(100, 200))

	var i int = 0
	defer fmt.Println(i)  //先进后出，存在栈里，输出10->second->0 
	defer fmt.Println("second")

	i = 10 
	fmt.Println(i)
}