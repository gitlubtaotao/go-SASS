package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AddCustomerKey_20191028_115955 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddCustomerKey_20191028_115955{}
	m.Created = "20191028_115955"

	migration.Register("AddCustomerKey_20191028_115955", m)
}

// Run the migrations
func (m *AddCustomerKey_20191028_115955) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE customer ADD FOREIGN KEY (company_id) REFERENCES company(id)")
	m.SQL("ALTER TABLE customer ADD FOREIGN KEY (audit_user_id) REFERENCES user(id)")
	m.SQL("ALTER TABLE customer ADD FOREIGN KEY (create_user_id) REFERENCES user(id)")
	m.SQL("ALTER TABLE customer ADD FOREIGN KEY (sale_user_id) REFERENCES user(id)")
}

// Reverse the migrations
func (m *AddCustomerKey_20191028_115955) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	//m.SQL("ALTER TABLE customer DROP FOREIGN KEY user_ibfk_1")
}
