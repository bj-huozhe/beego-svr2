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



func FindLendUserAll() interface{} {
	var lendUserList []models.LendUser
	instanceDao := orm.NewOrm()
	err := instanceDao.Begin()
	instanceDao.QueryTable("lend_user").All(&lendUserList)
	common.ConsoleLogs.Info("lendUserList=", lendUserList)
	if err != nil {
		instanceDao.Rollback()
	} else {
		instanceDao.Commit()
	}
	return lendUserList
}


func UpdateLendUser(lendUser models.LendUser)  {
	common.ConsoleLogs.Info("lendUser=", lendUser)
	instanceDao := orm.NewOrm()
	err := instanceDao.Begin()
	updateInfo := models.LendUser{Id:lendUser.Id}
	common.ConsoleLogs.Info("updateInfo=", updateInfo)
	if instanceDao.Read(&updateInfo) == nil {
	    updateInfo.LendId = lendUser.LendId
	    updateInfo.UserName = lendUser.UserName
	    updateInfo.Email = lendUser.Email
	    updateInfo.Age = lendUser.Age
	    updateInfo.IdNo = lendUser.IdNo
	    updateInfo.UserId = lendUser.UserId
	    updateInfo.UpdateTime = lendUser.UpdateTime
	    updateInfo.Version = lendUser.Version
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

