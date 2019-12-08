package controllers

import (
	"encoding/json"
	"errors"
	"github.com/beego/i18n"
	"quickstart/models"
	"quickstart/utils"
	"strconv"
	"strings"
)

//  ContactController operations for Contact
type ContactController struct {
	BaseController
}

// URLMapping ...
func (c *ContactController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Contact
// @Param	body		body 	models.Contact	true		"body for Contact content"
// @Success 201 {int} models.Contact
// @Failure 403 body is empty
// @router / [post]
func (c *ContactController) Post() {
	var v models.Contact
	_ = json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if status,err := v.Validate();status{
		if _, err := models.AddContact(&v); err == nil {
			c.jsonResult(200,"","OK")
		} else {
			c.jsonResult(500,err.Error(),"")
		}
	}else{
		c.jsonResult(500,err,"")
	}
}

// GetOne ...
// @Title Get One
// @Description get Contact by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Contact
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ContactController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetContactById(id)
	if err != nil {
		c.jsonResult(500,err.Error(),"")
	} else {
		c.jsonResult(200,"",v)
	}
}

// GetAll ...
// @Title Get All
// @Description get Contact
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Contact
// @Failure 403
// @router / [get]
func (c *ContactController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var offset int64
	limit := models.UserPerPage()
	page, _ := strconv.Atoi(c.GetString("page", "1"))
	offset = models.GetOffsetPage(int64(page))
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
	
	l, countNumber, err := models.GetAllContact(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.jsonResult(500, err.Error(), "")
		return
	}
	_, colNames := models.GetContactCols()
	if c.GetString("format") != "" {
		c.DownLoad(l, colNames)
	} else {
		mapValue := models.SetPaginator(countNumber)
		result := map[string]interface{}{
			"countPage": mapValue,
			"data":      l,
			"colNames":  colNames,
			"actions":   contactActions(),
		}
		c.jsonResult(200, "", result)
	}
}

//操作actions
func contactActions() []models.CustomerSlice {
	actions := []models.CustomerSlice{
		{"name": "修改", "url": "/contact/edit/:id", "remote": false},
		{"name": "删除", "url": "/contact/:id", "remote": true, "method": "delete"},
	}
	return actions
}

// Put ...
// @Title Put
// @Description update the Contact
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Contact	true		"body for Contact content"
// @Success 200 {object} models.Contact
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ContactController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Contact{Id: id}
	if status,err := v.Validate();status {
		_ = json.Unmarshal(c.Ctx.Input.RequestBody, &v)
		if err := models.UpdateContactById(&v); err == nil {
			c.jsonResult(200, "", "OK")
		} else {
			c.jsonResult(500, err.Error(), "")
		}
	}else{
		c.jsonResult(500,err,"")
	}
}

// Delete ...
// @Title Delete
// @Description delete the Contact
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ContactController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteContact(id); err == nil {
		c.jsonResult(200,"","OK")
	} else {
		c.jsonResult(500,err.Error(),"")
	}
}

func (c *ContactController) Index() {
	c.Data["JsName"] = "index"
	c.Data["Namespace"] = "customer_manage"
	c.Data["PageTitle"] = i18n.Tr(c.Lang,"module_name.contact")
	c.setTpl("contact/index.html")
}
func (c *ContactController) Edit() {
	c.Data["JsName"] = "index"
	c.Data["Namespace"] = "customer_manage"
	c.Data["PageTitle"] = utils.LocaleS(i18n.Tr(c.Lang,"edit"),i18n.Tr(c.Lang,"module_name.contact"))
	idStr := c.Ctx.Input.Params()["0"]
	c.Data["Id"] = idStr
	c.setTpl("contact/form.html")
}

func (c *ContactController) New() {
	c.Data["JsName"] = ""
	c.Data["Namespace"] = "customer_manage"
	c.Data["PageTitle"] = utils.LocaleS(i18n.Tr(c.Lang,"edit"),i18n.Tr(c.Lang,"module_name.contact"))
	c.setTpl("contact/form.html")
}
