package controllers

import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego/logs"
	"quickstart/models"
	"strconv"
	"strings"
)

//  CustomerController operations for Customer
type CustomerController struct {
	BaseController
}

// URLMapping ...
func (c *CustomerController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Customer
// @Param	body		body 	models.Customer	true		"body for Customer content"
// @Success 201 {int} models.Customer
// @Failure 403 body is empty
// @router / [post]
func (c *CustomerController) Post() {
	var v models.Customer
	//logs.Info(c.Ctx.Input.RequestBody)
	_ = json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	//RequestBody失效
	auditId := c.GetString("audit_id")
	saleId := c.GetString("sale_id")
	user := new(models.User)
	company := new(models.Company)
	user.Id = c.currentUser.Id
	v.CreateUser = user
	user.Id, _ = strconv.ParseInt(auditId, 0, 64)
	v.AuditUser = user
	user.Id, _ = strconv.ParseInt(saleId, 0, 64)
	v.SaleUser = user
	v.AccountPeriod = c.GetString("period")
	v.CompanyType, _ = strconv.Atoi(c.GetString("company_type"))
	company.Id, _ = strconv.ParseInt(c.GetString("company_id"),0,64)
	v.Company = company
	status, errs := v.Validate()
	if status {
		if _, err := models.AddCustomer(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		logs.Info(errs)
		c.Data["json"] = errs
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Customer by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Customer
// @Failure 403 :id is empty
// @router /:id [get]
func (c *CustomerController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetCustomerById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Customer
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Customer
// @Failure 403
// @router / [get]
func (c *CustomerController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	limit := models.UserPerPage()
	page, _ := strconv.Atoi(c.GetString("page", "1"))
	offset := models.GetOffsetPage(int64(page))
	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
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
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}
	_, colNames := models.GetCustomerCols()
	l, countPage, err := models.GetAllCustomer(query, fields, sortby, order, offset, limit, "customer")
	mapValue := models.SetPaginator(countPage)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = map[string]interface{}{
			"countPage": mapValue,
			"data":      l,
			"colNames":  colNames,
			"actions":   customerActions(),
		}
	}
	c.ServeJSON()
}

func customerActions() []models.CustomerSlice {
	actions := []models.CustomerSlice{
		{"name": "修改", "url": "/customer/edit/:id", "remote": false},
		{"name": "详情", "url": "/customer/show/:id", "remote": false},
		{"name": "删除", "url": "/customer/:id", "remote": true, "method": "delete"},
	}
	return actions
}

// Put ...
// @Title Put
// @Description update the Customer
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Customer	true		"body for Customer content"
// @Success 200 {object} models.Customer
// @Failure 403 :id is not int
// @router /:id [put]
func (c *CustomerController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Customer{Id: id}
	_ = json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	status, errs := v.Validate()
	if status {
		if err := models.UpdateCustomerById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = errs
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Customer
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *CustomerController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteCustomer(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

func (c *CustomerController) Get() {
	c.Data["JsName"] = "customer_index"
	c.Data["Namespace"] = "customer_manage"
	c.Data["PageTitle"] = "客户信息"
	c.setTpl("customer/index.html")
}
func (c *CustomerController) New() {
	c.Data["Namespace"] = "customer_manage"
	c.Data["PageTitle"] = "新增客户信息"
	c.setTpl("customer/form.html")
}
func (c *CustomerController) Edit() {

}

func (c *CustomerController) GetStatus() {
	actionType := c.GetString("actionType")
	switch actionType {
	case "Status":
		c.Data["json"] = models.CustomerStatusArray()
	case "AccountPeriod":
		c.Data["json"] = models.CustomerAccountPeriodArray()
	case "CompanyType":
		c.Data["json"] = models.CustomerTransportTypeArray()
		
	}
	c.ServeJSON()
}
