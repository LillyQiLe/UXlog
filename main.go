package main

import (
	_ "uxlog/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("static/avatar/", "avatar")
	beego.Run()
}

