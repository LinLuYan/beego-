package commonFunc

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"crypto/md5"
	"encoding/hex"
)

type Common struct {
	beego.Controller
}

//设置跨域
func (c *Common) CrossDomai(){
	c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
}
//请求返回
func (c *Common) ResHandler(code, msg, data string){
    type Response struct{
        Code string `json:"code"`
        Msg string `json:"msg"`
        Data string `json:"data"`
	}
	result := Response{}
	result.Code = code
	result.Msg = msg
	result.Data = data
	buffer, err := json.Marshal(&result)
	if err != nil{
		beego.Error(err)
	} else {
		c.Ctx.Output.Body(buffer)
	}
}
//登录状态
func (c *Common) IsSessionOn() bool{
	result := true
	v := c.GetSession("user")
	if v == nil{
		result = false
	}
	return result
}
//md5
func (c *Common)GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
