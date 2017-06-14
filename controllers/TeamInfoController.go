package controllers

import (
	"beego-svr2/pojo/request"
	"encoding/json"
	"github.com/astaxie/beego"
	"beego-svr2/service"
	"beego-svr2/util"
)




// Operations about Users
type TeamController struct {
	beego.Controller
}

func (this *TeamController) AddTeamInfo() {
	var teamInfoVo request.TeamInfoVo
	json.Unmarshal(this.Ctx.Input.RequestBody, &teamInfoVo)
	common.ConsoleLogs.Info("teamInfoVo=", teamInfoVo)
	this.Data["json"] = teamInfoVo
	this.ServeJSON()
	service.AddTeamInfo(teamInfoVo)
}

func (this *TeamController) FindTeamById() string {
	var result string
	id, err := this.GetInt("id")
	common.ConsoleLogs.Info("id=", id, "err=", err)
	if err != nil {
		panic(err)
	}
	teamInfo := service.FindTeamById(id)
	json, err := json.Marshal(teamInfo)
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