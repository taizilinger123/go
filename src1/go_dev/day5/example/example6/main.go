package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func main() {
	var stu Student = Student{
		Name:  "stu01",
		Age:   10,
		Score: 80,
	}

	data, err := json.Marshal(stu)
	if err != nil {
		fmt.Println("json encode stu failed, err:", err)
		return
	}
	fmt.Println(string(data))
}

// PS D:\project> go run "d:\project\src\go_dev\day5\example\example6\main.go"
// {"name":"stu01","age":10,"score":80}
