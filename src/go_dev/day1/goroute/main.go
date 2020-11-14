package main

import (
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		go test_goroute(i)
	}

	time.Sleep(time.Second)
}
/*
G:\go>go  build  go_dev/day1/goroute
C:\Users\sige>go  env
值类型:int,float,bool,string,数组,struct
引用类型:指针，slice,map,chan
*/