package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"quickstart/controllers"
)

func init() {
	beego.InsertFilter("/login_out",beego.BeforeRouter,filerFunc)
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/login_out", &controllers.LoginController{}, "get:LoginOut")
	
	beego.Router("/company",&controllers.CompaniesController{},"get:GetAll;post:Post")
	beego.Router("/company/index",&controllers.CompaniesController{},"get:Get")
	beego.Router("/company/new",&controllers.CompaniesController{},"get:New")
	beego.Router("/company/edit/:id",&controllers.CompaniesController{},"get:Edit")
	beego.Router("/company/:id",&controllers.CompaniesController{},"delete:Delete;get:GetOne;put:Put")
	
	beego.Router("/user/index",&controllers.UserController{},"get:Get")
	beego.Router("/user",&controllers.UserController{},"get:GetAll;post:Post")
	beego.Router("/user/new",&controllers.UserController{},"get:New")
	beego.Router("/user/edit/:id",&controllers.UserController{},"get:Edit")
	beego.Router("/user/:id",&controllers.UserController{},"delete:Delete;get:GetOne;put:Put")
	
	beego.Router("/department",&controllers.DepartmentController{},"get:GetAll;post:Post")
	beego.Router("/department/index",&controllers.DepartmentController{},"get:Get")
	beego.Router("/department/new",&controllers.DepartmentController{},"get:New")
	beego.Router("/department/edit/:id",&controllers.DepartmentController{},"get:Edit")
	beego.Router("/department/:id",&controllers.DepartmentController{},"put:Put;get:GetOne;delete:Delete")
	
	beego.Router("/order",&controllers.OrderController{},"get:GetAll;post:Post")
	beego.Router("/order/index",&controllers.OrderController{},"get:Get")
	beego.Router("/order/new",&controllers.OrderController{},"get:New")
	beego.Router("/order/edit/:id",&controllers.OrderController{},"get:Edit")
	beego.Router("/order/:id",&controllers.OrderController{},"put:Put;get:GetOne;delete:Delete")
	
	beego.Router("/customer",&controllers.CustomerController{},"get:GetAll;post:Post")
	beego.Router("/customer/index",&controllers.CustomerController{},"get:Get")
	beego.Router("/customer/new",&controllers.CustomerController{},"get:New")
	beego.Router("/customer/edit/:id",&controllers.CustomerController{},"get:Edit")
	beego.Router("/customer/:id",&controllers.CustomerController{},"put:Put;get:GetOne;delete:Delete")
	beego.Router("/customer/get_status",&controllers.CustomerController{},"get:GetStatus")
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

