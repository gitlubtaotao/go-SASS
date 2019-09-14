package controllers

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"quickstart/models"
)

//ArticleTypeController 文章类型
type  ArticleTypeController struct {
		BaseController
}

//Index 显示
//func (c *ArticleTypeController) Index()  {
//	c.TplName="article_types/index.html"
//}

// New : 新增文章类型
func (this *ArticleTypeController) New()  {
	o := orm.NewOrm()
	var article_types []models.ArticleType
	_,err := o.QueryTable("ArticleType").All(&article_types)
	if err != nil{
		fmt.Println("记录为空")
	}
	this.Data["article_types"] = article_types
	this.namespace = "article"
	this.Data["Namespace"] = "article"
	this.Data["PageTitle"] = "文章分类"
	this.setTpl("article_types/new.html")
	
}
//Create 创建
func (c *ArticleTypeController) Create()  {
	o := orm.NewOrm()
	name := c.GetString("TypeName")
	if name == ""{
		fmt.Println("名称为空")
		return
	}
	articleType := models.ArticleType{TypeName: name}
	_,err := o.Insert(&articleType)
	if err != nil{
		fmt.Println(err)
		return
	}
	c.Redirect("/article_type/create",302)
	
}