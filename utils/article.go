package utils
import "github.com/astaxie/beego"

func init()  {
	_= beego.AddFuncMap("showDefaultType",showDefaultType)
	_ = beego.AddFuncMap("isSelect", isSelect)
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
