package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/validation"
	"quickstart/utils"
	"reflect"
	"strings"
	"time"
	
	"github.com/astaxie/beego/orm"
)

//客户信息
type Customer struct {
	Id               int64     `orm:"pk;auto"`
	Name             string    `orm:"size(128);unique"`
	Telephone        string    `orm:"size(128);unique"`
	Address          string    `orm:"size(256)"`
	Email            string    `orm:"size(128);unique"`
	Remarks          string    `orm:"size(128)"`
	Website          string    `orm:"size(128)"`
	CreatedAt        time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt        time.Time `orm:"auto_now;type(datetime)"`
	CompanyType      int       `orm:"size(32);default(1)"`
	Aging            int
	Amount           int        `orm:"digits(12);decimals(4)"`
	AccountPeriod    string     `orm:"size(32);"`
	IsVip            bool       `orm:"default(false)"`
	Status           string     `orm:"size(32);default(init);"`
	AuditUser        *User      `orm:"rel(fk);index"`
	CreateUser       *User      `orm:"rel(fk);index"`
	SaleUser         *User      `orm:"rel(fk);index;NULL"`
	Company          *Company   `orm:"rel(fk);index" json:"Company"`
	BusinessTypeName string     `orm:"size(256)"`
	Contacts         []*Contact `orm:"reverse(many)"`
	//Contacts         []*cooperator.Contact `orm:"reverse(many)"`
}

func init() {
	orm.RegisterModel(new(Customer))
}

// AddCustomer insert a new Customer into database and returns
// last inserted Id on success.
func AddCustomer(m *Customer) (id int64, err error) {
	m.Status = "init"
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCustomerById retrieves Customer by Id. Returns error if
// Id doesn't exist
func GetCustomerById(id int64) (v *Customer, err error) {
	o := orm.NewOrm()
	v = &Customer{Id: id}
	if err = o.QueryTable(new(Customer)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCustomer retrieves all Customer matches certain condition. Returns empty list if
// no records exist
func GetAllCustomer(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64, typeValue string) (ml []interface{}, countPage int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Customer))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if v != "" {
			if k == "company_type" {
				logs.Info(query, "sdsds")
			} else {
				qs = qs.Filter(k, v)
			}
		}
	}
	//查询不同的客户类型
	if typeValue == "customer" {
		qs = qs.Filter("company_type__in", 0, 1, 3)
	} else if typeValue == "supplier" {
		qs = qs.Filter("company_type__in", 2, 3)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, 0, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, 0, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, 0, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, 0, errors.New("Error: unused 'order' fields")
		}
	}
	var l []Customer
	qs = qs.OrderBy(sortFields...).RelatedSel()
	count, _ := qs.Count()
	if len(fields) == 0 {
		fields = StrutFields(new(Customer))
	}
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		for _, v := range l {
			m := make(map[string]interface{})
			val := reflect.ValueOf(v)
			for _, fname := range fields {
				m[fname] = val.FieldByName(fname).Interface()
			}
			if ArrayExistItem("Status", fields) {
				m["Status"] = v.ShowStatus(true).(string)
			}
			if ArrayExistItem("IsVip", fields) {
				m["IsVip"] = v.ShowVip(true).(string)
			}
			if ArrayExistItem("CompanyType", fields) {
				m["CompanyType"] = v.ShowCompanyType(true).(string)
			}
			if ArrayExistItem("AccountPeriod", fields) {
				m["AccountPeriod"] = v.ShowPeriod(true).(string)
			}
			if ArrayExistItem("CreatedAt", fields) {
				m["CreatedAt"] = utils.LongTime(v.CreatedAt)
			}
			ml = append(ml, m)
		}
		return ml, count, nil
	}
	return nil, 0, err
}

