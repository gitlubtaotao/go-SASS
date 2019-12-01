package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/beego/i18n"
	"log"
	"quickstart/utils"
	"reflect"
	"strings"
	
	_ "github.com/go-sql-driver/mysql"
	
	"time"
)

//Company 公司信息
type Company struct {
	Id          int64         `orm:"pk;auto";form:"-"`
	Name        string        `orm:"size(128)";form:"Name"`
	Telephone   string        `orm:"size(128)"`
	Address     string        `orm:"size(256)"`
	Email       string        `orm:"size(128)"`
	Remarks     string        `orm:"size(128)"`
	Website     string        `orm:"size(128)"`
	CreatedAt   time.Time     `orm:"auto_now;type(datetime)"`
	User        []*User       `orm:"reverse(many)"`
	Department  []*Department `orm:"reverse(many)"`
	CompanyType string        `orm:"size(32)"`
}

func init() {
	
	orm.RegisterModel(new(Company))
}

//创建用户对应的验证
func (c *Company) Validate() (b bool, err map[string]interface{}) {
	status := true
	var returnErr map[string]interface{}
	returnErr = make(map[string]interface{})
	valid := validation.Validation{}
	valid.Required(c.Name, "name")
	valid.MaxSize(c.Name, 128, "max name")
	valid.Required(c.Telephone, "telephone")
	valid.Tel(c.Telephone, "telephone format")
	valid.Required(c.Email, "email")
	valid.Email(c.Email, "email format")
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

// AddCompany insert a new Company into database and returns
// last inserted Id on success.
func AddCompany(m *Company) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCompanyById retrieves Company by Id. Returns error if
// Id doesn't exist
func GetCompanyById(id int64) (v *Company, err error) {
	o := orm.NewOrm()
	v = &Company{Id: id}
	if err = o.QueryTable(new(Company)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCompany retrieves all Company matches certain condition. Returns empty list if
// no records exist
func GetAllCompany(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, countPage int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Company))
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
	
	var l []Company
	qs = qs.OrderBy(sortFields...).RelatedSel()
	if len(fields) == 0 {
		fields = StrutFields(new(Company))
	}
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
				if ArrayExistItem("CreatedAt", fields) {
					m["CreatedAt"] = utils.LongTime(v.CreatedAt)
				}
				ml = append(ml, m)
			}
		}
		return ml, count, nil
	}
	return nil, 0, err
}

// UpdateCompany updates Company by Id and returns error if
// the record to be updated doesn't exist
func UpdateCompanyById(m *Company) (err error) {
	o := orm.NewOrm()
	v := Company{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCompany deletes Company by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCompany(id int64) (err error) {
	o := orm.NewOrm()
	v := Company{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Company{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//获取对应的列
func GetCompanyCols(lang string) (array []CustomerSlice) {
	colNames := make([]CustomerSlice, 0)
	exceptColumns := []string{"Id", "CompanyType"}
	for _, column := range StrutFields(new(Company)) {
		if !ArrayExistItem(column, exceptColumns) {
			format := "company." + column
			hash := map[string]interface{}{
				"key": column, "value": i18n.Tr(lang, format), "class": "",
			}
			colNames = append(colNames, hash)
		}
	}
	return colNames
}
