package main

import (
	"fmt"
	"time"
)

type Car struct {
	Name string
	Age  int
}

func (c *Car) Set(name string, age int){
	c.Name = name 
	c.Age = age
}

type Car2 struct {
	Name string
}

type Train struct {
	//car   Car
	Car
	Car2
	createTime time.Time
	//count  int
	int
}

func (t *Train) Set(age int) {
	t.int = age
}

// func (t Train) Set(age int) {
// 	t.int = age
// }

// 与上面的函数一样的
// func Set(t Train, age int){
// 	t.int = age
// }
//Set(train, age)

func main() {
	var train Train
	train.int = 300
	train.Car.Set("huas", 100)
	//（&train）.Set(1000)与train.Set(1000)一样
	//Set(train, age)
	// train.Car.Name = "test"
	train.Car.Name = "test" //因为Car和Car2都有Name所以必须指定用哪个Car
	fmt.Println(train)
}
