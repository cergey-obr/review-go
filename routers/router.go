// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"review-go/controllers"
)

func init() {
	reviewNS := beego.NewNamespace("/:website([0-9]+)/review",
		beego.NSInclude(
			&controllers.ReviewController{},
		),
	)

	photoNS := beego.NewNamespace("/photo",
		beego.NSInclude(
			&controllers.ReviewController{},
		),
	)

	beego.AddNamespace(reviewNS, photoNS)
}
