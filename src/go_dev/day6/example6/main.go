package main

func main() {

	var intLink Link
	for i := 0; i < 10; i++ {
		//intLink.InsertHead(i)
		intLink.Insert(i)
	}

	intLink.Trans()
}

//PS D:\project> go run "d:\project\src\go_dev\day6\example6"   注意这里编译多个go文件一定要这样子
