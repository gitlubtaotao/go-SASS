package controllers

import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"quickstart/models"
	"strconv"
	"strings"
	
	"golang.org/x/crypto/bcrypt"
)

//  UserController operations for User
type UserController struct {
	BaseController
}

// URLMapping ...
func (c *UserController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create User
// @Param	body		body 	models.User	true		"body for User content"
// @Success 201 {int} models.User
// @Failure 403 body is empty
// @router / [post]
func (c *UserController) Post() {
	var v models.User
	o := orm.NewOrm()
	_ = json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	companyId, _ := strconv.Atoi(c.GetString("companyId"))
	company := models.Company{Id: int64(companyId)}
	_ = o.Read(&company)
	v.Company = &company
	//更新员工信息
	if v.Id != 0 {
		if v.Pwd != "" {
			encodePassword := c.generatePassword(v.Pwd)
			v.EncodePassword = encodePassword
			v.Pwd = ""
		}
		valid, vErrors := v.Validate()
		logs.Info(valid, vErrors)
		if valid {
			if err := models.UpdateUserById(&v); err == nil {
				c.Ctx.Output.SetStatus(201)
				c.Data["json"] = "OK"
			} else {
				c.Data["json"] = err.Error()
			}
		} else {
			c.Data["json"] = vErrors
		}
	} else {
		//创建员工信息
		encodePassword := c.generatePassword(v.Pwd)
		v.EncodePassword = encodePassword
		v.Pwd = ""
		//插入数据前进行验证
		valid, vErrors := v.Validate()
		logs.Info(valid, vErrors)
		if valid {
			if _, err := models.AddUser(&v); err == nil {
				c.Ctx.Output.SetStatus(201)
				c.Data["json"] = v
			} else {
				c.Data["json"] = err.Error()
			}
		} else {
			c.Data["json"] = vErrors
		}
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get User by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :id is empty
// @router /:id [get]
func (c *UserController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetUserById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get User
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.User
// @Failure 403
// @router / [get]
func (c *UserController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64
	page, _ := strconv.Atoi(c.GetString("page", "1"))
	offset = limit * (int64(page) - 1)
	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	if len(fields) == 0 {
		fields = append(fields, "Name",
			"Email", "Gender", "EntryTime",
			"Id", "Company")
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
	} else {
		sortby = append(sortby, "Id")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	} else {
		order = append(order, "desc")
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
	} else {
		query["Name"] = c.GetString("Name")
		query["Email"] = c.GetString("Email")
	}
	logs.Info(query)
	l, countPage, err := models.GetAllUser(query, fields, sortby, order, offset, limit)
	if err != nil {
		logs.Error(err)
		c.Data["json"] = err.Error()
	} else {
		mapValue := models.SetPaginator(countPage, int64(limit))
		logs.Info(mapValue)
		c.Data["json"] = map[string]interface{}{
			"countPage": mapValue,
			"editUrl":   "/user/edit/",
			"data":      l,
		}
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the User
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.User	true		"body for User content"
// @Success 200 {object} models.User
// @Failure 403 :id is not int
// @router /:id [put]
func (c *UserController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.User{Id: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	
	if err := models.UpdateUserById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the User
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *UserController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteUser(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

//用户列表
func (c *UserController) Get() {
	c.Data["Namespace"] = "company"
	c.Data["PageTitle"] = "员工信息"
	c.Data["JsName"] = "company"
	c.setTpl("user/index.html")
}

//新增用户
func (c *UserController) New() {
	c.namespace = "company"
	c.Data["JsName"] = "user_form"
	c.Data["Namespace"] = "company"
	c.Data["PageTitle"] = "新增员工信息"
	c.setTpl("user/new.html")
}

//修改用户

func (c *UserController) Edit() {
	c.Data["JsName"] = "user_form"
	c.Data["Namespace"] = "company"
	c.Data["PageTitle"] = "修改员工信息"
	logs.Info("dsdsdsssdsd")
	//获取 :Id
	idStr := c.Ctx.Input.Param(":id")
	logs.Info(idStr)
	id, _ := strconv.ParseInt(idStr, 0, 64)
	c.Data["UserId"] = id
	c.setTpl("user/edit.html")
}

//生成对应的密码
func (c *UserController) generatePassword(pwd string) (encodePassword string) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		logs.Error(err)
		return ""
	}
	encodePW := string(hash) // 保存在数据库的密码，虽然每次生成都不同，只需保存一份即可
	logs.Info(encodePW)
	return encodePW
}


