package dao

import (
	"github.com/astaxie/beego/orm"
	"beego-svr2/models"
	"beego-svr2/util"
)

func AddOrderInfo(orderInfo models.OrderInfo) {
	orderInfoDao := orm.NewOrm()
	orderInfoDao.Using("trade")
	err := orderInfoDao.Begin()
	affect, err := orderInfoDao.Insert(&orderInfo)
	common.ConsoleLogs.Info("affect=", affect, "err=", err)
	if err != nil {
		common.ConsoleLogs.Error("err=", err)
		orderInfoDao.Rollback()
	} else {
		orderInfoDao.Commit()
	}
}

func FindOrderById(id int) interface{} {
	common.ConsoleLogs.Info("FindOrderById--id=", id)
	instanceDao := orm.NewOrm()
	instanceDao.Using("default")
	err := instanceDao.Begin()
	modelInfo := models.OrderInfo{Id:id}
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

