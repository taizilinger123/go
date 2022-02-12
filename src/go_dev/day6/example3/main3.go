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

type BYD struct {
	Name string
}

func (p *BYD) GetName() string {
	return p.Name
}

func (p *BYD) Run() {
	fmt.Printf("%s is running\n", p.Name)
}

func (p *BYD) DiDi() {
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

	byd := &BYD{
		Name: "BYD",
	}
	car = byd
	car.Run()
}

//提交到https://github.com/taizilinger123/go
//用windows的GUI
//shenglei@LAPTOP-86RD7CIL MINGW64 /d/project (master)
//git add *
//git   commit   -m  "20210318"
//git   push   https://github.com/taizilinger123/go.git
