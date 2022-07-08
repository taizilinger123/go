package main

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func main() {

	//路由  url => controller
	//处理器函数 处理器： 函数 结构体4
	//beego/context/Context
	//绑定函数
	//以GET方式请求/通过绑定函数处理
	beego.Get("/", func(ctx *context.Context) {
		//用户数据的获取
		name := ctx.Input.Query("name")

		//给用户响应数据
		ctx.Output.Context.WriteString(fmt.Sprintf("你输入的名字是: %s", name))
	})
	beego.Post("/", func(ctx *context.Context) {
		name := ctx.Input.Query("name")
		ctx.Output.Context.WriteString(fmt.Sprintf("(POST)你输入的名字是: %s", name))
	})
	beego.Any("/any", func(ctx *context.Context) {
		name := ctx.Input.Query("name")
		ctx.Output.Context.WriteString(fmt.Sprintf("(%s)你输入的名字是: %s", ctx.Input.Method(), name))
	})
	//启动beego程序
	beego.Run()
}

/*
C:\Users\Administrator>　set GO111MODULE=on
C:\Users\Administrator>netstat -aon | findstr 8080
  TCP    0.0.0.0:8080           0.0.0.0:0              LISTENING       7824
  TCP    [::]:8080              [::]:0                 LISTENING       7824

C:\Users\Administrator>taskkill | findstr "7824"
错误: 无效语法。没有指定 /FI、/PID 或 /IM。
有关用法，请键入 "TASKKILL /?"。

C:\Users\Administrator>tasklist  | findstr "7824"
main.exe                      7824 Console                    1      7,752 K

C:\Users\Administrator>taskkill   /F   /IM  main.exe
成功: 已终止进程 "main.exe"，其 PID 为 7824。

C:\Users\Administrator>netstat -aon | findstr 8080

*/
