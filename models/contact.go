package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/validation"
	"quickstart/utils"
	"reflect"
	"strings"
	"time"
	
	"github.com/astaxie/beego/orm"
)

type Contact struct {
	Id        int64     `orm:"auto"`
	Name      string    `orm:"size(128);unique" description:"联系人姓名"`
	Email     string    `orm:"size(64);unique" description:"邮箱"`
	Phone     string    `orm:"size(64);unique"`
	WeiXinNo  string    `orm:"size(64);unique"`
	QqNo      string    `orm:"size(64);unique"`
	Address   string    `orm:"size(256);null"`
	Customer  *Customer `orm:"rel(fk);index" json:"Customer"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"CreatedAt"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Contact))
}

//获取数据库的表名
func (c *Contact) TableName() string {
	return TableName("contact")
}

// AddContact insert a new Contact into database and returns
// last inserted Id on success.
func AddContact(m *Contact) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetContactById retrieves Contact by Id. Returns error if
// Id doesn't exist
func GetContactById(id int64) (v *Contact, err error) {
	o := orm.NewOrm()
	v = &Contact{Id: id}
	if err = o.QueryTable(new(Contact)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllContact retrieves all Contact matches certain condition. Returns empty list if
// no records exist
func GetAllContact(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, countNumber int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Contact))
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
	
	var l []Contact
	qs = qs.OrderBy(sortFields...).RelatedSel()
	count, _ := qs.Count()
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			fields = StrutFields(new(Contact))
		}
		// trim unused fields
		for _, v := range l {
			m := make(map[string]interface{})
			val := reflect.ValueOf(v)
			for _, fname := range fields {
				m[fname] = val.FieldByName(fname).Interface()
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


// UpdateContact updates Contact by Id and returns error if
// the record to be updated doesn't exist
func UpdateContactById(m *Contact) (err error) {
	o := orm.NewOrm()
	v := Contact{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteContact deletes Contact by Id and returns error if
// the record to be deleted doesn't exist
func DeleteContact(id int64) (err error) {
	o := orm.NewOrm()
	v := Contact{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Contact{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

func GetContactCols() ([]string, []CustomerSlice) {
	attributes := []CustomerSlice{
		{"key": "Name", "value": "姓名", "class": "col-xs-1"},
		{"key": "Email", "value": "邮箱", "class": "col-xs-1"},
		{"key": "Phone", "value": "电话", "class": ""},
		{"key": "WeiXinNo","value": "微信号","class": ""},
		{"key": "QqNo","value": "QQ号","class": ""},
		{"key": "Address","value": "家庭地址","class": ""},
		{"key": "Customer", "value": "合作单位", "class": "col-xs-1"},
		{"key": "CreatedAt","value": "创建时间", "class": ""},
	}
	var fields []string
	return fields, attributes
}

func (u *Contact) Validate() (b bool, errors map[string]interface{}) {
	valid := validation.Validation{}
	valid.Required(u.Name, "姓名")
	valid.MaxSize(u.Name, 128, "姓名最大值")
	valid.Phone(u.Phone, "电话")
	valid.Email(u.Email, "邮箱")
	valid.MaxSize(u.Email, 128, "邮箱最大值")
	valid.Required(u.Customer, "合作单位")
	var returnErr =  make(map[string]interface{})
	status := true
	if valid.HasErrors() {
		status = false
		for _, err := range valid.Errors {
			returnErr[err.Key] = err.Message
		}
	}
	return status, returnErr
}
