package controllers

type HomeController struct {
	 BaseController
}

//Get 首页
func (this *HomeController) Get()  {
	this.namespace = "home"
	this.setTpl("home/index.html")
	this.Data["Namespace"] = "home"
	this.Data["PageTitle"] = "首页"
}
