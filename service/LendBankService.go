package service

import (
	"beego-svr2/models"
	"beego-svr2/dao"
	"time"
	"beego-svr2/util"
	"beego-svr2/pojo/request"
)

func AddLendBank(modelInfoVo request.LendBankVo){
	modelInfo := models.LendBank{BankName:modelInfoVo.BankName,BankNo:modelInfoVo.BankNo,CreateTime:time.Now(), UpdateTime:time.Now(),Version:0}
	dao.AddLendBank(modelInfo)
}

func FindLendBankById(id int) interface{} {
	orderInfo := dao.FindLendBankById(id)
	common.ConsoleLogs.Info(".....orderInfo=", orderInfo)
	return orderInfo
}


func FindLendBankAll() interface{} {
	orderInfoList := dao.FindLendBankAll()
	common.ConsoleLogs.Info(".....orderInfoList=", orderInfoList)
	return orderInfoList
}

func UpdateLendBank(updateInfoVo request.LendBankVo){
	updateInfo := models.LendBank{Id:updateInfoVo.Id,BankName:updateInfoVo.BankName,BankNo:updateInfoVo.BankNo,CreateTime:time.Now(), UpdateTime:time.Now(),Version:updateInfoVo.Version}
	dao.UpdateLendBank(updateInfo)
}









