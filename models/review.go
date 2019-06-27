package models

<<<<<<<<< Temporary merge branch 1
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

func init() {
	orm.RegisterModel(new(Review))
}

func GetAll() []*Review {
	o := orm.NewOrm()
	queryBuilder := o.QueryTable(Review{})
	queryBuilder = addLimit(queryBuilder, limit)
	queryBuilder = addOffset(queryBuilder, offset)
	num, _ := queryBuilder.All(&reviews)

	fmt.Println(num)

	return reviews
}

func addLimit(o orm.QuerySeter, limit int) orm.QuerySeter {
	return o.Limit(limit)
}

func addOffset(o orm.QuerySeter, offset int) orm.QuerySeter {
	return o.Offset(offset)
}

//func addFilter(o orm.QuerySeter, filter int) orm.QuerySeter {
//
//	return o
//}

//func GetProductReviews() map[string]*Review {
//
//	return ReviewList
//}

func GetProductReviews() []*Review {
	var review []*Review
	return review
}