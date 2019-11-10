package controllers

import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego/logs"
	"quickstart/models"
	"strconv"
	"strings"
)

// SupplierController operations for Supplier
type SupplierController struct {
	BaseController
}

// URLMapping ...
func (c *SupplierController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Supplier
// @Param	body		body 	models.Supplier	true		"body for Supplier content"
// @Success 201 {object} models.Supplier
// @Failure 403 body is empty
// @router / [post]
func (c *SupplierController) Post() {
	var v models.Customer
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	logs.Error(err)
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
// @Title GetOne
// @Description get Supplier by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Supplier
// @Failure 403 :id is empty
// @router /:id [get]
func (c *SupplierController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if v, err := models.GetCustomerById(id); err != nil {
		c.jsonResult(500, err.Error(), "")
	} else {
		c.jsonResult(200, "", v)
	}
}

// GetAll ...
// @Title GetAll
// @Description get Supplier
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Supplier
// @Failure 403
// @router / [get]
func (c *SupplierController) GetAll() {
	var fields []string
	var sortBy [] string
	var order []string
	var query = make(map[string]string)
	limit := models.UserPerPage()
	page, _ := strconv.Atoi(c.GetString("page", "1"))
	offset := models.GetOffsetPage(int64(page))
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	if v := c.GetString("sortby"); v != "" {
		sortBy = strings.Split(v, ",")
	}
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
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
	_, colNames := models.GetCustomerCols()
	l, countPage, err := models.GetAllCustomer(query, fields, sortBy,
		order, offset, limit,
		"supplier")
	if err != nil {
		c.jsonResult(500, err.Error(), "")
	} else {
		mapValue := models.SetPaginator(countPage)
		result := map[string]interface{}{
			"countPage": mapValue,
			"data":      l,
			"colNames":  colNames,
			"actions":   c.actions(),
		}
		c.jsonResult(200, "", result)
	}
}

//获取对应的actions
func (c *SupplierController) actions() []models.CustomerSlice {
	actions := []models.CustomerSlice{
		{"name": "修改", "url": "/supplier/edit/:id", "remote": false},
		{"name": "详情", "url": "/supplier/show/:id", "remote": false},
		{"name": "删除", "url": "/supplier/:id", "remote": true, "method": "delete"},
	}
	return actions
}

// Put ...
// @Title Put
// @Description update the Supplier
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Supplier	true		"body for Supplier content"
// @Success 200 {object} models.Supplier
// @Failure 403 :id is not int
// @router /:id [put]
func (c *SupplierController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Customer{Id: id}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	logs.Error(err)
	status, errs := v.Validate()
	if status {
		if err = models.UpdateCustomerById(&v); err == nil {
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
// @Description delete the Supplier
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *SupplierController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteCustomer(id); err == nil {
		c.jsonResult(200, "", "OK")
	} else {
		c.jsonResult(500, err.Error(), "")
	}
}

func (c *SupplierController) Index() {
	c.Data["JsName"] = "customer_index"
	c.Data["Namespace"] = "customer_manage"
	c.Data["PageTitle"] = "供应商信息"
	c.setTpl("customer/index.html")
}
func (c *SupplierController) Edit() {
	c.Data["JsName"] = "customer_form"
	idStr := c.Ctx.Input.Params()["0"]
	c.Data["Namespace"] = "customer_manage"
	c.Data["PageTitle"] = "修改供应商信息"
	c.Data["Id"] = idStr
	c.setTpl("customer/form.html")
}
func (c *SupplierController) New() {
	c.Data["JsName"] = "customer_form"
	c.Data["Namespace"] = "customer_manage"
	c.Data["PageTitle"] = "新增供应商信息"
	c.setTpl("customer/form.html")
}
