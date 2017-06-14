package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"beego-svr2/models"
	"beego-svr2/service"
	"fmt"
	"strconv"
)


// Operations about Users
type BusinessController struct {
	beego.Controller
}


func (business *BusinessController) AddOrderInfo() {
	var orderInfoVo models.OrderInfoVo
	json.Unmarshal(business.Ctx.Input.RequestBody, &orderInfoVo)
	fmt.Printf("orderInfoVo=", orderInfoVo)
	business.Data["json"] = orderInfoVo
	business.ServeJSON()
	service.AddOrderInfo(orderInfoVo)
}




func (business *BusinessController) FindById()  {
	//id, err := business.GetInt("id")
	id, err:= strconv.Atoi(business.Ctx.Input.Param(":id"))
	fmt.Printf("err=", err)
	if err != nil {
			panic(err)
		}
	fmt.Printf("id=", id)
    	//pid, err := strconv.Atoi(id)
	//if err != nil {
	//	panic("AAA")
	//}
	orderInfo := service.FindById(id)
	json, err := json.Marshal(orderInfo)
	if err != nil {
	    fmt.Println("error:", err)
	}
	business.Ctx.WriteString(string(json))
	//return orderInfo
}