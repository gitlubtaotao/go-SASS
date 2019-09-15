package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

//User person
type User struct {
	Id        int    `orm:"pk;auto"`
	Name      string `orm:"unique"`
	Pwd       string
	Articles []*Article `orm:"rel(m2m)"`
	Company *Company `orm:"rel(fk)"`
}
//Article 文章标题
type Article struct {
	Id        int    `orm:"pk;auto"`
	ArtiName  string `orm:"default(null)"`
	Atime     time.Time  `orm:"auto_now;type(datetime)"`
	Account   int `orm:"default(0);null"`
	Acontent  string
	Aimg      string
	ArticleType *ArticleType `orm:"rel(fk)"`
	Users []*User `orm:"reverse(many)"`

}
//ArticleType 文章类型
type ArticleType struct {
	Id int
	TypeName string `orm:"size(64);"`
	Articles []*Article `orm:"reverse(many)"`
}

func init()  {
	orm.RegisterModel(new(Article), new(User),new(ArticleType))
}

//ModelCount 统计数据的总条数
func ModelCount(tableName string) int64 {
	o := orm.NewOrm()
	cnt, _ := o.QueryTable(tableName).Count() // SELECT COUNT(*) FROM USER
	return cnt
}
