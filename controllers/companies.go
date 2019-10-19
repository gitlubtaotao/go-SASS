package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"strings"
	
	"quickstart/models"
	"strconv"
)

// CompaniesController operations for Companies
type CompaniesController struct {
	BaseController
}

// URLMapping ...
//func (c *CompaniesController) URLMapping() {
//	c.Mapping("Post", c.Post)
//	c.Mapping("GetOne", c.GetOne)
//	c.Mapping("GetAll", c.GetAll)
//	c.Mapping("Put", c.Put)
//	c.Mapping("Delete", c.Delete)
//
//}

//New 新增公司信息
func (this *CompaniesController) New() {
	this.namespace = "company"
	this.Data["Namespace"] = "company"
	this.Data["PageTitle"] = "新增公司"
	this.Data["JsName"] = "company_form"
	this.setTpl("companies/new.html")
}

// Post ...
// @Title Create
// @Description create Companies
// @Param	body		body 	models.Companies	true		"body for Companies content"
// @Success 201 {object} models.Companies
// @Failure 403 body is empty
// @router / [post]
func (c *CompaniesController) Post() {
	company := models.Company{}
	_ = json.Unmarshal(c.Ctx.Input.RequestBody, &company)
	if err := c.ParseForm(&company); err != nil {
		fmt.Print(err)
	} else {
		o := orm.NewOrm()
		status, returnErr := company.Validate()
		if status {
			if _, err = o.Insert(&company); err != nil {
				c.setTpl("company/new")
				return
			}
			c.redirectCustomer("/company")
		} else {
			fmt.Print(returnErr)
		}
	}
}

//Get 首页
func (c *CompaniesController) Get() {
	c.Data["JsName"] = "company"
	c.Data["Namespace"] = "company"
	c.Data["PageTitle"] = "公司信息"
	c.LayoutSections = make(map[string]string)
	c.setTpl("companies/index.html")
}

// GetOne ...
// @Title GetOne
// @Description get Companies by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Companies
// @Failure 403 :id is empty
// @router /:id [get]
func (c *CompaniesController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Companies
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Companies
// @Failure 403
// @router / [get]
func (c *CompaniesController) GetAll() {
	var fields []string
	companyFiler := map[string]string{
		"Name":      c.GetString("Name"),
		"Telephone": c.GetString("Telephone"),
		"Email":     c.GetString("Email"),
		"Address":   c.GetString("Address"),
	}
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	if len(fields) == 0 {
		fields = append(fields, "Id", "Name", "Telephone", "Address",
			"Email", "Website", "Remarks", "CreatedAt")
	}
	sortBy := make([]string, 1)
	order := make([]string, 1)
	sortBy[0] = "Id"
	order[0] = "desc"
	//进行数据的分页
	limit := models.UserPerPage()
	var offset int64
	page, _ := strconv.Atoi(c.GetString("page", "1"))
	offset = models.GetOffsetPage(int64(page))
	colNames := models.GetCompanyCols()
	companies, countPage, err := models.GetAllCompany(companyFiler, fields, sortBy, order, offset, limit)
	if err != nil {
		logs.Error(err)
		c.ServeJSON()
	} else {
		mapValue := models.SetPaginator(countPage)
		c.Data["json"] = map[string]interface{}{
			"countPage": mapValue,
			"data":      companies,
			"colNames":  colNames,
			"actions":   map[string]string{"edit": "company/edit/:id", "destroy": ""},
		}
		c.ServeJSON()
	}
}

// Put ...
// @Title Put
// @Description update the Companies
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Companies	true		"body for Companies content"
// @Success 200 {object} models.Companies
// @Failure 403 :id is not int
// @router /:id [put]
func (c *CompaniesController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Companies
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *CompaniesController) Delete() {
	var result map[string]string
	result = make(map[string]string)
	if c.IsAjax() {
		id := c.Ctx.Input.Param(":id")
		newId, _ := strconv.ParseInt(id, 10, 64)
		company := models.Company{Id: newId}
		o := orm.NewOrm()
		if _, err := o.Delete(&company); err == nil {
		} else {
			result["message"] = "删除失败"
			c.Data["json"] = result
		}
		c.ServeJSON()
	} else {
		c.redirectCustomer("/company")
	}
}

func (c *CompaniesController) Edit() {

}
