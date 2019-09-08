package routers

import (
	"quickstart/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.InsertFilter("/article/*",beego.BeforeRouter, filerFunc)
	beego.InsertFilter("/login_out",beego.BeforeRouter,filerFunc)
	beego.InsertFilter("/article_type/*",beego.BeforeRouter,filerFunc)
	
	beego.Router("/", &controllers.MainController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/login_out", &controllers.LoginController{}, "get:LoginOut")
	beego.Router("/article", &controllers.ArticleController{})
	beego.Router("/article/add", &controllers.ArticleController{}, "get:Add")
	beego.Router("/article/show",&controllers.ArticleController{},"get:Show")
	beego.Router("/article/update",&controllers.ArticleController{},"get:Edit;post:Update")
	beego.Router("/article_type/create",&controllers.ArticleTypeController{},"get:New;post:Create")
	//beego.Router("/article_type",&controllers.ArticleTypeController{},"get:Index")
}
//过滤器的使用
var filerFunc = func(ctx *context.Context) {
	userName := ctx.Input.Session("userName")
	if userName == nil {
		ctx.Redirect(302,"/login")
	}
}
