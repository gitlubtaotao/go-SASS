package routers

import (
	"quickstart/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	
	//beego.InsertFilter("/article/*",beego.BeforeRouter, filerFunc)
	beego.InsertFilter("/login_out",beego.BeforeRouter,filerFunc)
	beego.InsertFilter("/article_type/*",beego.BeforeRouter,filerFunc)
	//beego.InsertFilter("/dist/*",beego.BeforeStatic, FilterNoCache)
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/login_out", &controllers.LoginController{}, "get:LoginOut")
	beego.Router("/company",&controllers.CompaniesController{})
	beego.Router("/company/new",&controllers.CompaniesController{},"get:New")
	beego.Router("/company/:id",&controllers.CompaniesController{},"delete:Delete;get:GetOne;put:Put")
	beego.Router("/user/index",&controllers.UserController{},"get:Get")
	beego.Router("/user",&controllers.UserController{},"get:GetAll;post:Post")
	beego.Router("/user/new",&controllers.UserController{},"get:New")
	beego.Router("/user/edit/:id",&controllers.UserController{},"get:Edit")
	beego.Router("/user/:id",&controllers.UserController{},"delete:Delete;get:GetOne;put:Put")
	beego.Router("/department",&controllers.DepartmentController{},"get:GetAll;post:Post")
}
//过滤器的使用
var filerFunc = func(ctx *context.Context) {
	userName := ctx.Input.Session("userName")
	if userName == nil {
		ctx.Redirect(302,"/login")
	}
}

//var FilterNoCache = func(ctx *context.Context) {
//	ctx.Output.Header("Cache-Control", "no-cache, no-store, must-revalidate")
//	ctx.Output.Header("Pragma", "no-cache")
//	ctx.Output.Header("Expires","0")
//}

