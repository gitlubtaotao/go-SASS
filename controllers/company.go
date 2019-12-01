package controllers

import (
	"encoding/json"
	"errors"
	"github.com/beego/i18n"
	"quickstart/utils"
	"strings"
	
	"quickstart/models"
	"strconv"
)

// CompanyController operations for Companies
type CompanyController struct {
	BaseController
}

// URLMapping ...
func (c *CompanyController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

//New 新增公司信息
func (this *CompanyController) New() {
	this.namespace = "company"
	this.Data["Namespace"] = "company"
	this.Data["PageTitle"] = utils.LocaleS(i18n.Tr(this.Lang,"new"),
		i18n.Tr(this.Lang,"module_name.company"))
	this.setTpl("companies/form.html")
}

// Post ...
// @Title Create
// @Description create Companies
// @Param	body		body 	models.Companies	true		"body for Companies content"
// @Success 201 {object} models.Companies
// @Failure 403 body is empty
// @router / [post]
func (c *CompanyController) Post() {
	company := models.Company{}
	_ = json.Unmarshal(c.Ctx.Input.RequestBody, &company)
	status, errors := company.Validate()
	if status {
		if _, err := models.AddCompany(&company); err == nil {
			c.jsonResult(200, "", "OK")
		} else {
			c.jsonResult(500, err.Error(), "")
		}
	} else {
		c.jsonResult(500, errors, "")
	}
}

//Get 首页
func (c *CompanyController) Index() {
	c.Data["JsName"] = "company_index"
	c.Data["Namespace"] = "company"
	c.Data["PageTitle"] = i18n.Tr(c.Lang,"module_name.company")
	c.setTpl("companies/index.html")
}

// GetOne ...
// @Title GetOne
// @Description get Companies by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Companies
// @Failure 403 :id is empty
// @router /:id [get]
func (c *CompanyController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if v, err := models.GetCompanyById(id); err != nil {
		c.jsonResult(500, err.Error(), "")
	} else {
		c.jsonResult(200, "", v)
	}
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
func (c *CompanyController) GetAll() {
	var fields []string
	var query = make(map[string]string)
	sortBy := make([]string, 1)
	order := make([]string, 1)
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	sortBy[0] = "Id"
	order[0] = "desc"
	//进行数据的分页
	limit := models.UserPerPage()
	page, _ := strconv.Atoi(c.GetString("page", "1"))
	offset := models.GetOffsetPage(int64(page))
	
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.jsonResult(500, errors.New("Error: invalid query key/value pair"), "")
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	colNames := models.GetCompanyCols(c.Lang)
	companies, countPage, err := models.GetAllCompany(query, fields, sortBy, order, offset, limit)
	if err != nil {
		c.jsonResult(500, err.Error(), "")
		return
	}
	if c.GetString("format") != "" {
		c.DownLoad(companies, colNames)
		return
	}
	mapValue := models.SetPaginator(countPage)
	result := map[string]interface{}{
		"countPage": mapValue,
		"data":      companies,
		"colNames":  colNames,
		"actions":   c.companyActions(),
	}
	c.jsonResult(200, "", result)
}

func (c *CompanyController)companyActions() []models.CustomerSlice {
	actions := []models.CustomerSlice{
		{"name": i18n.Tr(c.Lang,"edit"), "url": "/company/edit/:id", "remote": false},
		{"name": i18n.Tr(c.Lang,"delete"), "url": "/company/:id", "remote": true, "method": "delete"},
	}
	return actions
}

// Put ...
// @Title Put
// @Description update the Companies
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Companies	true		"body for Companies content"
// @Success 200 {object} models.Companies
// @Failure 403 :id is not int
// @router /:id [put]
func (c *CompanyController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Company{Id: id}
	_ = json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	status, errors := v.Validate()
	if status {
		if err := models.UpdateCompanyById(&v); err == nil {
			c.jsonResult(200, "", "OK")
		} else {
			c.jsonResult(500, err.Error(), "")
		}
	} else {
		c.jsonResult(500, errors, "")
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Companies
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *CompanyController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteCompany(id); err == nil {
		c.jsonResult(200, "", "OK")
	} else {
		c.jsonResult(500, err.Error(), "")
	}
}

func (this *CompanyController) Edit() {
	idStr := this.Ctx.Input.Params()["0"]
	this.namespace = "company"
	this.Data["Namespace"] = "company"
	this.Data["PageTitle"] = utils.LocaleS(i18n.Tr(this.Lang,"edit"),i18n.Tr(this.Lang,"module_name.company"))
	this.Data["Id"] = idStr
	this.setTpl("companies/form.html")
}
