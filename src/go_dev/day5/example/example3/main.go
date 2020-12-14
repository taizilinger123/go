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

func insertTail(p *Student){
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

func main() {
	var head Student
	head.Name = "hua"
	head.Age = 18
	head.Score = 100

    insertTail(&head)
	trans(&head)
}
