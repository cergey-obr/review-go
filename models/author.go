package models

import "github.com/astaxie/beego/orm"

type Author struct {
	Id         int       `orm:"auto;pk"`
	Nickname   string    `orm:"size(128)"`
	Email      string    `orm:"size(128);null;unique"`
	CustomerId int       `orm:"index;default(0)"`
	Review     []*Review `orm:"reverse(many)"`
}

func init() {
	orm.RegisterModel(new(Author))
}