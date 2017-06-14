package controllers



import (
	"beego-svr2/pojo/request"
	"encoding/json"
	"github.com/astaxie/beego"
	"beego-svr2/service"
	"beego-svr2/util"
)




// Operations about Users
type LendUserController struct {
	beego.Controller
}

func (this *LendUserController) AddLendUser() {
	var lendUserVo request.LendUserVo
	json.Unmarshal(this.Ctx.Input.RequestBody, &lendUserVo)
	common.ConsoleLogs.Info("lendUserVo=", lendUserVo)
	this.Data["json"] = lendUserVo
	this.ServeJSON()
	service.AddLendUser(lendUserVo)
}

func (this *LendUserController) FindLendUserById() string {
	var result string
	id, err := this.GetInt("id")
	common.ConsoleLogs.Info("id=", id, "err=", err)
	if err != nil {
		panic(err)
	}
	lendUser := service.FindLendUserById(id)
	json, err := json.Marshal(lendUser)
	common.ConsoleLogs.Info("json=", json)
	if err != nil {
		common.ConsoleLogs.Error("error:", err)
	}
	result = string(json)
	common.ConsoleLogs.Info("result=", result)
	this.Ctx.WriteString(result)
	//return orderInfo
	//business.Ctx.WriteString(json)
	return result
}




