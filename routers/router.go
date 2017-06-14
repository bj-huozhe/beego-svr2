// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"beego-svr2/controllers"

	"github.com/astaxie/beego"
)

func init() {

	order_info_ns := beego.NewNamespace("/order",
		beego.NSRouter("/AddOrderInfo", &controllers.OrderController{}, "post:AddOrderInfo"),
		beego.NSRouter("/FindOrderById", &controllers.OrderController{}, "get:FindOrderById"),
	)
	team_info_ns := beego.NewNamespace("/team",
		beego.NSRouter("/AddTeamInfo", &controllers.TeamController{}, "post:AddTeamInfo"),
		beego.NSRouter("/FindTeamById", &controllers.TeamController{}, "get:FindTeamById"),
	)

	lend_info_ns := beego.NewNamespace("/lend",
		beego.NSRouter("/AddLendInfo", &controllers.LendController{}, "post:AddLendInfo"),
		beego.NSRouter("/FindLendInfoById", &controllers.LendController{}, "get:FindLendInfoById"),
	)

	lend_user_ns := beego.NewNamespace("/lendUser",
		beego.NSRouter("/AddLendUser", &controllers.LendUserController{}, "post:AddLendUser"),
		beego.NSRouter("/FindLendUserById", &controllers.LendUserController{}, "get:FindLendUserById"),
	)

	beego.AddNamespace(order_info_ns, team_info_ns, lend_info_ns, lend_user_ns)

}
