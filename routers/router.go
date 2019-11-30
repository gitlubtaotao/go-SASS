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
	
	beego.Router("/company",&controllers.CompanyController{},"get:GetAll;post:Post")
	beego.Router("/company/:id",&controllers.CompanyController{},"delete:Delete;get:GetOne;put:Put")
	beego.AutoRouter(&controllers.CompanyController{})
	
	
	beego.Router("/user",&controllers.UserController{},"get:GetAll;post:Post")
	beego.Router("/user/:id",&controllers.UserController{},"delete:Delete;get:GetOne;put:Put")
	beego.AutoRouter(&controllers.UserController{})
	
	beego.Router("/department",&controllers.DepartmentController{},"get:GetAll;post:Post")
	beego.Router("/department/:id",&controllers.DepartmentController{},"put:Put;get:GetOne;delete:Delete")
	beego.AutoRouter(&controllers.DepartmentController{})
	
	beego.Router("/order",&controllers.OrderController{},"get:GetAll;post:Post")
	beego.Router("/order/:id",&controllers.OrderController{},"put:Put;get:GetOne;delete:Delete")
	beego.AutoRouter(&controllers.OrderController{})
	
	beego.Router("/customer",&controllers.CustomerController{},"get:GetAll;post:Post")
	beego.Router("/customer/:id",&controllers.CustomerController{},"put:Put;get:GetOne;delete:Delete")
	beego.AutoRouter(&controllers.CustomerController{})
	
	beego.Router("/supplier",&controllers.SupplierController{},"get:GetAll;post:Post")
	beego.Router("/supplier/:id",&controllers.SupplierController{},"put:Put;get:GetOne;delete:Delete")
	beego.AutoRouter(&controllers.SupplierController{})
	
	beego.Router("/contact",&controllers.ContactController{},"get:GetAll;post:Post")
	beego.Router("/contact/:id",&controllers.ContactController{},"put:Put;get:GetOne;delete:Delete")
	beego.AutoRouter(&controllers.ContactController{})
	
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

