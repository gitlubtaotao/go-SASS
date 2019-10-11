package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
	"quickstart/models"
	"time"
)

//LoginController 登录页面
type LoginController struct {
	BaseController
}

//Get 登录页面
func (c *LoginController) Get() {
	name := c.Ctx.GetCookie("Name")
	if name != "" {
		c.Data["remember"] = "checked"
		c.Data["userName"] = name
	}
	c.Data["JsName"] = "login_in"
	c.SetSession("redirectUrl",c.GetString("url"))
	c.TplName = "login/index.html"
}

//Post 提交数据
func (c *LoginController) Post() {
	type Login struct {
		UserName string
		Password string
		Remember bool
	}
	v := Login{}
	_ = json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	logs.Info(v)
	Account := v.UserName
	Password := v.Password
	
	//验证账号和密码是否为空
	if Account == "" || Password == "" {
		c.Data["json"] = map[string]interface{}{"status": false,
			"message": "账号或者密码不能为空"}
		c.ServeJSON()
		c.StopRun()
	}
	//验证登录的账号是手机或者邮箱
	o := orm.NewOrm()
	user := models.User{}
	user.Email = Account
	//验证是否邮箱登录
	err := o.Read(&user, "Email")
	if err != nil {
		//	邮箱登录错误，进行手机进行验证
		user.Phone = Account
		err = o.Read(&user, "Phone")
		logs.Info(err)
		if err != nil {
			c.Data["json"] = map[string]interface{}{"status": false,
				"message": "账号或者密码错误"}
			c.ServeJSON()
			c.StopRun()
		}
	}
	//验证密码是否正确
	status := c.validatePassword(Password, user.EncodePassword)
	if !status {
		c.Data["json"] = map[string]interface{}{"status": false,
			"message": "账号或者密码错误"}
		c.ServeJSON()
		c.StopRun()
	}
	remember := v.Remember
	if remember {
		c.Ctx.SetCookie("Name", Account, time.Second*3600)
	}
	c.SetSession("currentName", user)
	var url interface{}
	if c.GetSession("redirectUrl") != "" {
		url = c.GetSession("redirectUrl")
	} else {
		url = "/"
	}
	logs.Info(url)
	c.Data["json"] = map[string]interface{}{"status": true,
		"url": url}
	c.ServeJSON()
}

//LoginOut 退出登录
func (c *LoginController) LoginOut() {
	c.DelSession("userName")
	c.DelSession("currentUser")
	c.pageLogin()
}

//验证输入的密码是否正确
func (c *LoginController) validatePassword(password string, encodePassword string) (status bool) {
	err := bcrypt.CompareHashAndPassword([]byte(encodePassword), []byte(password))
	logs.Info("sdsdsdsssssss ")
	logs.Info(err)
	if err != nil {
		return false
	} else {
		return true
	}
}
