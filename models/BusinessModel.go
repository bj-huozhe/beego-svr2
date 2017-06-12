package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type OrderInfo struct {
	Id int32  `primary key Id`
	Name string  `Name`
	Age int32 `Age`
	Create_time time.Time
	Update_time time.Time
	Version int32  `version `
}

type OrderInfoVo struct {
	Id int32  `primary key Id`
	Name string  `Name`
	Age int32 `Age`
}




func RegisterDB() {
    //注册 model
    orm.RegisterModel(new(OrderInfo))
    //注册驱动
    orm.RegisterDriver("mysql", orm.DRMySQL)
    //注册默认数据库
    orm.RegisterDataBase("default", "mysql", "root:root@tcp(172.16.2.133:3306)/trade?charset=utf8")//密码为空格式

}