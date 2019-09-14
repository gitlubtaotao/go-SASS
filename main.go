package main

import (
	"github.com/astaxie/beego/orm"
	_ "quickstart/routers"
	
	_ "quickstart/models"
	
	"github.com/astaxie/beego/logs"
	
	"github.com/astaxie/beego"
	_ "quickstart/utils"
)


func main() {
	orm.Debug = true
	log := logs.NewLogger(10000)
	_ = logs.SetLogger(logs.AdapterConsole, `{"level":7,"color":true}`)
	log.EnableFuncCallDepth(true)
	log.Async()
	beego.Run()
}




func init()  {
	//链接数据库
	dataConnection()
}

func dataConnection()  {
	_ = orm.RegisterDataBase("default", "mysql",
		"root:qweqwe123@tcp(127.0.0.1:3306)/go_quick_start?charset=utf8")
	_ = orm.RunSyncdb("default", false, true)
}
