package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"review-go/models"
	_ "review-go/routers"
)

func init() {
	// register model
	orm.RegisterModel(
		new(models.Author),
		new(models.Photo),
		new(models.Review))

	// set default database
	orm.RegisterDataBase(
		"default",
		"mysql",
		"root:123456@tcp(localhost:3307)/review",
		30)
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
