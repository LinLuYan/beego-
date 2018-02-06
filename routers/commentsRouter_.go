package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["study/controllers:UserController"] = append(beego.GlobalControllerRouter["study/controllers:UserController"],
		beego.ControllerComments{
			Method: "IsLogin",
			Router: `/isLogin`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study/controllers:UserController"] = append(beego.GlobalControllerRouter["study/controllers:UserController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study/controllers:UserController"] = append(beego.GlobalControllerRouter["study/controllers:UserController"],
		beego.ControllerComments{
			Method: "Start",
			Router: `/start`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study/controllers:UserController"] = append(beego.GlobalControllerRouter["study/controllers:UserController"],
		beego.ControllerComments{
			Method: "Test",
			Router: `/test`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["study/controllers:ViewsController"] = append(beego.GlobalControllerRouter["study/controllers:ViewsController"],
		beego.ControllerComments{
			Method: "ReStart",
			Router: `/reStart`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
