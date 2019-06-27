package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Review struct {
	Id        int               `orm:"auto;pk"`
	ProductId int               `orm:"index"`
	StatusId  int               `orm:"default(2)"`
	CreatedAt orm.DateTimeField `orm:"null,default(set_null)"`
	Title     string            `orm:"size(256)"`
	Author    *Author           `orm:"rel(fk)"`
	Detail    string
	Recommend bool `orm:"null"`
	Website   int
	Photo     []*Photo `orm:"reverse(many)"`
	Overall   int      `orm:"null"`
	Quality   int      `orm:"null"`
	Fit       int      `orm:"null"`
	Style     int      `orm:"null"`
}

func GetAll() []*Review {
	o := orm.NewOrm()
	var review []*Review
	num, err := o.QueryTable("review").All(&review)
	fmt.Printf("Returned Rows Num: %s, %s", num, err)

	return review
}

func GetProductReviews() []*Review {
	var review []*Review
	return review
}