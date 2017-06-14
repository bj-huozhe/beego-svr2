package dao

import (
	"github.com/astaxie/beego/orm"
	"beego-svr2/models"
	"fmt"
	"beego-svr2/util"
)

func AddTeamInfo(teamInfo models.TeamInfo) {
	instanceDao := orm.NewOrm()
	instanceDao.Using("default")
	err := instanceDao.Begin()
	affect, err := instanceDao.Insert(&teamInfo)
	fmt.Println("affect=", affect, "err=", err)
	if err != nil {
		instanceDao.Rollback()
	} else {
		instanceDao.Commit()
	}
}


func FindTeamById(id int) interface{}  {
	common.ConsoleLogs.Info("FindTeamById--id=", id)
	instanceDao := orm.NewOrm()
	instanceDao.Using("default")
	err := instanceDao.Begin()
	modelInfo := models.TeamInfo{Id:id}
	common.ConsoleLogs.Info("modelInfo=", modelInfo)
	err = instanceDao.Read(&modelInfo)
	common.ConsoleLogs.Info("modelInfo=", modelInfo)
	common.ConsoleLogs.Info("err=", err)
	if err != nil {
		instanceDao.Rollback()
	} else {
		instanceDao.Commit()
	}
	return modelInfo
}



func FindTeamAll() interface{} {
	var teamInfoList []models.TeamInfo
	instanceDao := orm.NewOrm()
	err := instanceDao.Begin()
	instanceDao.QueryTable("team_info").All(&teamInfoList)
	common.ConsoleLogs.Info("teamInfoList=", teamInfoList)
	if err != nil {
		instanceDao.Rollback()
	} else {
		instanceDao.Commit()
	}
	return teamInfoList
}


func UpdateTeam(teamInfo models.TeamInfo)  {
	common.ConsoleLogs.Info("teamInfo=", teamInfo)
	instanceDao := orm.NewOrm()
	err := instanceDao.Begin()
	updateInfo := models.TeamInfo{Id:teamInfo.Id}
	common.ConsoleLogs.Info("updateInfo=", updateInfo)
	if instanceDao.Read(&updateInfo) == nil {
	    updateInfo.TeamName = teamInfo.TeamName
	    updateInfo.OrderId = teamInfo.OrderId
	    updateInfo.CreateTime = teamInfo.CreateTime
	    updateInfo.UpdateTime = teamInfo.UpdateTime
	    updateInfo.Version = teamInfo.Version
	    if num, err := instanceDao.Update(&updateInfo); err == nil {
		fmt.Println(num)
	    }
	}
	common.ConsoleLogs.Info("updateInfo=", updateInfo)
	if err != nil {
		instanceDao.Rollback()
	} else {
		instanceDao.Commit()
	}
}


