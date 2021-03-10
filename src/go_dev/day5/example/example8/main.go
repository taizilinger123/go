package main

import (
	"fmt"
)

type integer int 

func  (p integer) print(){
	fmt.Println("p is ", p)
}

func  (p *integer) set(b integer){
	*p = b
}


type Student struct {
	Name  string 
	Age   int 
	Score int 
	Sex   int 
}
//p是结构体Student,函数init里面有3个参数，后面可以接int返回类型,这里没写
//func (p Student) init(name string, age int, score int) int{}
func (p *Student) init(name string, age int, score int) {
	p.Name = name 
	p.Age = age 
	p.Score = score
	fmt.Println(p)
}

func (p Student) get() Student {
	return p
}

func main(){
	var stu Student
	// (&stu).init("stu", 10, 200)
	stu.init("stu", 10, 200)
	stu1 := stu.get()
	fmt.Println(stu1)

	var a integer
	a = 100
	a.print()

	// (&a).set(1000)
	a.set(1000)
	a.print()
}