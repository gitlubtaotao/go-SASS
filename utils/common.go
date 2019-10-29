package utils

import (
	"github.com/astaxie/beego"
)

func init() {
	_ = beego.AddFuncMap("ShowAssets", ShowAssets)
}

//是否显示对应的静态资源
func ShowAssets(namespace string) bool {
	if namespace != "" {
		return true
	} else {
		return false
	}
}


