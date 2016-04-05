// @APIVersion 1.0.0
// @Title Highlego API
// @Description  Highlego restful API
// @Contact anyshisong@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/shawncode/highlego-api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),

		beego.NSNamespace("/goods",
			beego.NSInclude(
				&controllers.GoodsController{},
			),
		),

		beego.NSNamespace("/lottery",
			beego.NSInclude(
				&controllers.LotteryController{},
			),
		),

		beego.NSNamespace("/lotteryOrder",
			beego.NSInclude(
				&controllers.LotteryOrderController{},
			),
		),

		beego.NSNamespace("/address",
			beego.NSInclude(
				&controllers.AddressController{},
			),
		),

		beego.NSNamespace("/exchange",
			beego.NSInclude(
				&controllers.ExchangeController{},
			),
		),

		beego.NSNamespace("/lotteryTopList",
			beego.NSInclude(
				&controllers.LotteryTopListController{},
			),
		),

		beego.NSNamespace("/shareBoard",
			beego.NSInclude(
				&controllers.ShareBoardController{},
			),
		),

		beego.NSNamespace("/usertasks",
			beego.NSInclude(
				&controllers.UsertasksController{},
			),
		),


	)
	beego.AddNamespace(ns)

}
