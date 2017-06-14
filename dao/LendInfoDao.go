package dao

import (
	"github.com/astaxie/beego/orm"
	"beego-svr2/models"
	"beego-svr2/util"
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

