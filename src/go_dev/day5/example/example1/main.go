package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	score float32
}

func main() {
	var stu Student
	stu.Age = 18
	stu.Name = "hua"
	stu.score = 80
	
	var stu1 *Student = &Student{
		Age: 20, 
		Name: "hua",
	}
	
	var stu3 = Student{
		Age: 20, 
		Name: "hua",
	}
	fmt.Println(stu)
	fmt.Println(stu1.Name)
	fmt.Println(stu3)
	fmt.Println("Name:%p\n", &stu.Name)
	fmt.Println("Age:%p\n", &stu.Age)
	fmt.Println("score:%p\n", &stu.score)
}
