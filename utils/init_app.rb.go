package utils

import (
	"github.com/astaxie/beego/orm"
	"quickstart/controllers"
)

//数据库连接
func DataBaseConnection() {
	//需要配置时间为东八区,否则取出来的时间少8个小时
	_ = orm.RegisterDataBase("default", "mysql",
		"root:qweqwe123@tcp(127.0.0.1:3306)/go_quick_start?charset=utf8&loc=Asia%2FShanghai")
	_ = orm.RunSyncdb("default", false, true)
	orm.SetMaxIdleConns("default", 30)
	orm.SetMaxOpenConns("default", 30)
}

//初始化App
func InitApp() {
	controllers.InitLocales()
	Initialize()
}
