package controllers

import (
	"github.com/astaxie/beego"
)

type ViewsController struct {
	beego.Controller
}

// @Title 入口
// @Description 远程重启
// @Success 200 {object} string
// @Failure 403 body is empty
// @router /reStart [get]
func (c *ViewsController) ReStart() {
	c.TplName = "index.html"
	c.Render()
}
