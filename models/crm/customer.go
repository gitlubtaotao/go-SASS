package crm

import (
	"errors"
	"fmt"
	"quickstart/models"
	"quickstart/models/oa"
	"reflect"
	"strings"
	"time"
	
	"github.com/astaxie/beego/orm"
)

//客户信息
type Customer struct {
	Id               int64     `orm:"pk;auto"`
	Name             string    `orm:"size(128);unique;NOT NUL"`
	Telephone        string    `orm:"size(128)"`
	Address          string    `orm:"size(256)"`
	Email            string    `orm:"size(128);unique"`
	Remarks          string    `orm:"size(128)"`
	Website          string    `orm:"size(128)"`
	CreatedAt        time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt        time.Time `orm:"auto_now;type(datetime)"`
	CompanyType      int64     `orm:"size(32);default(1)"`
	Aging            int64
	Amount           float64 `orm:"digits(12);decimals(4)"`
	AccountPeriod    int
	IsVip            bool        `orm:"default(false)"`
	Status           string      `orm:"size(32);default(init)"`
	AuditUser        *oa.User    `orm:"rel(fk);index"`
	CreateUser       *oa.User    `orm:"rel(fk);index"`
	SaleUser         *oa.User    `orm:"rel(fk);index"`
	Company          *oa.Company `orm:"rel(fk);index"`
	BusinessTypeName string      `orm:"size(256)"`
}

func init() {
	orm.RegisterModel(new(Customer))
}

// AddCustomer insert a new Customer into database and returns
// last inserted Id on success.
func AddCustomer(m *Customer) (id int64, err error) {
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
	offset int64, limit int64) (ml []interface{}, countPage int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Customer))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
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
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, count,nil
	}
	return nil,0, err
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

func GetCustomerCols() ([]string, []models.CustomerSlice) {
	var fields []string
	colNames := []models.CustomerSlice{
		{"key": "Name", "value": "公司名称", "class": ""},
		{"key": "Telephone", "value": "公司电话", "class": ""},
		{"key": "Email", "value": "公司邮箱", "class": ""},
		{"key": "AccountPeriod", "value": "账期", "class": ""},
		{"key": "Aging", "value": "账龄", "class": ""},
		{"key": "Amount", "value": "额度", "class": ""},
		{"key": "IsVip", "value": "是否VIP", "class": ""},
		{"key": "BusinessTypeName", "value": "业务类型", "class": ""},
		{"key": "status", "value": "状态", "class": ""},
		{"key": "CreatedAt", "value": "创建时间", "class": ""},
		{"key": "AuditUser", "value": "审核者", "class": ""},
		{"key": "CreateUser", "value": "创建者", "class": ""},
		{"key": "SaleUser", "value": "业务员", "class": ""},
		{"key": "Company", "value": "所属公司", "class": ""},
	}
	return fields, colNames
}
