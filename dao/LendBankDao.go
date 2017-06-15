package dao

import (
	"github.com/astaxie/beego/orm"
	"beego-svr2/models"
	"beego-svr2/util"
	"fmt"
)

func AddLendBank(modelInfo models.LendBank) {
	instanceDao := orm.NewOrm()
	err := instanceDao.Begin()
	affect, err := instanceDao.Insert(&modelInfo)
	common.ConsoleLogs.Info("affect=", affect, "err=", err)
	if err != nil {
		common.ConsoleLogs.Error("err=", err)
		instanceDao.Rollback()
	} else {
		instanceDao.Commit()
	}
}

func FindLendBankById(id int) interface{} {
	common.ConsoleLogs.Info("FindLendBankById--id=", id)
	instanceDao := orm.NewOrm()
	err := instanceDao.Begin()
	modelInfo := models.LendBank{Id:id}
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

//find all users
//func listUsers() {
//    var users []User
//    orm.NewOrm().QueryTable("t_user").All(&users)
//    for _, user := range users {
//        fmt.Println(user.ToString())
//    }
//}

func FindLendBankAll() interface{} {
	var modelList []models.LendBank
	instanceDao := orm.NewOrm()
	err := instanceDao.Begin()
	instanceDao.QueryTable("lend_bank").All(&modelList)
	common.ConsoleLogs.Info("modelList=", modelList)
	if err != nil {
		instanceDao.Rollback()
	} else {
		instanceDao.Commit()
	}
	return modelList
}


//o := orm.NewOrm()
//user := User{Id: 1}
//if o.Read(&user) == nil {
//    user.Name = "MyName"
//    if num, err := o.Update(&user); err == nil {
//        fmt.Println(num)
//    }
//}


func UpdateLendBank(modelInfo models.LendBank)  {
	common.ConsoleLogs.Info("modelInfo=", modelInfo)
	instanceDao := orm.NewOrm()
	err := instanceDao.Begin()
	updateInfo := models.LendBank{Id:modelInfo.Id}
	common.ConsoleLogs.Info("updateInfo=", updateInfo)
	if instanceDao.Read(&updateInfo) == nil {
	    updateInfo.BankName = modelInfo.BankName
	    updateInfo.BankNo = modelInfo.BankNo
	    updateInfo.CreateTime = modelInfo.CreateTime
	    updateInfo.UpdateTime = modelInfo.UpdateTime
	    updateInfo.Version = modelInfo.Version
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