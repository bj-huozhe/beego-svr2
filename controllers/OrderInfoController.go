package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"beego-svr2/service"
	"beego-svr2/util"
	"beego-svr2/pojo/request"
)


// Operations about Users
type OrderController struct {
	beego.Controller
}


func (this *OrderController) AddOrderInfo() {
	var orderInfoVo request.OrderInfoVo
	json.Unmarshal(this.Ctx.Input.RequestBody, &orderInfoVo)
	common.ConsoleLogs.Info("orderInfoVo=", orderInfoVo)
	this.Data["json"] = orderInfoVo
	this.ServeJSON()
	service.AddOrderInfo(orderInfoVo)
}




func (this *OrderController) FindOrderById() string{
	var result string
	//id, err := business.GetInt("id")
	id, err:= this.GetInt("id")
	common.ConsoleLogs.Info("err=", err)
	if err != nil {
			panic(err)
		}
	common.ConsoleLogs.Info("id=", id)
	orderInfo := service.FindOrderById(id)
	json, err := json.Marshal(orderInfo)
	common.ConsoleLogs.Info("json=", json)
	if err != nil {
	    common.ConsoleLogs.Error("error:", err)
	}
	result = string(json)
	common.ConsoleLogs.Info("result=", result)
	//this.Ctx.WriteString(string(json))
	//return orderInfo
	//business.Ctx.WriteString(json)
	return result
}