package controllers

import (
	"study/utils"
	"encoding/json"
	"github.com/astaxie/beego"
	"os/exec"
	"bytes"
)

type UserController struct {
	commonFunc.Common
}

type User struct{
	Account string `json:"account"`
	PassWord string `json:"passWord"`
}

//密码账号
var account = ""
var passWord = ""

// @Title Login
// @Description 登录
// @Success 200 {string} login success
// @Failure 403 body is empty
// @router /login [post]
func (c *UserController) Login() {
	//设置跨域
    c.CrossDomai()
    code := ""
    msg := ""
    data := ""
    //声明最后执行
	defer func(){
		c.ResHandler(code, msg, data)
	}()
    //判断登录状态
    if c.IsSessionOn(){
    	code = "0001"
    	msg = "您已登录"
    	return
	}
    //进行登录操作
    userInfo := User{}
    err := json.Unmarshal(c.Ctx.Input.RequestBody, &userInfo)
    if err != nil{
    	code = "0002"
    	msg = "参数解析失败"
    	beego.Error(err)
    	return
	}
    if userInfo.Account == c.GetMd5String(account) && userInfo.PassWord == c.GetMd5String(passWord){
		code = "0000"
		msg = "登录成功"
		c.SetSession("user",c.GetMd5String(account))
		return
	}else {
		code = "0003"
		msg = "账号密码错误"
		return
	}
}

// @Title start
// @Description 登录
// @Success 200 {string} login success
// @Failure 403 body is empty
// @router /start [post]
func (c *UserController) Start() {
	//设置跨域
	c.CrossDomai()
	code := ""
	msg := ""
	data := ""
	//声明最后执行
	defer func(){
		c.ResHandler(code, msg, data)
	}()

	//判断登录状态
	if !c.IsSessionOn(){
		code = "0001"
		msg = "未登录"
		return
	}

	cmd := exec.Command("/bin/bash", "-c", "supervisorctl restart ezonNew")
	//cmd := exec.Command("cmd")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Run()

	code = "0000"
	msg = out.String()
}

// @Title isLogin
// @Description 验证登录
// @Success 200 {string} login success
// @Failure 403 body is empty
// @router /isLogin [post]
func (c *UserController) IsLogin() {
	//设置跨域
	c.CrossDomai()
	code := ""
	msg := ""
	data := ""
	//声明最后执行
	defer func(){
		c.ResHandler(code, msg, data)
	}()

	//判断登录状态
	if c.IsSessionOn(){
		code = "0000"
		msg = "登录"
		return
	}else{
		code = "0001"
		msg = "未登录"
	}

}

// @Title Test
// @Description 登录
// @Success 200 {string} test success
// @Failure 403 body is empty
// @router /test [get]
func (c *UserController) Test() {
	c.Ctx.WriteString("22222222222")
}