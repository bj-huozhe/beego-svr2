package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"beego-svr2/service"
	"beego-svr2/util"
	"beego-svr2/pojo/request"
)


// Operations about Users
type LendBankController struct {
	beego.Controller
}


func (this *LendBankController) AddLendBank() {
	var modelInfoVo request.LendBankVo
	json.Unmarshal(this.Ctx.Input.RequestBody, &modelInfoVo)
	common.ConsoleLogs.Info("modelInfoVo=", modelInfoVo)
	this.Data["json"] = modelInfoVo
	this.ServeJSON()
	service.AddLendBank(modelInfoVo)
}




func (this *LendBankController) FindLendBankById() string{
	var result string
	//id, err := business.GetInt("id")
	id, err:= this.GetInt("id")
	common.ConsoleLogs.Info("err=", err)
	if err != nil {
			panic(err)
		}
	common.ConsoleLogs.Info("id=", id)
	modelInfoVo := service.FindLendBankById(id)
	json, err := json.Marshal(modelInfoVo)
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


func (this *LendBankController) FindLendBankAll() string{
	var result string
	modelInfoList := service.FindLendBankAll()
	json, err := json.Marshal(modelInfoList)
	common.ConsoleLogs.Info("json=", json)
	if err != nil {
	    common.ConsoleLogs.Error("error:", err)
	}
	result = string(json)
	common.ConsoleLogs.Info("result=", result)
	this.Ctx.WriteString(result)
	//return orderInfo
	//business.Ctx.WriteString(json)
	//this.ServeJSONP()
	return result
}

func (this *LendBankController) UpdateLendBank() {
	var modelInfoVo request.LendBankVo
	json.Unmarshal(this.Ctx.Input.RequestBody, &modelInfoVo)
	common.ConsoleLogs.Info("modelInfoVo=", modelInfoVo)
	this.Data["json"] = modelInfoVo
	this.ServeJSON()
	service.UpdateLendBank(modelInfoVo)
}

func (this *LendBankController) FindLendBankAllList() {
	modelInfoList := service.FindLendBankAll()
	this.Data["json"] = modelInfoList
	this.ServeJSON()
}










