package main

import (
	"fmt"
	"math/rand"
)

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

//尾部插入，正序
func insertTail(p *Student) {
	var tail = p
	for i := 0; i < 10; i++ {
		stu := Student{
			Name:  fmt.Sprintf("stu%d", i),
			Age:   rand.Intn(100),
			Score: rand.Float32() * 100,
		}

		tail.next = &stu
		tail = &stu
	}
}

//头部插入,倒叙
func insertHead(p **Student) {
	//var tail = p
	for i := 0; i < 10; i++ {
		stu := Student{
			Name:  fmt.Sprintf("stu%d", i),
			Age:   rand.Intn(100),
			Score: rand.Float32() * 100,
		}
		stu.next = *p
		*p = &stu
		//tail.next = &stu
		//tail = &stu
	}
}

func delNode(p *Student) {
	var prev *Student = p
	for p != nil {
		if p.Name == "stu6" {
			prev.next = p.next
			break
		}
		prev = p
		p = p.next
	}

}

func addNode(p *Student, newNode *Student){
	for p != nil {
		if p.Name == "stu5" {
			newNode.next = p.next
			p.next = newNode
			break
		}
		p = p.next
	}
}

func main() {
	var head *Student = new(Student)
	head.Name = "hua"
	head.Age = 18
	head.Score = 100

	//insertTail(head)
	//trans(head)
	insertHead(&head)
	trans(head)
	delNode(head)
	trans(head)
	var newNode *Student = new(Student)
	newNode.Name = "stu1000"
	newNode.Age = 18
	newNode.Score = 100
	addNode(head, newNode)
	trans(head)
}
