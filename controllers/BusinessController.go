package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"beego-svr2/models"
	"beego-svr2/service"
	"fmt"
	"beego-svr2/util"
)


// Operations about Users
type BusinessController struct {
	beego.Controller
}


func (business *BusinessController) AddOrderInfo() {
	var orderInfoVo models.OrderInfoVo
	json.Unmarshal(business.Ctx.Input.RequestBody, &orderInfoVo)
	common.ConsoleLogs.Info("orderInfoVo=", orderInfoVo)
	business.Data["json"] = orderInfoVo
	business.ServeJSON()
	service.AddOrderInfo(orderInfoVo)
}




func (this *BusinessController) FindById() interface{}{
	//id, err := business.GetInt("id")
	id, err:= this.GetInt("id")
	common.ConsoleLogs.Info("err=", err)
	if err != nil {
			panic(err)
		}
	common.ConsoleLogs.Info("id=", id)
    	//pid, err := strconv.Atoi(id)
	//if err != nil {
	//	panic("AAA")
	//}
	orderInfo := service.FindById(id)
	json, err := json.Marshal(orderInfo)
	common.ConsoleLogs.Info("json=", json)
	if err != nil {
	    fmt.Println("error:", err)
	}
	this.Ctx.WriteString(string(json))
	//return orderInfo
	//business.Ctx.WriteString(json)
	return json
}