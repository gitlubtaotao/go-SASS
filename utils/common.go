package utils

import (
	"github.com/astaxie/beego"
)

func init() {
	_ = beego.AddFuncMap("MenuActive", MenuActive)
	_ = beego.AddFuncMap("ShowAssets", ShowAssets)
}

//是否显示对应的静态资源
func ShowAssets(namespace string) bool  {
	assetSlice := []string {"company"}
	status := false
	for _,value := range assetSlice{
		if value == namespace{
			status = true
			break
		}
	}
	return status
}

//MenuActive 当前菜单高亮
func MenuActive(namespace string) string {
	var value string
	switch namespace {
	case "home":
		value= "active open"
	case "article":
		value="active open"
		
	}
	return value
}
