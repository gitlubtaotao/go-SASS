package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type AddForeignKey_20191027_155112 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &AddForeignKey_20191027_155112{}
	m.Created = "20191027_155112"

	migration.Register("AddForeignKey_20191027_155112", m)
}

// Run the migrations
func (m *AddForeignKey_20191027_155112) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("ALTER TABLE user ADD FOREIGN KEY (company_id) REFERENCES company(id)")
	m.SQL("ALTER TABLE department ADD FOREIGN KEY (company_id) REFERENCES company(id)")
}

// Reverse the migrations
func (m *AddForeignKey_20191027_155112) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("ALTER TABLE user DROP FOREIGN KEY user_ibfk_1")
	m.SQL("ALTER TABLE department DROP FOREIGN KEY department_ibfk_1")
}
