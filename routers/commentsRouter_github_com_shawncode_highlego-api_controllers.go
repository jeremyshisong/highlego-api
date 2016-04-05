package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/shawncode/highlego-api/controllers:GoodsController"] = append(beego.GlobalControllerRouter["github.com/shawncode/highlego-api/controllers:GoodsController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shawncode/highlego-api/controllers:GoodsController"] = append(beego.GlobalControllerRouter["github.com/shawncode/highlego-api/controllers:GoodsController"],
		beego.ControllerComments{
			"GetALLGoodsCate",
			`/cates`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shawncode/highlego-api/controllers:GoodsController"] = append(beego.GlobalControllerRouter["github.com/shawncode/highlego-api/controllers:GoodsController"],
		beego.ControllerComments{
			"GetGoodsByCate",
			`/cate/:cid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shawncode/highlego-api/controllers:GoodsController"] = append(beego.GlobalControllerRouter["github.com/shawncode/highlego-api/controllers:GoodsController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shawncode/highlego-api/controllers:LotteryController"] = append(beego.GlobalControllerRouter["github.com/shawncode/highlego-api/controllers:LotteryController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shawncode/highlego-api/controllers:LotteryController"] = append(beego.GlobalControllerRouter["github.com/shawncode/highlego-api/controllers:LotteryController"],
		beego.ControllerComments{
			"GetLotteryByGid",
			`/goods/:gid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shawncode/highlego-api/controllers:LotteryController"] = append(beego.GlobalControllerRouter["github.com/shawncode/highlego-api/controllers:LotteryController"],
		beego.ControllerComments{
			"GetAll",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shawncode/highlego-api/controllers:LotteryOrderController"] = append(beego.GlobalControllerRouter["github.com/shawncode/highlego-api/controllers:LotteryOrderController"],
		beego.ControllerComments{
			"Post",
			`/`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["github.com/shawncode/highlego-api/controllers:LotteryOrderController"] = append(beego.GlobalControllerRouter["github.com/shawncode/highlego-api/controllers:LotteryOrderController"],
		beego.ControllerComments{
			"GetOne",
			`/:id`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shawncode/highlego-api/controllers:LotteryOrderController"] = append(beego.GlobalControllerRouter["github.com/shawncode/highlego-api/controllers:LotteryOrderController"],
		beego.ControllerComments{
			"GetWaitByUser",
			`/wait/user/:uid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shawncode/highlego-api/controllers:LotteryOrderController"] = append(beego.GlobalControllerRouter["github.com/shawncode/highlego-api/controllers:LotteryOrderController"],
		beego.ControllerComments{
			"GetPublishedByUser",
			`/all/user/:uid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shawncode/highlego-api/controllers:LotteryOrderController"] = append(beego.GlobalControllerRouter["github.com/shawncode/highlego-api/controllers:LotteryOrderController"],
		beego.ControllerComments{
			"GetWinByUser",
			`/win/user/:uid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shawncode/highlego-api/controllers:LotteryOrderController"] = append(beego.GlobalControllerRouter["github.com/shawncode/highlego-api/controllers:LotteryOrderController"],
		beego.ControllerComments{
			"GetFailByUser",
			`/fail/user/:uid`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shawncode/highlego-api/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/shawncode/highlego-api/controllers:UserController"],
		beego.ControllerComments{
			"Login",
			`/login/:deviceId`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["github.com/shawncode/highlego-api/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/shawncode/highlego-api/controllers:UserController"],
		beego.ControllerComments{
			"UpdateUser",
			`/:id`,
			[]string{"put"},
			nil})

}
