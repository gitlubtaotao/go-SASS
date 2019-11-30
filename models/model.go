package models

import (
	"github.com/astaxie/beego"
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
	t := reflect.TypeOf(models)
	typeOfT := s.Type()
	if t.Kind() == reflect.Ptr || t.Kind() == reflect.Struct {
		for i := 0; i < s.NumField(); i++ {
			rel := s.Field(i)
			if rel.Kind().String() != "slice" {
				fields = append(fields, typeOfT.Field(i).Name)
			}
		}
	}
	return fields
}

//获取struct 字段对应的类型
func StructFieldType(model interface{}, field string) (TypeName interface{}) {
	t := reflect.TypeOf(model)
	if t.Kind() == reflect.Struct || t.Kind() == reflect.Ptr {
		for i := 0; i < t.NumField(); i++ {
			if t.Field(i).Name == field {
				return t.Field(i).Type
			}
		}
	}
	return TypeName
}

//获取对象所有的方法
func structMethods(model interface{}) (methods []string) {
	t := reflect.TypeOf(model)
	if t.Kind() == reflect.Struct || t.Kind() == reflect.Ptr {
		for i := 0; i < t.NumMethod(); i++ {
			methods = append(methods, t.Method(i).Name)
		}
	}
	return methods
}

//数组元素是否存在某元素
//ArrayExistItem(2,[]int{1,23})
func ArrayExistItem(obj interface{}, target interface{}) bool {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true
		}
	}
	return false
}

//struct to map
func Struct2Map(obj interface{}) map[string]interface{} {
	var fields []string
	var data = make(map[string]interface{})
	s := reflect.ValueOf(obj).Elem()
	fields =  StrutFields(obj)
	for _, fname := range fields {
		data[fname] = s.FieldByName(fname).Interface()
	}
	return  data
}

//获取数据库的表名
func TableName(name string) string {
	prefix := beego.AppConfig.String("db_dt_prefix")
	return prefix + name
}
