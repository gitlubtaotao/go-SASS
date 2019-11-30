package main

import (
	"github.com/astaxie/beego/orm"
	"quickstart/utils"
	"quickstart/utils/redis"
	
	_ "quickstart/routers"
	
	_ "quickstart/models"
	
	"github.com/astaxie/beego/logs"
	
	"github.com/astaxie/beego"
)

func main() {
	orm.Debug = true
	log := logs.NewLogger(10000)
	_ = logs.SetLogger(logs.AdapterConsole, `{"level":7,"color":true}`)
	log.EnableFuncCallDepth(true)
	log.Async()
	beego.Run()
}

func init() {
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.SetStaticPath("/assets", "assets")
	beego.SetStaticPath("/dist", "dist")
	beego.SetStaticPath("/views", "views")
	
	//链接redis
	redis.RedisNewClient()
	//链接数据库
	utils.DataBaseConnection()
	utils.InitApp()
}
