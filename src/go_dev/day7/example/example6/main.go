package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("len of args:%d\n", len(os.Args))
	for i, v := range os.Args {
		fmt.Printf("args[%d]=%s\n", i, v)
	}
}

// PS D:\project> go run "d:\project\src\go_dev\day7\example\example6\main.go"
// len of args:1
// args[0]=C:\Users\shenglei\AppData\Local\Temp\go-build707277822\b001\exe\main.exe
// PS D:\project> go run "d:\project\src\go_dev\day7\example\example6\main.go"  -c  test.txt
// len of args:3
// args[0]=C:\Users\shenglei\AppData\Local\Temp\go-build698349198\b001\exe\main.exe
// args[1]=-c
// args[2]=test.txt
