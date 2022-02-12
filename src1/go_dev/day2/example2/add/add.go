package add  

import(
    _  "go_dev/day2/example2/test"
)

// var  Name  string = "hello world"
// var  Age  int = 10 
var Name string = "xxxxx"
var Age int = 100

func init(){
	 Name = "hello world"
	 Age = 10
}


//执行顺序:import->全局变量->init函数->main函数