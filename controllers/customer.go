package controllers

import "C"
import (
	"encoding/json"
	"errors"
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
	c.Mapping("Get", c.GetAll)
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
	user := new(models.User)
	user.Id = c.currentUser.Id
	v.CreateUser = user
	status, errs := v.Validate()
	if status {
		if _, err := models.AddCustomer(&v); err == nil {
			c.jsonResult(200, "", "OK")
		} else {
			c.jsonResult(500, err.Error(), "")
		}
	} else {
		c.jsonResult(500, errs, "")
	}
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
		c.jsonResult(500, err.Error(), "")
	} else {
		c.jsonResult(200, "", v)
	}
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
				c.jsonResult(500, errors.New("Error: invalid query key/value pair"), "")
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
		c.jsonResult(500, err.Error(), "")
	} else {
		result := map[string]interface{}{
			"countPage": mapValue,
			"data":      l,
			"colNames":  colNames,
			"actions":   customerActions(),
		}
		c.jsonResult(200, "", result)
	}
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
			c.jsonResult(200, "", "OK")
		} else {
			c.jsonResult(500, err.Error(), "")
		}
	} else {
		c.jsonResult(500, errs, "")
	}
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
		c.jsonResult(200, "", "OK")
	} else {
		c.jsonResult(500, err.Error(), "")
	}
}

func (c *CustomerController) Get() {
	c.Data["JsName"] = "customer_index"
	c.Data["Namespace"] = "customer_manage"
	c.Data["PageTitle"] = "客户信息"
	c.setTpl("customer/index.html")
}
func (c *CustomerController) Index()  {
	c.Data["JsName"] = "customer_index"
	c.Data["Namespace"] = "customer_manage"
	c.Data["PageTitle"] = "客户信息"
	c.setTpl("customer/index.html")
}
func (c *CustomerController) New() {
	c.Data["JsName"] = "customer_form"
	c.Data["Namespace"] = "customer_manage"
	c.Data["PageTitle"] = "新增客户信息"
	c.setTpl("customer/form.html")
}
func (c *CustomerController) Edit() {
	c.Data["JsName"] = "customer_form"
	c.Data["Namespace"] = "customer_manage"
	c.Data["PageTitle"] = "修改客户信息"
	idStr := c.Ctx.Input.Params()["0"]
	c.Data["Id"] = idStr
	c.setTpl("customer/form.html")
}

func (c *CustomerController) Status() {
	actionType := c.GetString("actionType")
	var result []models.CustomerSlice
	returnJson := make(map[string]interface{})
	switch actionType {
	case "Status":
		result = models.CustomerStatusArray()
	case "AccountPeriod":
		result = models.CustomerAccountPeriodArray()
	case "CompanyType":
		result = models.CustomerTransportTypeArray()
	case "IsVip":
		result = models.CustomerIsVipArray()
	default:
		returnJson["Status"] = models.CustomerStatusArray()
		returnJson["AccountPeriod"] = models.CustomerAccountPeriodArray()
		returnJson["CompanyType"] = models.CustomerTransportTypeArray()
		returnJson["IsVip"] = models.CustomerIsVipArray()
	}
	if actionType != "all" {
		c.jsonResult(200, "", result)
	}else{
		c.jsonResult(200,"",returnJson)
	}
}
