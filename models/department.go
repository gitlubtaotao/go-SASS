package models

import (
	"errors"
	"fmt"
	"github.com/beego/i18n"
	"quickstart/utils"
	"reflect"
	"strings"
	"time"
	
	"github.com/astaxie/beego/orm"
)

type Department struct {
	Id        int64     `orm:"pk;auto"`
	Name      string    `orm:"size(128)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"CreatedAt"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
	Company   *Company  `orm:"rel(fk);index" json:"Company"`
}

func init() {
	orm.RegisterModel(new(Department))
}

// AddDepartment insert a new Department into database and returns
// last inserted Id on success.
func AddDepartment(m *Department) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func GetDepartmentCols(lang string) []CustomerSlice {
	
	colNames := make([]CustomerSlice,0)
	exceptColumns := []string{"Id"}
	for _,column := range	StrutFields(new(Department)){
		if !ArrayExistItem(column,exceptColumns){
			format := "department."+column
			hash := map[string]interface{}{
				"key": column,"value": i18n.Tr(lang,format),"class": "",
			}
			colNames = append(colNames,hash)
		}
	}
	return  colNames
}

// GetDepartmentById retrieves Department by Id. Returns error if
// Id doesn't exist
func GetDepartmentById(id int64) (v *Department, err error) {
	o := orm.NewOrm()
	v = &Department{Id: id}
	if err = o.QueryTable(new(Department)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllDepartment retrieves all Department matches certain condition. Returns empty list if
// no records exist
func GetAllDepartment(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, countPage int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Department))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if v != "" {
			qs = qs.Filter(k, v)
		}
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
	var l []Department
	qs = qs.OrderBy(sortFields...).RelatedSel()
	count, _ := qs.Count()
	if len(fields) == 0 {
		fields = StrutFields(new(Department))
	}
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
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
			if ArrayExistItem("UpdatedAt",fields){
				m["UpdatedAt"] = utils.LongTime(v.UpdatedAt)
			}
			ml = append(ml, m)
		}
		return ml, count, nil
	}
	return nil, 0, err
}

// UpdateDepartment updates Department by Id and returns error if
// the record to be updated doesn't exist
func UpdateDepartmentById(m *Department) (err error) {
	o := orm.NewOrm()
	v := Department{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteDepartment deletes Department by Id and returns error if
// the record to be deleted doesn't exist
func DeleteDepartment(id int64) (err error) {
	o := orm.NewOrm()
	v := Department{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Department{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
