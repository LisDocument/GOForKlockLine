package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

var u = "1111"


func (l *LoginController) Get()  {
	const user_name = "admin"
	const user_pass = "admin"
	l.Data["username"] = "li"
	l.Data["password"] = "bin"

	var username string
	l.Ctx.Input.Bind(&username,"username")
	var password string
	l.Ctx.Input.Bind(&password,"password")
	fmt.Println(username,password)
	if(username == user_name && password == user_pass){
		l.Data["json"] = map[string]interface{}{"name":"登录成功"}
	}else{
		l.Data["json"] = map[string]interface{}{"name":"密码错误"}
	}
	l.ServeJSON()
	return
}

func (l *LoginController) Post(){

}