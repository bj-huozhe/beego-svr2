package controllers


import (
	"beego-svr2/pojo/request"
	"encoding/json"
	"github.com/astaxie/beego"
	"beego-svr2/service"
	"beego-svr2/util"
)




// Operations about Users
type LendController struct {
	beego.Controller
}

func (this *LendController) AddLendInfo() {
	var lendInfoVo request.LendInfoVo
	json.Unmarshal(this.Ctx.Input.RequestBody, &lendInfoVo)
	common.ConsoleLogs.Info("lendInfoVo=", lendInfoVo)
	this.Data["json"] = lendInfoVo
	this.ServeJSON()
	service.AddLendInfo(lendInfoVo)
}

func (this *LendController) FindLendInfoById() string {
	var result string
	id, err := this.GetInt("id")
	common.ConsoleLogs.Info("id=", id, "err=", err)
	if err != nil {
		panic(err)
	}
	lendInfo := service.FindLendInfoById(id)
	json, err := json.Marshal(lendInfo)
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






















