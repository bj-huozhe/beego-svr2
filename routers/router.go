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
		beego.NSRouter("/FindOrderAll", &controllers.OrderController{}, "get:FindOrderAll"),
		beego.NSRouter("/UpdateOrder", &controllers.OrderController{}, "post:UpdateOrder"),
	)
	team_info_ns := beego.NewNamespace("/team",
		beego.NSRouter("/AddTeamInfo", &controllers.TeamController{}, "post:AddTeamInfo"),
		beego.NSRouter("/FindTeamById", &controllers.TeamController{}, "get:FindTeamById"),
		beego.NSRouter("/FindTeamAll", &controllers.TeamController{}, "get:FindTeamAll"),
		beego.NSRouter("/UpdateTeam", &controllers.TeamController{}, "post:UpdateTeam"),
	)

	lend_info_ns := beego.NewNamespace("/lend",
		beego.NSRouter("/AddLendInfo", &controllers.LendController{}, "post:AddLendInfo"),
		beego.NSRouter("/FindLendInfoById", &controllers.LendController{}, "get:FindLendInfoById"),
		beego.NSRouter("/FindLendInfoAll", &controllers.LendController{}, "get:FindLendInfoAll"),
		beego.NSRouter("/UpdateLendInfo", &controllers.LendController{}, "post:UpdateLendInfo"),
	)

	lend_user_ns := beego.NewNamespace("/lendUser",
		beego.NSRouter("/AddLendUser", &controllers.LendUserController{}, "post:AddLendUser"),
		beego.NSRouter("/FindLendUserById", &controllers.LendUserController{}, "get:FindLendUserById"),
		beego.NSRouter("/FindLendUserAll", &controllers.LendUserController{}, "get:FindLendUserAll"),
		beego.NSRouter("/UpdateLendUser", &controllers.LendUserController{}, "post:UpdateLendUser"),
	)

	lend_bank_ns := beego.NewNamespace("/lendBank",
		beego.NSRouter("/AddLendBank", &controllers.LendBankController{}, "post:AddLendBank"),
		beego.NSRouter("/FindLendBankById", &controllers.LendBankController{}, "get:FindLendBankById"),
		beego.NSRouter("/FindLendBankAll", &controllers.LendBankController{}, "get:FindLendBankAll"),
		beego.NSRouter("/FindLendBankAllList", &controllers.LendBankController{}, "get:FindLendBankAllList"),
		beego.NSRouter("/UpdateLendBank", &controllers.LendBankController{}, "post:UpdateLendBank"),
	)

	beego.AddNamespace(order_info_ns, team_info_ns, lend_info_ns, lend_user_ns, lend_bank_ns)

}
