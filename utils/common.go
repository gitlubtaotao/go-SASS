package utils

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/beego/i18n"
	"math/rand"
	"time"
)

func InitMap() {
	_ = beego.AddFuncMap("ShowAssets", ShowAssets)
	err := beego.AddFuncMap("ShowSaleUser", ShowSaleUser)
	err = beego.AddFuncMap("RandNumber", RandNumber)
	_ = beego.AddFuncMap("i18n", i18n.Tr)
	_ = beego.AddFuncMap("LocaleS", LocaleS)
	logs.Error(err)
	
}

//是否显示对应的静态资源
func ShowAssets(namespace string) bool {
	if namespace != "" {
		return true
	} else {
		return false
	}
}

func ShowSaleUser(controllerName string) bool {
	if controllerName == "CustomerController" {
		return true
	} else {
		return false
	}
}

func RandNumber() int {
	return rand.Intn(10)
}

//较长的时间格式输出
func LongTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
func LocaleS(args ...string) string {
	logs.Info(args)
	stringA := make([]string, 1, len(args)+1)
	if len(stringA) > 0 {
		stringA = append(stringA, args...)
	}
	var output string
	for _, value := range stringA {
		output += value + " "
		
	}
	return output
}
