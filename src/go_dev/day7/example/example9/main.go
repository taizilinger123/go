package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	UserName string
	NickName string
	Age      int
	Birthday string
	Sex      string
	Email    string
	Phone    string
}

func testStruct() (ret string, err error) {
	user1 := &User{
		UserName: "user1",
		NickName: "上课看似",
		Age:      18,
		Birthday: "2008/8/8",
		Sex:      "男",
		Email:    "mahuateng@qq.com",
		Phone:    "110",
	}

	data, err := json.Marshal(user1)
	if err != nil {
		fmt.Errorf("json.marshal failed, err:", err)
		return
	}
	ret = string(data)
	return
}

func test() {
	data, err := testStruct()
	if err != nil {
		fmt.Println("test struct failed, ", err)
		return
	}

	var user1 User
	err = json.Unmarshal([]byte(data), &user1)
	if err != nil {
		fmt.Println("Unmarshal failed, ", err)
		return
	}
	fmt.Println(user1)
}

func main() {
	test()
}

// PS D:\project> go run "d:\project\src\go_dev\day7\example\example9\main.go"
// {user1 上课看似 18 2008/8/8 男 mahuateng@qq.com 110}
