package dao

import (
	"github.com/astaxie/beego/orm"
	"beego-svr2/models"
	"fmt"
	"beego-svr2/util"
)

func AddLendUser(lendUser models.LendUser) {
	instanceDao := orm.NewOrm()
	instanceDao.Using("trade")
	err := instanceDao.Begin()
	affect, err := instanceDao.Insert(&lendUser)
	fmt.Println("affect=", affect, "err=", err)
	if err != nil {
		instanceDao.Rollback()
	} else {
		instanceDao.Commit()
	}
}


func FindLendUserById(id int) interface{}  {
	common.ConsoleLogs.Info("FindLendUserById--id=", id)
	instanceDao := orm.NewOrm()
	instanceDao.Using("default")
	err := instanceDao.Begin()
	modelInfo := models.LendUser{Id:id}
	common.ConsoleLogs.Info("modelInfo=", modelInfo)
	err = instanceDao.Read(&modelInfo)
	common.ConsoleLogs.Info("orderInfo=", modelInfo)
	common.ConsoleLogs.Info("err=", err)
	if err != nil {
		instanceDao.Rollback()
	} else {
		instanceDao.Commit()
	}
	return modelInfo
}

