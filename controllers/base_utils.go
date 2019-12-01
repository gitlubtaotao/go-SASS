package controllers

import (
	"github.com/astaxie/beego/logs"
	"github.com/beego/i18n"
	"quickstart/enums"
	"quickstart/models"
	"quickstart/utils"
	"strings"
)

//检查是否登录
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

//获取登录用户的信息
func (this *BaseController) adapterUserInfo() {
	user := this.GetSession("currentName")
	if user != nil {
		this.currentUser = user.(models.User)
		this.Data["currentUser"] = user
	}
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

//检查用户是否进行登录
func (this *BaseController) jsonResult(code enums.JsonResultCode, msg interface{}, obj interface{}) {
	r := &models.JsonResult{Code: code, Msg: msg, Obj: obj}
	this.Data["json"] = r
	this.ServeJSON()
	this.StopRun()
}

//设置locale
func (this *BaseController) setLangVer() bool  {
	isNeedRedir := false
	hasCookie := false
	// 1. Check URL arguments.
	lang := this.Input().Get("lang")
	// 2. Get language information from cookies.
	if len(lang) == 0 {
		lang = this.Ctx.GetCookie("lang")
		hasCookie = true
	} else {
		isNeedRedir = true
	}
	
	if !i18n.IsExist(lang) {
		lang = ""
		isNeedRedir = false
		hasCookie = false
	}
	
	// 3. Get language information from 'Accept-Language'.
	if len(lang) == 0 {
		al := this.Ctx.Request.Header.Get("Accept-Language")
		if len(al) > 4 {
			al = al[:5] // Only compare first 5 letters.
			if i18n.IsExist(al) {
				lang = al
			}
		}
	}
	// 4. Default language is English.
	if len(lang) == 0 {
		lang = "zh-CN"
		isNeedRedir = false
	}
	curLang := utils.LangType{
		Lang: lang,
	}
	// Save language information in cookies.
	if !hasCookie {
		this.Ctx.SetCookie("lang", curLang.Lang, 1<<31-1, "/")
	}
	restLangs := make([]*utils.LangType, 0, len(utils.LangTypes)-1)
	for _, v := range utils.LangTypes {
		if lang != v.Lang {
			restLangs = append(restLangs, v)
		} else {
			curLang.Name = v.Name
		}
	}
	this.Lang = lang
	this.Data["Lang"] = curLang.Lang
	this.Data["CurLang"] = curLang.Name
	this.Data["RestLangs"] = restLangs
	logs.Info(curLang,isNeedRedir,hasCookie)
	return isNeedRedir
}
