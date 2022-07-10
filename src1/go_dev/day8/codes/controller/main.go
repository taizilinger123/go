package main

import (
	"github.com/astaxie/beego"
)

//定义控制器HomeController
type HomeController struct {
	beego.Controller
}

func main() {
	beego.Run()
}
