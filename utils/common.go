package utils

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"math/rand"
)

func init() {
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


