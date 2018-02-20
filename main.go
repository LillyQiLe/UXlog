package main

import (
	_ "uxlog/routers"
	"github.com/astaxie/beego"
)

func main() {
	//beego.BConfig.WebConfig.StaticDir["/static"] = "static"
	//beego.SetStaticPath("/avatar", "/avatar")
	beego.Run()
}

