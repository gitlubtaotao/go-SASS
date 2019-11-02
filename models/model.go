package models

import (
	"github.com/astaxie/beego/orm"
	"reflect"
)

type CustomerSlice map[string]interface{}
type CustomerBoolSlice map[bool]string

//ModelCount 统计数据的总条数
func ModelCount(tableName string) int64 {
	o := orm.NewOrm()
	cnt, _ := o.QueryTable(tableName).Count() // SELECT COUNT(*) FROM USER
	return cnt
}


//获取struct所有的字段
func StrutFields(models interface{}) []string {
	var fields []string
	s := reflect.ValueOf(models).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		rel := s.Field(i)
		if rel.Kind().String() != "slice" {
			fields = append(fields, typeOfT.Field(i).Name)
		}
		//logs.Info(rel.Kind(),rel.Type(),rel)
		//fmt.Printf("%d: %s %s = %v\n", i, f.Type(), f.Interface())
	}
	return fields
}
