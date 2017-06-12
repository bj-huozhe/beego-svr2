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

