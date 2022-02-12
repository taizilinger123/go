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
	fmt.Printf("%s is running\n", p.Name)
}

func (p *BMW) DiDi() {
	fmt.Printf("%s is didi\n", p.Name)
}

func main() {
	var car Carer
	fmt.Println(car)

	bmw := &BMW{
		Name: "BMW",
	}
	car = bmw
	car.Run()
}
