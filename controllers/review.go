package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/syyongx/php2go"
	"review-go/models"
)

// Operations about Users
type ReviewController struct {
	beego.Controller
}

// @Description get all reviews
// @Param order query []string "Order criterion"
// @Param limit query int 100 "Limit"
// @Param offset query int 0 "Offset"
// @Success 200 {object} models.Review
// @router / [get]
func (r *ReviewController) GetAll() {
	limit, _ := r.GetInt("limit", 100)
	offset, _ := r.GetInt("offset", 0)

	query := "filter%5Bfirst%5D=test&filter%5Bfourth%5D%5B%5D=123&filter%5Bfourth%5D%5B%5D=456&filter%5Bsecond%5D=hello&filter%5Bthird%5D=123&limit=t&offset=w"

	result := make(map[string]interface{})
	php2go.ParseStr(query, result)
	fmt.Println(result["filter"])

	//type column struct{
	//	Name string
	//	Data string
	//}
	//columns := make(map[string]string)
	//
	//if err := r.Ctx.Input.Bind(&columns, "filter"); err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println("OK, got:", columns)
	//}

	reviews := models.GetAll(offset, limit)

	r.Data["json"] = reviews
	r.ServeJSON()
}

func (r *ReviewController) ReviewController() {
	r.Data["json"] = "test"
	r.ServeJSON()
}
