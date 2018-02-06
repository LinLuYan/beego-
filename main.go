package main

import (
	_ "study/routers"
	"github.com/astaxie/beego"
	//"study/files"
)


func main() {
	//files
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 60*10
	beego.BConfig.WebConfig.StaticDir["/build"] = "build"
	beego.Run()
}

