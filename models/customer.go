package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/validation"
	"github.com/beego/i18n"
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
	offset int64, limit int64, typeValue string,lang string) (ml []interface{}, countPage int64, err error) {
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
				m["Status"] = v.ShowStatus(true,lang).(string)
			}
			if ArrayExistItem("IsVip", fields) {
				m["IsVip"] = v.ShowVip(true,lang).(string)
			}
			if ArrayExistItem("CompanyType", fields) {
				m["CompanyType"] = v.ShowCompanyType(true,lang).(string)
			}
			if ArrayExistItem("AccountPeriod", fields) {
				m["AccountPeriod"] = v.ShowPeriod(true,lang).(string)
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

func GetCustomerCols(lang string) []CustomerSlice {
	colNames := make([]CustomerSlice, 0)
	exceptColumns := []string{"Id","UpdatedAt"}
	for _, column := range StrutFields(new(Customer)) {
		if !ArrayExistItem(column, exceptColumns) {
			format := "customer." + column
			hash := map[string]interface{}{
				"key": column, "value": i18n.Tr(lang, format), "class": "",
			}
			colNames = append(colNames, hash)
		}
	}
	return  colNames
}

//状态数组
func CustomerStatusArray(lang string) []CustomerSlice {
	data := []CustomerSlice{
		{"label": i18n.Tr(lang,"customer.status/init"), "code": "init"},
		{"label": i18n.Tr(lang,"customer.status/pass"), "code": "pass"},
		{"label": i18n.Tr(lang,"customer.status/fail"), "code": "fail"},
	}
	return data
}

func (c *Customer) ShowStatus(isString bool,lang string) interface{} {
	data := CustomerStatusArray(lang)
	for _, v := range data {
		if v["code"] == c.Status {
			if isString {
				return v["label"]
			} else {
				return v
			}
		}
	}
	return i18n.Tr(lang,"customer.status/init")
}

//获取对应的账期
func CustomerAccountPeriodArray(lang string) []CustomerSlice {
	data := []CustomerSlice{
		{"label": i18n.Tr(lang,"customer.accountPeriod/month"), "code": "month"},
		{"label": i18n.Tr(lang,"customer.accountPeriod/ticket"), "code": "ticket"},
	}
	return data
}

//显示账期
func (c *Customer) ShowPeriod(isString bool,lang string) interface{} {
	data := CustomerAccountPeriodArray(lang)
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
func CustomerTransportTypeArray(lang string) []CustomerSlice {
	data := []CustomerSlice{
		{"label": i18n.Tr(lang,"customer.companyType/1"), "code": 1},
		{"label": i18n.Tr(lang,"customer.companyType/2"), "code": 2},
		{"label": i18n.Tr(lang,"customer.companyType/3"), "code": 3},
	}
	return data
}

//显示对应的公司类型
func (c *Customer) ShowCompanyType(isString bool,lang string) interface{} {
	data := CustomerTransportTypeArray(lang)
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
func CustomerIsVipArray(lang string) []CustomerSlice {
	data := []CustomerSlice{
		{"label": i18n.Tr(lang,"customer.isVip/true"), "code": true,},
		{"label": i18n.Tr(lang,"customer.isVip/false"), "code": false},
	}
	return data
}
func (c *Customer) ShowVip(isString bool,lang string) interface{} {
	data := CustomerStatusArray(lang)
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
func (c *Customer) Validate(lang string) (b bool, err map[string]interface{}) {
	b = true
	valid := validation.Validation{}
	valid.Required(c.Name, i18n.Tr(lang,"customer.Name"))
	valid.MaxSize(c.Name, 128, i18n.Tr(lang,"error.max"))
	valid.Required(c.Telephone, i18n.Tr(lang,"customer.Telephone"))
	valid.Tel(c.Telephone, i18n.Tr(lang,"error.phone"))
	valid.Required(c.Email,i18n.Tr(lang,"customer.Email") )
	valid.Email(c.Email, i18n.Tr(lang,"error.email"))
	valid.Required(c.AccountPeriod, i18n.Tr(lang,"customer.AccountPeriod"))
	valid.Required(c.Company, i18n.Tr(lang,"customer.Company"))
	valid.Required(c.AuditUser, i18n.Tr(lang,"customer.AuditUser"))
	valid.Required(c.SaleUser, i18n.Tr(lang,"customer.SaleUser"))
	valid.Required(c.CreateUser, i18n.Tr(lang,"customer.CreateUser"))
	err = make(map[string]interface{})
	if valid.HasErrors() {
		b = false
		for _, item := range valid.Errors {
			err[item.Key] = item.Message
		}
	}
	return b, err
}
