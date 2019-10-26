package models

import "github.com/astaxie/beego/orm"

type CustomerSlice map[string]string
//ModelCount 统计数据的总条数
func ModelCount(tableName string) int64 {
	o := orm.NewOrm()
	cnt, _ := o.QueryTable(tableName).Count() // SELECT COUNT(*) FROM USER
	return cnt
}


