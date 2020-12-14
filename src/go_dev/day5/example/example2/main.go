package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score float32
	next  *Student
}

func trans(p *Student) {
	for p != nil {
		fmt.Println(*p)
		p = p.next
	}
}

func main() {
	var head Student
	head.Name = "hua"
	head.Age = 18
	head.Score = 100

	var stu1 Student
	stu1.Name = "stu1"
	stu1.Age = 23
	stu1.Score = 23
	// stu1.next = nil

	head.next = &stu1
	//var p *Student = &head
	//trans(&head)

	var stu2 Student
	stu2.Name = "stu2"
	stu2.Age = 23
	stu2.Score = 23

	stu1.next = &stu2
	trans(&head)

	var stu3 Student
	stu3.Name = "stu3"
	stu3.Age = 23
	stu3.Score = 23

	stu2.next = &stu3
	trans(&head)
}
