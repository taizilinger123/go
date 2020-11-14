package main

import (
	"fmt"
	"go_dev/day1/package_example/calc"
)

func main() {
	sum := calc.Add(100, 300)
	sub := calc.Sub(100, 300)

	fmt.Println("sum=", sum)
	fmt.Println("sub=", sub)
}

/*
G:\go>go  build  go_dev/day1/package_example/main

G:\go>dir
 驱动器 G 中的卷是 娱乐F
 卷的序列号是 0001-8A0C

 G:\go 的目录

2020/06/23  23:16    <DIR>          .
2020/06/23  23:16    <DIR>          ..
2020/06/21  23:40    <DIR>          .vscode
2020/06/23  22:05         2,076,160 example1.exe
2020/06/23  22:38         2,078,720 goroute.exe
2020/06/23  23:16         2,077,696 main.exe
2020/06/23  23:15    <DIR>          pkg
2020/06/23  22:53    <DIR>          src
               3 个文件      6,232,576 字节
               5 个目录 20,028,391,424 可用字节

G:\go>main.exe
sum= 400
sub= -200
*/
/*
把环境变量设置为GOPATH G:\project
G:\PROJECT>go build  -o  bin/calC_test.exe  go_dev/day1/package_example/main
G:\PROJECT>cd  bin
G:\project\bin>calC_test.exe
sum= 400
sub= -200
*/
