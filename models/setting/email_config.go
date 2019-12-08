package setting

import "github.com/astaxie/beego/orm"

type EmailConfig struct {
	Id       int64  `orm:"pk;auto"`
	Address  string `orm:"size(128)"`
	Port     int8
	Domain   string `orm:"size(128)"`
	UserName string `orm:"size(128)"`
	Password string `orm:"size(128)"`
}

func init()  {
	orm.RegisterModel(new(EmailConfig))
}


