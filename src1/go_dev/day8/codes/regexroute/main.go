package main

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func main() {
	//正则路由
	//URL中定义正则字符串方式进行匹配
	//匹配 /数字/ 格式的路由  => 并把匹配的值放入到:id参数
	beego.Get("/name/?:name(\\w+)/", func(ctx *context.Context) {
		// ctx.WriteString("匹配name")
		name := ctx.Input.Param(":name")
		ctx.WriteString(fmt.Sprintf("匹配name: %s", name))
	})
	beego.Get("/id/:id(\\d+)/", func(ctx *context.Context) {
		// ctx.WriteString("匹配id")
		id := ctx.Input.Param(":id")
		ctx.WriteString(fmt.Sprintf("匹配id: %s", id))
	})
	beego.Get("/any/:context/", func(ctx *context.Context) {
		// ctx.WriteString("匹配context")
		context := ctx.Input.Param(":context")
		ctx.WriteString(fmt.Sprintf("匹配context: %s", context))
	})
	beego.Get("/file/*", func(ctx *context.Context) {
		// ctx.WriteString("匹配file")
		splat := ctx.Input.Param(":splat")
		ctx.WriteString(fmt.Sprintf("匹配file: %s", splat))
	})
	beego.Get("/ext/*.*", func(ctx *context.Context) {
		// ctx.WriteString("匹配ext")
		path := ctx.Input.Param(":path")
		ext := ctx.Input.Param(":ext")
		ctx.WriteString(fmt.Sprintf("匹配ext: %s[%s]", path, ext))
	})
	beego.Run()
}

//postman里get方法：http://localhost:8080/ext/a/b/c.json 输出：匹配ext: a/b/c[json]
//http://localhost:8080/name/xxxx,  http://localhost:8080/name
