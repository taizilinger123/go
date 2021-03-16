package main

import "fmt"

type Carer interface {
	GetName() string
	Run()
	DiDi()
}

type BMW struct {
	Name string
}

func (p *BMW) GetName() string {
	return p.Name
}

func (p *BMW) Run() {
	fmt.Println("%s is running\n", p.Name)
}

func (p *BMW) DiDi() {
	fmt.Println("%s is didi\n", p.Name)
}

func main() {
	var car Carer
	fmt.Println(car)

	var bmw BMW
	car = &bmw
	car.Run()
}