// UpdateCustomer updates Customer by Id and returns error if
// the record to be updated doesn't exist
func UpdateCustomerById(m *Customer) (err error) {
	o := orm.NewOrm()
	v := Customer{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCustomer deletes Customer by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCustomer(id int64) (err error) {
	o := orm.NewOrm()
	v := Customer{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Customer{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func GetCustomerCols() ([]string, []CustomerSlice) {
	var fields []string
	colNames := []CustomerSlice{
		{"key": "Name", "value": "公司名称", "class": ""},
		{"key": "Telephone", "value": "公司电话", "class": ""},
		{"key": "Email", "value": "公司邮箱", "class": ""},
		{"key": "CompanyType", "value": "类型", "class": ""},
		{"key": "AccountPeriod", "value": "账期", "class": ""},
		{"key": "Aging", "value": "账龄", "class": ""},
		{"key": "Amount", "value": "额度", "class": ""},
		{"key": "IsVip", "value": "是否VIP", "class": ""},
		{"key": "BusinessTypeName", "value": "业务类型", "class": ""},
		{"key": "Status", "value": "状态", "class": ""},
		{"key": "CreatedAt", "value": "创建时间", "class": ""},
		{"key": "AuditUser", "value": "审核者", "class": ""},
		{"key": "CreateUser", "value": "创建者", "class": ""},
		{"key": "SaleUser", "value": "业务员", "class": ""},
		{"key": "Company", "value": "所属公司", "class": ""},
	}
	return fields, colNames
}

//状态数组
func CustomerStatusArray() []CustomerSlice {
	data := []CustomerSlice{
		{"label": "等待审核", "code": "init"},
		{"label": "审核通过", "code": "pass"},
		{"label": "审核失败", "code": "fail"},
	}
	return data
}

func (c *Customer) ShowStatus(isString bool) interface{} {
	data := CustomerStatusArray()
	for _, v := range data {
		if v["code"] == c.Status {
			if isString {
				return v["label"]
			} else {
				return v
			}
		}
	}
	return "等待审核"
}

//获取对应的账期
func CustomerAccountPeriodArray() []CustomerSlice {
	data := []CustomerSlice{
		{"label": "月结", "code": "month"},
		{"label": "票结", "code": "ticket"},
	}
	return data
}

//显示账期
func (c *Customer) ShowPeriod(isString bool) interface{} {
	data := CustomerAccountPeriodArray()
	for _, v := range data {
		if v["code"] == c.AccountPeriod {
			if isString {
				return v["label"].(string)
			} else {
				return v
			}
		}
	}
	return ""
}
func CustomerTransportTypeArray() []CustomerSlice {
	data := []CustomerSlice{
		{"label": "客户", "code": 1},
		{"label": "供应商", "code": 2},
		{"label": "客户&供应商", "code": 3},
	}
	return data
}

//显示对应的公司类型
func (c *Customer) ShowCompanyType(isString bool) interface{} {
	data := CustomerTransportTypeArray()
	for _, v := range data {
		if v["code"] == c.CompanyType {
			if isString {
				return v["label"]
			} else {
				return v
			}
		}
	}
	return ""
}
func CustomerIsVipArray() []CustomerSlice {
	data := []CustomerSlice{
		{"label": "是", "code": true,},
		{"label": "否", "code": false},
	}
	return data
}
func (c *Customer) ShowVip(isString bool) interface{} {
	data := CustomerStatusArray()
	for _, v := range data {
		if c.IsVip == v["code"] {
			if isString {
				return v["code"]
			} else {
				return v
			}
		}
	}
	return ""
}

//创建客户对应的验证
func (c *Customer) Validate() (b bool, err map[string]interface{}) {
	b = true
	valid := validation.Validation{}
	valid.Required(c.Name, "客户名称")
	valid.MaxSize(c.Name, 128, "最大值")
	valid.Required(c.Telephone, "联系电话")
	valid.Required(c.Telephone, "联系电话")
	valid.Tel(c.Telephone, "联系电话格式")
	valid.Required(c.Email, "邮箱")
	valid.Email(c.Email, "邮箱格式")
	valid.Required(c.AccountPeriod, "账期")
	valid.Required(c.Company, "所属公司")
	valid.Required(c.AuditUser, "审核者")
	valid.Required(c.SaleUser, "业务员")
	valid.Required(c.CreateUser, "创建者")
	err = make(map[string]interface{})
	if valid.HasErrors() {
		b = false
		for _, item := range valid.Errors {
			err[item.Key] = item.Message
		}
	}
	return b, err
}
