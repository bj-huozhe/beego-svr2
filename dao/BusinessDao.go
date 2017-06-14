package dao

import (
	"github.com/astaxie/beego/orm"
	"beego-svr2/models"
	"fmt"
	"beego-svr2/util"
)

func Add(orderInfo models.OrderInfo) {
	orderInfoDao := orm.NewOrm()
	orderInfoDao.Using("trade")
	err := orderInfoDao.Begin()
	affect, err := orderInfoDao.Insert(&orderInfo)
	fmt.Println("affect=", affect, "err=", err)
	if err != nil {
		orderInfoDao.Rollback()
	} else {
		orderInfoDao.Commit()
	}
}


func FindById(id int) interface{}  {
	common.ConsoleLogs.Info("OrderInfoDao--id=", id)
	orderInfoDao := orm.NewOrm()
	orderInfoDao.Using("default")
	err := orderInfoDao.Begin()
	orderInfo := models.OrderInfo{Id:id}
	common.ConsoleLogs.Info("orderInfo=", orderInfo)
	err = orderInfoDao.Read(&orderInfo)
	common.ConsoleLogs.Info("orderInfo=", orderInfo)
	common.ConsoleLogs.Info("err=", err)
	if err != nil {
		orderInfoDao.Rollback()
	} else {
		orderInfoDao.Commit()
	}
	return orderInfo
}

