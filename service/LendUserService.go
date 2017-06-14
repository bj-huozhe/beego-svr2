package service

import (
	"beego-svr2/models"
	"beego-svr2/dao"
	"time"
	"beego-svr2/util"
	"beego-svr2/pojo/request"
)

func AddLendUser(lendUserVo request.LendUserVo){
	orderInfo := models.LendUser{LendId:lendUserVo.LendId,UserName:lendUserVo.UserName,Email:lendUserVo.Email,Age:lendUserVo.Age,IdNo:lendUserVo.IdNo,UserId:lendUserVo.UserId,CreateTime:time.Now(), UpdateTime:time.Now(),Version:0}
	dao.AddLendUser(orderInfo)
}

func FindLendUserById(id int) interface{} {
	orderInfo := dao.FindLendUserById(id)
	common.ConsoleLogs.Info(".....orderInfo=", orderInfo)
	return orderInfo
}












