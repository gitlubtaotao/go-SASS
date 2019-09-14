package utils

import "github.com/astaxie/beego"

func init() {
	_ = beego.AddFuncMap("MenuActive", MenuActive)
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
