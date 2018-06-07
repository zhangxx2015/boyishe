// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"ManualApi/bll"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/users",
			beego.NSInclude(
				&bll.UserController{},
			),
		),

		beego.NSNamespace("/mgr",
			beego.NSInclude(
				&bll.MgrController{},
			),
		),
		beego.NSNamespace("/comm",
			beego.NSInclude(
				&bll.CommController{},
			),
		),
		beego.NSNamespace("/dbg",
			beego.NSInclude(
				&bll.DbgController{},
			),
		),
		beego.NSNamespace("/show",
			beego.NSInclude(
				&bll.ShowController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
