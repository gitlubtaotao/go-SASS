package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"quickstart/models/oa"
)

// RegisterController 注册功能
type RegisterController struct {
	beego.Controller
}

//Get 注册页面
func (c *RegisterController) Get() {
	c.TplName = "registers/index.html"
}

//Post 提交数据
func (c *RegisterController) Post() {
	userName := c.GetString("userName")
	pwd := c.GetString("password")
	if userName == "" || pwd == "" {
		fmt.Println("数据较验错误")
		c.Data["message"] = "用户名或者密码不能为空"
		c.TplName = "register/index.html"
		
	}
	o := orm.NewOrm()
	user := oa.User{Name: userName, Pwd: pwd}
	_, err := o.Insert(&user)
	if err != nil {
		fmt.Println("注册错误：", user)
		c.Data["message"] = err
		c.TplName = "register/index.html"
	}
	c.Redirect("/login", 302)
}
