package main

import "fmt"

type student struct {
	name string
}

func main() {
	var stuChan chan *student
	stuChan = make(chan *student, 10)
	stu := student{name: "stu01"}

	stuChan <- &stu
	fmt.Println(stuChan)
}
