package controllers

import "C"
import (
	"github.com/astaxie/beego"
	"github.com/lhtzbj12/sdrms/enums"
	"html/template"
	"quickstart/models"
	"strings"
)

//BaseController
//controllerName: 获取当前controller名称
//actionName: 获取当前action名称
//currentUser: 获取当前用户信息
type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
	currentUser    models.User
	namespace      string //命名空间
}

//Prepare before action
func (this *BaseController) Prepare() {
	//跨站请求伪造
	this.TplExt = "html"
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
	//this.controllerName, this.actionName = this.GetControllerAndAction()
	this.Data["ControllerName"] = this.controllerName
	this.adapterUserInfo()
	//登录页面可以不需要进行登录判断
	if this.controllerName != "LoginController" {
		if this.actionName != "Get" && this.actionName != "Post" {
			this.checkLogin()
		}
	}
}

//Finish after method
func (this *BaseController) Finish() {
}

//检查用户是否进行登录
func (this *BaseController) checkLogin() {
	//表示用户没有进行登录
	if this.currentUser.Id == 0 {
		urlStr := this.URLFor("LoginController.Get") + "?url="
		returnURL := this.Ctx.Request.URL.Path
		//如果ajax请求则返回相应的错码和跳转的地址
		if this.Ctx.Input.IsAjax() {
			//由于是ajax请求，因此地址是header里的Referer
			returnURL = this.Ctx.Input.Refer()
			this.jsonResult(enums.JRCode302, "请登录", urlStr+returnURL)
		}
		this.Redirect(urlStr+returnURL, 302)
		this.StopRun()
	}
}

func (this *BaseController) jsonResult(code enums.JsonResultCode, msg interface{}, obj interface{}) {
	r := &models.JsonResult{Code: code, Msg: msg, Obj: obj}
	this.Data["json"] = r
	this.ServeJSON()
	this.StopRun()
}

//获取用户当前登录用户的信息
func (this *BaseController) adapterUserInfo() {
	user := this.GetSession("currentName")
	if user != nil {
		this.currentUser = user.(models.User)
		this.Data["currentUser"] = user
	}
}

//跳转到登录页面
func (this *BaseController) pageLogin() {
	url := this.URLFor("LoginController.Get")
	this.Redirect(url, 302)
	this.StopRun()
}

//重定向方法
func (this *BaseController) redirectCustomer(url string) {
	this.Redirect(url, 302)
	this.StopRun()
}

//模版渲染
//可以传入多个模版

func (this *BaseController) setTpl(template ...string) {
	var tplName string
	//默认的layouts
	layout := "layouts/application.html"
	switch len(template) {
	case 1:
		tplName = template[0]
	case 2:
		tplName = template[0]
		layout = template[1]
	default:
		ctrlName := strings.ToLower(this.controllerName[0 : len(this.controllerName)-10])
		actionName := strings.ToLower(this.actionName)
		tplName = ctrlName + "/" + actionName + ".html"
	}
	//设置默认的layout
	this.Layout = layout
	//设置模版名称
	this.TplName = tplName
}




