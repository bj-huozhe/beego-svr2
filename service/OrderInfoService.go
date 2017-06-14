package service

import (
	"beego-svr2/models"
	"beego-svr2/dao"
	"time"
	"beego-svr2/util"
	"beego-svr2/pojo/request"
)

func AddOrderInfo(orderInfoVo request.OrderInfoVo){
	orderInfo := models.OrderInfo{Name:orderInfoVo.Name,Age:orderInfoVo.Age,Create_time:time.Now(), Update_time:time.Now(),Version:0}
	dao.AddOrderInfo(orderInfo)
}

func FindOrderById(id int) interface{} {
	orderInfo := dao.FindOrderById(id)
	common.ConsoleLogs.Info(".....orderInfo=", orderInfo)
	return orderInfo
}












