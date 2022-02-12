package main  

import(
	 a  "go_dev/day2/example2/add"
	"fmt"
)

func main(){
	fmt.Println("Name=", a.Name)
	fmt.Println("age=", a.Age)
}

/*
G:\project>go  build  -o  bin/example2.exe  go_dev/day2/example2/main
G:\project>cd  bin
G:\project\bin>example2.exe
Name= hello world
age= 10
*/