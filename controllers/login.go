package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"quickstart/models"
	"time"
)

//LoginController 登录页面
type LoginController struct {
	beego.Controller
}

//Get 登录页面
func (c *LoginController) Get() {
	name := c.Ctx.GetCookie("Name")
	fmt.Println(name)
	if name != ""{
		c.Data["remember"] = "checked"
		c.Data["userName"] = name
	}
	c.TplName = "login/index.html"
}

//Post 提交数据
func (c *LoginController) Post() {
	UserName := c.GetString("userName")
	Pwd := c.GetString("password")
	if UserName == "" || Pwd == "" {
		fmt.Println("输入数据有误")
		c.TplName = "login/index.html"
		//c.Redirect("/login",302)
		return
	}
	o := orm.NewOrm()
	user := models.User{}
	user.Pwd = Pwd
	user.Name = UserName
	err := o.Read(&user, "Name", "Pwd")
	if err != nil {
		fmt.Println("账号或者密码不正确", err)
		c.Data["message"] = "账号或者密码不能为空"
		c.TplName = "login/index.html"
		return
	}
	remember := c.GetString("remember")
	fmt.Println(remember)
	if remember == "true" {
		c.Ctx.SetCookie("Name", UserName, time.Second*3600)
	}else{
		c.Ctx.SetCookie("Name","",)
	}
	c.SetSession("userName",UserName)
	c.Redirect("/",302)
	
}

//LoginOut 退出登录
func (c *LoginController) LoginOut() {
	c.DelSession("userName")
	c.Redirect("/login",302)
}
