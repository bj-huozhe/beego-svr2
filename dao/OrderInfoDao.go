package dao

import (
	"github.com/astaxie/beego/orm"
	"beego-svr2/models"
	"beego-svr2/util"
	"fmt"
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

//find all users
//func listUsers() {
//    var users []User
//    orm.NewOrm().QueryTable("t_user").All(&users)
//    for _, user := range users {
//        fmt.Println(user.ToString())
//    }
//}

func FindOrderAll() interface{} {
	var orderList []models.OrderInfo
	instanceDao := orm.NewOrm()
	err := instanceDao.Begin()
	instanceDao.QueryTable("order_info").All(&orderList)
	common.ConsoleLogs.Info("orderList=", orderList)
	if err != nil {
		instanceDao.Rollback()
	} else {
		instanceDao.Commit()
	}
	return orderList
}


//o := orm.NewOrm()
//user := User{Id: 1}
//if o.Read(&user) == nil {
//    user.Name = "MyName"
//    if num, err := o.Update(&user); err == nil {
//        fmt.Println(num)
//    }
//}


func UpdateOrder(orderInfo models.OrderInfo)  {
	common.ConsoleLogs.Info("orderInfo=", orderInfo)
	instanceDao := orm.NewOrm()
	err := instanceDao.Begin()
	updateInfo := models.OrderInfo{Id:orderInfo.Id}
	common.ConsoleLogs.Info("updateInfo=", updateInfo)
	if instanceDao.Read(&updateInfo) == nil {
	    updateInfo.Name = orderInfo.Name
	    updateInfo.Age = orderInfo.Age
	    updateInfo.CreateTime = orderInfo.CreateTime
	    updateInfo.UpdateTime = orderInfo.UpdateTime
	    updateInfo.Version = orderInfo.Version
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