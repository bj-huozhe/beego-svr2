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

	AddOrderInfo := beego.NewNamespace("/business",
	    beego.NSRouter("/AddOrderInfo", &controllers.BusinessController{}, "post:AddOrderInfo"),
	    )
	FindById := beego.NewNamespace("/business",
	    beego.NSRouter("/FindById/:id([0-9]+", &controllers.BusinessController{}, "get:FindById"),
	    )
	beego.AddNamespace(AddOrderInfo,FindById)
}
