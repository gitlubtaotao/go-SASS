package utils

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/beego/i18n"
	"strings"
)

//数据库连接
func DataBaseConnection(database string) {
	//需要配置时间为东八区,否则取出来的时间少8个小时
	_ = orm.RegisterDataBase("default", "mysql",
		"root:qweqwe123@tcp(127.0.0.1:3306)/"+database+"?charset=utf8&loc=Asia%2FShanghai")
	_ = orm.RunSyncdb("default", false, true)
	orm.SetMaxIdleConns("default", 30)
	orm.SetMaxOpenConns("default", 30)
}

//初始化App
func InitApp() {
	initLocales()
	InitMap()
}

var LangTypes []*LangType // Languages are supported.
// langType represents a language type.
type LangType struct {
	Lang, Name string
}

//初始化locales
func initLocales() {
	langs := strings.Split(beego.AppConfig.String("lang::types"), "|")
	names := strings.Split(beego.AppConfig.String("lang::names"), "|")
	LangTypes = make([]*LangType, 0, len(langs))
	for i, v := range langs {
		LangTypes = append(LangTypes, &LangType{
			Lang: v,
			Name: names[i],
		})
	}
	for _, lang := range langs {
		logs.Trace("Loading language: " + lang)
		//moreFile :=  "conf/"+"locale_model_"+lang+".ini"
		if err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini","conf/"+"locale_model_"+lang+".ini"); err != nil {
			logs.Error("Fail to set message file: " + err.Error())
			return
		}
	}
}
