package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"beego-svr2/models"
	"beego-svr2/service"
	"fmt"
)


// Operations about Users
type BusinessController struct {
	beego.Controller
}


func (u *BusinessController) AddOrderInfo() {
	var orderInfoVo models.OrderInfoVo
	json.Unmarshal(u.Ctx.Input.RequestBody, &orderInfoVo)
	fmt.Printf("orderInfoVo=", orderInfoVo)
	u.Data["json"] = orderInfoVo
	u.ServeJSON()
	service.AddOrderInfo(orderInfoVo)
}

