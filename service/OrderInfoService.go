package service

import (
	"beego-svr2/models"
	"beego-svr2/dao"
	"time"
	"beego-svr2/util"
	"beego-svr2/pojo/request"
)

func AddOrderInfo(orderInfoVo request.OrderInfoVo){
	orderInfo := models.OrderInfo{Name:orderInfoVo.Name,Age:orderInfoVo.Age,CreateTime:time.Now(), UpdateTime:time.Now(),Version:0}
	dao.AddOrderInfo(orderInfo)
}

func FindOrderById(id int) interface{} {
	orderInfo := dao.FindOrderById(id)
	common.ConsoleLogs.Info(".....orderInfo=", orderInfo)
	return orderInfo
}


func FindOrderAll() interface{} {
	orderInfoList := dao.FindOrderAll()
	common.ConsoleLogs.Info(".....orderInfoList=", orderInfoList)
	return orderInfoList
}

func UpdateOrder(orderInfoVo request.OrderInfoVo){
	orderInfo := models.OrderInfo{Id:orderInfoVo.Id,Name:orderInfoVo.Name,Age:orderInfoVo.Age,CreateTime:time.Now(), UpdateTime:time.Now(),Version:orderInfoVo.Version}
	dao.UpdateOrder(orderInfo)
}









