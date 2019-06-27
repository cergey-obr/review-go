package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"review-go/models"
)

// Operations about Users
type ReviewController struct {
	beego.Controller
}

// @Description get all reviews
// @Param filter query []string "Filter criterion"
// @Param order query []string "Order criterion"
// @Param limit query int 100 "Limit"
// @Param offset query int 0 "Offset"
// @Success 200 {object} models.Review
// @router / [get]
func (r *ReviewController) GetAll() {
	reviews := models.GetAllReviews()

	filters := r.GetStrings("filter[]")
	orders := r.GetStrings("order[]")
	limit, _ := r.GetInt("limit", 100)
	offset, _ := r.GetInt("offset", 0)

	fmt.Println("params ", filters, orders, limit, offset)

	r.Data["json"] = reviews
	r.ServeJSON()
}

// @Description Get reviews by product id
// @Param productId query []int "Array of product id"
// @Param filter query []string "Filter criterion"
// @Param order query []string "Order criterion"
// @Param limit query int 100 "Limit"
// @Param offset query int 0 "Offset"
// @Success 200 {string} login success
// @router /product [get]
func (r *ReviewController) GetProductReviews() {
	reviews := models.GetProductReviews()

	r.Data["json"] = reviews
	r.ServeJSON()
}
