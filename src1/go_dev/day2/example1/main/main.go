package main 

import(
	"fmt"
)

func list(n int){
	for i := 0; i <= n;i++ {
		fmt.Printf("%d+%d+%d\n", i , n -i, n)
	}
}

func main(){
	list(10)
}

/*
ctrl+~->TERMINAL
G:\project>go  build  -o  bin/list.exe  go_dev/day2/example1/main
G:\project>dir  bin
G:\project>cd  bin
G:\project\bin>list.exe
*/