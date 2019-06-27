package models

import "github.com/astaxie/beego/orm"

type Photo struct {
	Id     int     `orm:"auto;pk"`
	Review *Review `orm:"rel(fk)"`
	Url    string  `orm:"size(128);null"`
}

func init() {
	orm.RegisterModel(new(Photo))
}
