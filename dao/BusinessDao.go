package dao

import (
	"github.com/astaxie/beego/orm"
	"beego-svr2/models"
	"fmt"
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


func FindById(id int) models.OrderInfo {
	orderInfoDao := orm.NewOrm()
	orderInfoDao.Using("trade")
	err := orderInfoDao.Begin()
	orderInfo := new(models.OrderInfo)
	orderInfo.Id = id
	err = orderInfoDao.Read(&orderInfo)
	fmt.Println("err=", err)
	if err != nil {
		orderInfoDao.Rollback()
	} else {
		orderInfoDao.Commit()
	}
	return orderInfo
}

