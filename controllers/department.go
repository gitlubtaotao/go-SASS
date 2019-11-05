package controllers

import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego/logs"
	"quickstart/models"
	"strconv"
	"strings"
)

//  DepartmentController operations for Department
type DepartmentController struct {
	BaseController
}

// URLMapping ...
func (c *DepartmentController) URLMapping() {
	logs.Info(c)
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Department
// @Param	body		body 	models.Department	true		"body for Department content"
// @Success 201 {int} models.Department
// @Failure 403 body is empty
// @router / [post]
func (c *DepartmentController) Post() {
	var v models.Department
	_ = json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	company := new(models.Company)
	company.Id, _ = strconv.ParseInt(c.GetString("company_id"), 0, 64)
	v.Company = company
	if _, err := models.AddDepartment(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = "OK"
		c.jsonResult(200, "", "OK")
	} else {
		c.jsonResult(500, err.Error(), "")
	}
}

// GetOne ...
// @Title Get One
// @Description get Department by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Department
// @Failure 403 :id is empty
// @router /:id [get]
func (c *DepartmentController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetDepartmentById(id)
	if err != nil {
		c.jsonResult(500, err.Error(), "")
	} else {
		c.jsonResult(200, "", v)
		c.Data["json"] = v
	}
}

// GetAll ...
// @Title Get All
// @Description get Department
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Department
// @Failure 403
// @router / [get]
func (c *DepartmentController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	limit := models.UserPerPage()
	var offset int64
	page, _ := strconv.Atoi(c.GetString("page", "1"))
	offset = models.GetOffsetPage(int64(page))
	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	_, colNames := models.GetDepartmentCols()
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
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
	
	l, countPage, err := models.GetAllDepartment(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.jsonResult(500, err.Error(), "")
	} else {
		mapValue := models.SetPaginator(countPage)
		result := map[string]interface{}{
			"countPage": mapValue,
			"data":      l,
			"colNames":  colNames,
			"actions":   departmentActions(),
		}
		c.jsonResult(200, "", result)
	}
}

func departmentActions() []models.CustomerSlice {
	actions := []models.CustomerSlice{
		{"name": "修改", "url": "/department/edit/:id", "remote": false},
		{"name": "删除", "url": "/department/:id", "remote": true, "method": "delete"},
	}
	return actions
}

// Put ...
// @Title Put
// @Description update the Department
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Department	true		"body for Department content"
// @Success 200 {object} models.Department
// @Failure 403 :id is not int
// @router /:id [put]
func (c *DepartmentController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Department{Id: id}
	_ = json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateDepartmentById(&v); err == nil {
		c.jsonResult(200, "", "OK")
	} else {
		c.jsonResult(500, err.Error(), "")
	}
}

// Delete ...
// @Title Delete
// @Description delete the Department
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *DepartmentController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteDepartment(id); err == nil {
		c.jsonResult(200, "OK", "")
	} else {
		c.jsonResult(500, err.Error(), "")
	}
}

func (c *DepartmentController) Get() {
	c.namespace = "company"
	c.Data["JsName"] = "index"
	c.Data["Namespace"] = "company"
	c.Data["PageTitle"] = "部门信息"
	c.setTpl("department/index.html")
}

//新增
func (c *DepartmentController) New() {
	c.namespace = "company"
	c.Data["Namespace"] = "company"
	c.Data["PageTitle"] = "新增部门信息"
	c.setTpl("department/form.html")
}

//
func (c *DepartmentController) Edit() {
	idStr := c.Ctx.Input.Param(":id")
	c.namespace = "company"
	c.Data["Namespace"] = "company"
	c.Data["PageTitle"] = "修改部门信息"
	c.Data["Id"] = idStr
	c.setTpl("department/form.html")
}
