package utils

import "github.com/astaxie/beego"

func init() {
	_ = beego.AddFuncMap("ShowPerPage", showPerPage)
	_ = beego.AddFuncMap("showNextPage", showNextPage)
}

// ShowPerPage 显示上一页
func showPerPage(data int) int {
	//pageTemp, _ := strconv.Atoi(data)
	pageIndex := data - 1
	if pageIndex <= 0 {
		return 1
	} else {
		return pageIndex
	}
}

//showNextPage 显示上一页
func showNextPage(data int) int {
	pageIndex := data + 1
	if pageIndex <= 0 {
		return 1
	} else {
		return pageIndex
	}
}
