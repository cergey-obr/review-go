package models

import (
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

func GetAll(offset int, limit int) []*Review {
	reviews := []*Review{}
	o := orm.NewOrm()
	queryReviewBuilder := o.QueryTable(Review{})
	queryReviewBuilder = addLimit(queryReviewBuilder, 1000)
	//queryBuilder = addOffset(queryBuilder, offset)
	queryReviewBuilder.All(&reviews)

	photos := []*Photo{}
	oP := orm.NewOrm()
	queryPhotoBuilder := oP.QueryTable(Photo{})
	cond := orm.NewCondition()
	for _, review := range reviews {
		cond = cond.Or("review__id", review.Id)
	}

	queryPhotoBuilder.SetCond(cond).All(&photos)

	for _, review := range reviews {
		for _, photo := range photos {
			if photo.Review.Id == review.Id {
				review.Photo = append(review.Photo, photo)
				continue
			}
		}
	}

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
