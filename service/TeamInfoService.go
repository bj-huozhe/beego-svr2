package service

import (
	"beego-svr2/models"
	"beego-svr2/dao"
	"time"
	"beego-svr2/util"
	"beego-svr2/pojo/request"
)

func AddTeamInfo(teamInfoVo request.TeamInfoVo){
	teamInfo := models.TeamInfo{TeamName:teamInfoVo.TeamName,OrderId:teamInfoVo.OrderId,CreateTime:time.Now(), UpdateTime:time.Now(),Version:0}
	dao.AddTeamInfo(teamInfo)
}

func FindTeamById(id int) interface{} {
	teamInfo := dao.FindTeamById(id)
	common.ConsoleLogs.Info(".....teamInfo=", teamInfo)
	return teamInfo
}












