package dao

import (
	"github.com/astaxie/beego/orm"
	"beego-svr2/models"
	"beego-svr2/util"
	"fmt"
)

func AddLendInfo(modelInfo models.LendInfo) {
	instanceDao := orm.NewOrm()
	instanceDao.Using("trade")
	err := instanceDao.Begin()
	affect, err := instanceDao.Insert(&modelInfo)
	common.ConsoleLogs.Info("affect=", affect, "err=", err)
	if err != nil {
		instanceDao.Rollback()
	} else {
		instanceDao.Commit()
	}
}


func FindLendInfoById(id int) interface{}  {
	common.ConsoleLogs.Info("OrderInfoDao--id=", id)
	instanceDao := orm.NewOrm()
	instanceDao.Using("default")
	err := instanceDao.Begin()
	modelInfo := models.LendInfo{Id:id}
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


func FindLendInfoAll() interface{} {
	var lendInfoList []models.LendInfo
	instanceDao := orm.NewOrm()
	err := instanceDao.Begin()
	instanceDao.QueryTable("lend_info").All(&lendInfoList)
	common.ConsoleLogs.Info("lendInfoList=", lendInfoList)
	if err != nil {
		instanceDao.Rollback()
	} else {
		instanceDao.Commit()
	}
	return lendInfoList
}


func UpdateLendInfo(lendInfo models.LendInfo)  {
	common.ConsoleLogs.Info("lendInfo=", lendInfo)
	instanceDao := orm.NewOrm()
	err := instanceDao.Begin()
	updateInfo := models.LendInfo{Id:lendInfo.Id}
	common.ConsoleLogs.Info("updateInfo=", updateInfo)
	if instanceDao.Read(&updateInfo) == nil {
	    updateInfo.LendId = lendInfo.LendId
	    updateInfo.RequestContent = lendInfo.RequestContent
	    updateInfo.ResponseContent = lendInfo.ResponseContent
	    updateInfo.Status = lendInfo.Status
	    updateInfo.UserId = lendInfo.UserId
	    updateInfo.UpdateTime = lendInfo.UpdateTime
	    updateInfo.Version = lendInfo.Version
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





