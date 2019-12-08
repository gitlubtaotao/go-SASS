package controllers

import (
	"github.com/beego/i18n"
)

type HomeController struct {
	BaseController
}

//Get 首页
func (this *HomeController) Get() {
	this.namespace = "home"
	this.setTpl("home/index.html")
	this.Data["Namespace"] = "home"
	this.Data["PageTitle"] = i18n.Tr(this.Lang, "home")
	
}
