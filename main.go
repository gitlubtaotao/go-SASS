package main

import (
	"github.com/astaxie/beego/orm"
	_ "quickstart/routers"
	
	_ "quickstart/models"
	
	"github.com/astaxie/beego/logs"
	
	"github.com/astaxie/beego"
)

func main() {
	orm.Debug = true
	_ = beego.AddFuncMap("ShowPerPage", showPerPage)
	_= beego.AddFuncMap("showNextPage",showNextPage)
	_= beego.AddFuncMap("showDefaultType",showDefaultType)
	_ = beego.AddFuncMap("isSelect", isSelect)
	log := logs.NewLogger(10000)
	_ = log.SetLogger("console", "")
	log.EnableFuncCallDepth(true)
	log.Async()
	beego.Run()
}

// ShowPerPage 显示上一页
func showPerPage(data int) int {
	//pageTemp, _ := strconv.Atoi(data)
	pageIndex := data - 1
	if pageIndex <= 0{
		return 1
	}else {
		return pageIndex
	}
}
//showNextPage 显示上一页
func showNextPage(data int) int  {
	pageIndex := data + 1
	if pageIndex <= 0{
		return 1
	}else {
		return pageIndex
	}
}
//显示默认的类型
func showDefaultType(value string) bool  {
	if value == ""{
		return true
	}else{
		return false
	}
}

func isSelect(params string,value string) bool  {
	if params == value{
		return true
	}else{
		return false
	}
}
