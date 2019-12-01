package controllers
import (
	"bytes"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/beego/i18n"
	"github.com/tealeg/xlsx"
	"html/template"
	"net/http"
	"quickstart/models"
	"reflect"
	"strconv"
	"strings"
	"time"
	
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
	i18n.Locale
}

var (
	AppVer string
	IsPro  bool
)

//Prepare before action
func (this *BaseController) Prepare() {
	//跨站请求伪造
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
	this.controllerName, this.actionName = this.GetControllerAndAction()
	this.Data["ControllerName"] = this.controllerName
	this.Data["JsName"] = ""
	this.Data["AppVer"] = AppVer
	this.Data["IsPro"] = IsPro
	this.adapterUserInfo()
	//登录页面可以不需要进行登录判断
	if this.controllerName != "LoginController" {
		if this.actionName != "Get" && this.actionName != "Post" {
			this.checkLogin()
		}
	}
	if this.setLangVer() {
		i := strings.Index(this.Ctx.Request.RequestURI, "?")
		this.Redirect(this.Ctx.Request.RequestURI[:i], 302)
		return
	}
}

//Finish after method
func (this *BaseController) Finish() {
}









//进行附件的下载
func (this *BaseController) DownLoad(data []interface{}, cols []models.CustomerSlice) {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error
	file = xlsx.NewFile()
	sheet, err = file.AddSheet("Sheet1")
	if err != nil {
		logs.Info(err.Error())
	}
	row = sheet.AddRow()
	for _, iv := range cols {
		cell = row.AddCell()
		cell.Value = iv["value"].(string)
	}
	cell = row.AddCell()
	for _, v := range data {
		row = sheet.AddRow()
		for _, iv := range cols {
			value := v.(map[string]interface{})
			name := iv["key"].(string)
			cell = row.AddCell()
			if reflect.TypeOf(value[name]) != nil {
				if reflect.TypeOf(value[name]).Kind() == reflect.Ptr {
					cell.Value = models.Struct2Map(value[name])["Name"].(string)
				} else if reflect.TypeOf(value[name]).Kind() == reflect.String {
					if value[name] != "" {
						cell.Value = value[name].(string)
					}
				} else if reflect.TypeOf(value[name]).Kind() == reflect.Int {
					cell.Value = strconv.Itoa(value[name].(int))
				} else if reflect.TypeOf(value[name]).Kind() == reflect.Int64 {
					cell.Value = strconv.FormatInt(int64(value[name].(int64)), 10)
				} else if reflect.TypeOf(value[name]).Kind() == reflect.Bool{
					cell.Value = strconv.FormatBool(value[name].(bool))
				}else{
				}
			}
		}
	}
	this.Ctx.ResponseWriter.Header().Add("Content-Disposition", "attachment")
	this.Ctx.ResponseWriter.Header().Add("Content-Type",
		"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	
	var buffer bytes.Buffer
	if err = file.Write(&buffer); err != nil {
		logs.Info(err)
	}
	r := bytes.NewReader(buffer.Bytes())
	http.ServeContent(this.Ctx.ResponseWriter, this.Ctx.Request, "", time.Now(), r)
	if err != nil {
		fmt.Printf(err.Error())
	}
}






