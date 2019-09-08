package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"path"
	"quickstart/models"
	"time"
)

//ArticleController 文章controller
type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) Get() {
	c.Layout="layouts/application.html"
	c.TplName = "articles/new.html"
}

func (c *ArticleController) Add() {
	c.Layout="layouts/application.html"
	o := orm.NewOrm()
	var articleTypes []models.ArticleType
	_,err := o.QueryTable("ArticleType").All(&articleTypes)
	if err != nil{
		fmt.Println("记录为空")
	}
	c.Data["article_types"] = articleTypes
	c.Data["pageTitle"] = "添加文章"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["ContentHeader"] = "layouts/header.html"
	c.TplName = "articles/new.html"
}

//Post() 创建文章
func (c *ArticleController) Post() {
	articleName := c.GetString("articleName")
	articleContent := c.GetString("content")
	c.Layout="layouts/application.html"
	//上传文件
	f, h, err := c.GetFile("uploadname")
	defer f.Close()
	if err != nil {
		c.TplName = "articles/new.html"
		return
	}
	fileext := path.Ext(h.Filename)
	fmt.Println(fileext)
	if fileext != ".jpg" && fileext != ".png" {
		c.TplName = "articles/new.html"
		fmt.Println("上传图片类型错误")
		return
	}
	if h.Size > 5000000 {
		c.TplName = "articles/new.html"
		fmt.Println("上传图片过大")
		return
	}
	filename := time.Now().Format("2006-01-02 15:04:05") + h.Filename
	_ = c.SaveToFile("uploadname", "./static/upload/"+filename)
	if articleContent == "" || articleName == "" {
		c.TplName = "articles/new.html"
		return
	}
	println(articleName, articleContent)
	o := orm.NewOrm()
	atricle := models.Article{}
	atricle.ArtiName = articleName
	atricle.Acontent = articleContent
	atricle.Aimg = "/static/upload/" + filename
	articleTypeId, errSelect := c.GetInt("select")
	if errSelect != nil{
		return
	}
	fmt.Println(articleTypeId)
	articleType := models.ArticleType{Id: articleTypeId}
	_ = o.Read(&articleType)
	atricle.ArticleType = &articleType
	_, err = o.Insert(&atricle)
	if err != nil {
		c.TplName = "articles/new.html"
		fmt.Println(err)
		return
	}
	
	c.Redirect("/", 302)
	
}



//Show 文章详情
func (c *ArticleController) Show() {
	c.Layout="layouts/application.html"
	id, err := c.GetInt("id")
	if err != nil {
		logs.Debug("获取文章错误", err)
		return
	}
	o := orm.NewOrm()
	article := models.Article{Id: id}
	err = o.Read(&article)
	if err != nil {
		logs.Debug(err)
		return
	}
	//多对多插入
	m2m := o.QueryM2M(&article,"Users")
	userName :=  c.GetSession("userName")
	user := models.User{}
	user.Name = userName.(string)
	_ = o.Read(&user,"Name")
	_, _ = m2m.Add(&user)
	article.Account += 1
	_, _ = o.Update(&article)
	
	//many to many search
	//_, _ = o.LoadRelated(&article, "Users",1,1000,0,"-id")
	//logs.Info(article.Users)
	
	//_= o.QueryTable("Article").Filter("Users__User__Name","admin").One(&article)
	//logs.Info(article)
	
	var users []models.User
	_, _ = o.QueryTable("User").Filter("Articles__Article__Id", id).Filter("Name", userName).Distinct().All(&users)
	//var articles []models.Article
	//_, err = o.QueryTable("Article").Filter("Users__User__Name", userName).All(&articles)
	logs.Info(users)
	c.Data["article"] = article
	c.Data["users"] = users
	c.TplName = "articles/content.html"
}

//Edit 编辑
func (c *ArticleController) Edit() {
	c.Layout="layouts/application.html"
	id, err := c.GetInt("id")
	if err != nil {
		fmt.Println("错误", err)
	}
	o := orm.NewOrm()
	article := models.Article{Id: id}
	err = o.Read(&article)
	if err != nil {
		fmt.Println("错误", err)
	}
	c.Data["article"] = article
	c.Data["pageTitle"] = "编辑文章详情"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["ContentHeader"] = "layouts/header.html"
	c.TplName = "articles/update.html"
}

//Update 更新
func (c *ArticleController) Update() {
	c.Layout="layouts/application.html"
	f, h, err := c.GetFile("uploadname")
	filename := ""
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	} else {
		fileext := path.Ext(h.Filename)
		fmt.Println(fileext)
		if fileext != ".jpg" && fileext != ".png" {
			c.TplName = "articles/new.html"
			fmt.Println("上传图片类型错误")
			return
		}
		if h.Size > 5000000 {
			c.TplName = "articles/new.html"
			fmt.Println("上传图片过大")
			return
		}
		filename = time.Now().Format("2006-01-02 15:04:05") + h.Filename
		_ = c.SaveToFile("uploadname", "./static/upload/"+filename)
	}
	articleName := c.GetString("articleName")
	articleContent := c.GetString("content")
	Id,idErr := c.GetInt("Id")
	fmt.Println(Id)
	if idErr != nil{
		fmt.Println("获取Id 失败")
		return
	}
	if articleName == "" || articleContent == ""{
		fmt.Println("更加数据库失败")
		return
	}
	o := orm.NewOrm()
	atricle := models.Article{Id: Id}
	err = o.Read(&atricle)
	if err != nil{
		fmt.Println(err)
		c.TplName = "articles/update.html"
	}
	atricle.ArtiName = articleName
	atricle.Acontent = articleContent
	atricle.Id = Id
	if filename != "" {
		atricle.Aimg = "/static/upload/" + filename
	}
	_,err  =  o.Update(&atricle)
	if err != nil{
		fmt.Println("保存失败")
		c.TplName="articles/update.html"
	}
	
	c.Redirect("/", 302)
}
//Delete: 删除请求
func (c *ArticleController) Delete()  {

}
