package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"math"
	"quickstart/models"
	"strconv"
)

//MainController init
type MainController struct {
	beego.Controller
}

//Get method
func (c *MainController) Get() {
	userName := c.GetSession("userName")
	if userName == nil {
		c.Redirect("/login",302)
		return
	}
	//o := orm.NewOrm()
	//user := models.User{}
	//user.Name="taotao"
	//user.Pwd="qweqwe123"
	//_,err := o.Insert(&user)
	//if err != nil{
	//	fmt.Println("插入数据失败",err)
	//}
	
	//根据Id 进行查询
	//o := orm.NewOrm()
	//user := models.User{}
	//user.Id = 1
	//err := o.Read(&user)
	
	//更新
	//o := orm.NewOrm()
	//user := models.User{}
	//user.Name="taotao"
	//if o.Read(&user,"Name") == nil{
	//  user.Name = "MyName"
	//  user.Pwd="123456"
	//  if num, err := o.Update(&user,"Name","Pwd"); err == nil {
	//	  fmt.Println(num)
	//  }
	//}
	//fmt.Println(user)
	//c.Data["User"] = user
	
	//删除对象
	//o := orm.NewOrm()
	//user := models.User{}
	//user.id = 1
	//_,err := o.Delete(&user)
	//if err != nil{
	//	fmt.Println("删除失败",user)
	//}
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	
	//查询对象
	o := orm.NewOrm()
	var articles []models.Article
	qs := o.QueryTable("Article")
	//_, err := qs.All(&articles)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	
	//fmt.Println(Count)
	pageSize := 10
	pageIndex, err := strconv.Atoi(c.GetString("pageIndex"))
	if err != nil {
		pageIndex = 1
	}
	start := pageSize * (pageIndex - 1)
	temp_value := qs.Limit(pageSize, start).RelatedSel("ArticleType")
	select_type := c.GetString("select_value")
	var Count int64
	if select_type != "" {
		_, _ = temp_value.Filter("ArticleType__TypeName", select_type).All(&articles)
		Count, _ = temp_value.Filter("ArticleType__TypeName", select_type).Count()
	} else {
		_, _ = temp_value.All(&articles)
		Count, _ = temp_value.Count()
	}
	pageCount := math.Ceil(float64(Count) / float64(pageSize))
	c.Data["articles"] = articles
	c.Data["count"] = Count
	c.Data["pageIndex"] = pageIndex
	c.Data["pageCount"] = pageCount
	if pageIndex == 1 {
		c.Data["firstPage"] = true
	} else {
		c.Data["firstPage"] = false
	}
	o = orm.NewOrm()
	var articleTypes []models.ArticleType
	_, err = o.QueryTable("ArticleType").All(&articleTypes)
	if err != nil {
		fmt.Println("记录为空")
	}
	c.Data["article_types"] = articleTypes
	c.Data["selectValue"] = select_type
	c.Data["pageTitle"] = "文章列表"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["ContentHeader"] = "layouts/header.html"
	c.Layout="layouts/application.html"
	c.TplName = "index.html"
}
