package routers

import (
	"study/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//restful路由
	ns := beego.NewNamespace("v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)

    beego.AddNamespace(ns)

	ns0 := beego.NewNamespace("/build",
		beego.NSInclude(
			&controllers.ViewsController{},
		),
	)
	beego.AddNamespace(ns0)
}
