package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type Addcolumntocompany20191011220632 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Addcolumntocompany20191011220632{}
	m.Created = "20191011_220632"
	
	_ = migration.Register("AddColumnToCompany_20191011_220632", m)
}

// Run the migrations
func (m *Addcolumntocompany20191011220632) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE company ADD company_type VARCHAR(32)")
}

// Reverse the migrations
func (m *Addcolumntocompany20191011220632) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("ALTER TABLE company  DROP COLUMN company_type")
}

