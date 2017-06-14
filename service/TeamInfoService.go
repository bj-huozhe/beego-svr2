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

func FindTeamAll() interface{} {
	teamInfoList := dao.FindTeamAll()
	common.ConsoleLogs.Info(".....teamInfoList=", teamInfoList)
	return teamInfoList
}


func UpdateTeam(teamrInfoVo request.TeamInfoVo){
	teamInfo := models.TeamInfo{Id:teamrInfoVo.Id,TeamName:teamrInfoVo.TeamName,OrderId:teamrInfoVo.OrderId,CreateTime:time.Now(), UpdateTime:time.Now(),Version:teamrInfoVo.Version}
	dao.UpdateTeam(teamInfo)
}







