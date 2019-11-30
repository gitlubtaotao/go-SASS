package utils

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"math/rand"
	"time"
)

func Initialize() {
	_ = beego.AddFuncMap("ShowAssets", ShowAssets)
	err := beego.AddFuncMap("ShowSaleUser", ShowSaleUser)
	err = beego.AddFuncMap("RandNumber",RandNumber)
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
	if controllerName == "CustomerController"{
		return true
	}else{
		return false
	}
}


func RandNumber() int  {
	return rand.Intn(10)
}


//较长的时间格式输出
func LongTime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
//简短的时间格式输出
func ShortTime(t time.Time) string  {
	return t.Format("2006-01-02")
}

