package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/validation"
	"github.com/beego/i18n"
	"log"
	"quickstart/utils"
	"reflect"
	"strings"
	"time"
	
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id             int64  `orm:"pk;auto"`
	Name           string `orm:"size(128);unique"`
	Email          string `orm:"size(128);unique"`
	Phone          string `orm:"size(64);unique"`
	EncodePassword string `orm:"size(512);unique"`
	Pwd            string
	Gender         string      `orm:"size(64)"`
	Positions      []*Position `orm:"rel(m2m)"`
	EntryTime      time.Time
	CreatedAt      time.Time   `orm:"auto_now_add;type(datetime)" json:"CreatedAt"`
	UpdatedAt      time.Time   `orm:"auto_now;type(datetime)"`
	Company        *Company    `orm:"rel(fk);index" json:"Company"`
	Department     *Department `orm:"rel(fk);index;NULL" json:"Department"`
}

func init() {
	orm.RegisterModel(new(User))
}

// AddUser insert a new User into database and returns
// last inserted Id on success.
//真的密码的生成
func AddUser(m *User) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	logs.Info(err)
	return id, err
}

// GetUserById retrieves User by Id. Returns error if
// Id doesn't exist
func GetUserById(id int64) (v *User, err error) {
	o := orm.NewOrm()
	v = &User{Id: id}
	if err = o.QueryTable(new(User)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllUser retrieves all User matches certain condition. Returns empty list if
// no records exist
func GetAllUser(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, countPage int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(User))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if v != "" {
			qs = qs.Filter(k, v)
		}
	}
	count, _ := qs.Count()
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
	var l []User
	qs = qs.OrderBy(sortFields...).RelatedSel()
	if len(fields) == 0 {
		fields = StrutFields(new(User))
	}
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		// trim unused fields
		for _, v := range l {
			m := make(map[string]interface{})
			val := reflect.ValueOf(v)
			for _, fname := range fields {
				m[fname] = val.FieldByName(fname).Interface()
			}
			if ArrayExistItem("EntryTime", fields) {
				m["EntryTime"] = utils.LongTime(v.EntryTime)
			}
			if ArrayExistItem("UpdatedAt", fields) {
				m["UpdatedAt"] = utils.LongTime(v.UpdatedAt)
			}
			m["Company.Name"] = v.Company.Name
			m["Company.Email"] = v.Company.Email
			ml = append(ml, m)
		}
		return ml, count, nil
	}
	return nil, 0, err
}

// UpdateUser updates User by Id and returns error if
// the record to be updated doesn't exist
func UpdateUserById(m *User) (err error) {
	o := orm.NewOrm()
	v := User{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

//获取用户对应的colNames
func GetUserCols(lang string) (cols []CustomerSlice) {
	colNames := make([]CustomerSlice,0)
	exceptColumns := []string{"Id","CreatedAt","EncodePassword","Pwd"}
	for _,column := range	StrutFields(new(User)){
		if !ArrayExistItem(column,exceptColumns){
			format := "user."+column
			hash := map[string]interface{}{
				"key": column,"value": i18n.Tr(lang,format),"class": "",
			}
			colNames = append(colNames,hash)
		}
	}
	return colNames
}

// DeleteUser deletes User by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUser(id int64) (err error) {
	o := orm.NewOrm()
	v := User{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&User{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//创建用户的对应validate

func (u *User) Validate(lang string) (b bool, errors map[string]interface{}) {
	valid := validation.Validation{}
	valid.Required(u.Name, i18n.Tr(lang,"user.Name")).Message(i18n.Tr(lang,"error.required"))
	valid.MaxSize(u.Name, 128, i18n.Tr(lang,"user.Name")).Message(i18n.Tr(lang,"error.max_size")+",%d",128)
	valid.Phone(u.Phone, i18n.Tr(lang,"user.Phone")).Message(i18n.Tr(lang,"error.phone"))
	valid.Email(u.Email, i18n.Tr(lang,"user.Email")).Message(i18n.Tr(lang,"error.email"))
	valid.Required(u.EncodePassword,i18n.Tr(lang,"user.Pwd")).Message(i18n.Tr(lang,"error.required"))
	valid.Required(u.Company, i18n.Tr(lang,"user.Company")).Message(i18n.Tr(lang,"error.required"))
	var returnErr map[string]interface{}
	returnErr = make(map[string]interface{})
	status := true
	logs.Error(valid.HasErrors())
	if valid.HasErrors() {
		status = false
		for _, err := range valid.Errors {
			returnErr[err.Key] = err.Message
			log.Println(err.Key, err.Message)
		}
	}
	return status, returnErr
}
