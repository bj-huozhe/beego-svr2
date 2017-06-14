package service

import (
	"beego-svr2/models"
	"beego-svr2/dao"
	"time"
	"beego-svr2/util"
	"beego-svr2/pojo/request"
)

func AddLendInfo(lendInfoVo request.LendInfoVo){
	lendInfo := models.LendInfo{LendId:lendInfoVo.LendId,RequestContent:lendInfoVo.RequestContent,ResponseContent:lendInfoVo.ResponseContent,Status:lendInfoVo.Status,UserId:lendInfoVo.UserId,CreateTime:time.Now(), UpdateTime:time.Now(),Version:0}
	dao.AddLendInfo(lendInfo)
}

func FindLendInfoById(id int) interface{} {
	lendInfo := dao.FindLendInfoById(id)
	common.ConsoleLogs.Info(".....lendInfo=", lendInfo)
	return lendInfo
}


func FindLendInfoAll() interface{} {
	lendInfoList := dao.FindLendInfoAll()
	common.ConsoleLogs.Info(".....lendInfoList=", lendInfoList)
	return lendInfoList
}


func UpdateLendInfo(lendInfoVo request.LendInfoVo){
	lendInfo := models.LendInfo{Id:lendInfoVo.Id,LendId:lendInfoVo.LendId,RequestContent:lendInfoVo.RequestContent,ResponseContent:lendInfoVo.ResponseContent,Status:lendInfoVo.Status,UserId:lendInfoVo.UserId,CreateTime:time.Now(), UpdateTime:time.Now(),Version:0}
	dao.UpdateLendInfo(lendInfo)
}







