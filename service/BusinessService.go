package service

import (
	"beego-svr2/models"
	"beego-svr2/dao"
	"time"
)

func AddOrderInfo(orderInfoVo models.OrderInfoVo){
	orderInfo := models.OrderInfo{Name:orderInfoVo.Name,Age:orderInfoVo.Age,Create_time:time.Now(), Update_time:time.Now(),Version:0}
	dao.Add(orderInfo)
}

func FindById(id int) interface{} {
	orderInfo := dao.FindById(id)
	return orderInfo
}












