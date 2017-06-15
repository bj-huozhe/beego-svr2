package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type OrderInfo struct {
	Id int  `primary key Id`
	Name string  `Name`
	Age int `Age`
	CreateTime time.Time `create_time`
	UpdateTime time.Time `update_time`
	Version int  `version `
}

type TeamInfo struct {
	Id int  `primary key Id`
	TeamName string  `team_name`
	OrderId int `order_id`
	CreateTime time.Time `create_time`
	UpdateTime time.Time `update_time`
	Version int  `version `
}

type LendInfo struct {
	Id int  `primary key Id`
	LendId string  `lend_id`
	RequestContent string  `request_content`
	ResponseContent string  `response_content`
	Status string  `status`
	UserId int `user_id`
	CreateTime time.Time `create_time`
	UpdateTime time.Time `update_time`
	Version int  `version `
}

type LendUser struct {
	Id int  `primary key Id`
	LendId string  `LendId`
	UserName string  `user_name`
	Email string  `email`
	Age int `age`
	IdNo string  `id_no`
	UserId int `user_id`
	CreateTime time.Time `create_time`
	UpdateTime time.Time `update_time`
	Version int  `version `
}

type LendBank struct {
	Id int  `primary key Id`
	BankNo string  `bank_no`
	BankName string  `bank_name`
	CreateTime time.Time `create_time`
	UpdateTime time.Time `update_time`
	Version int  `version `
}



func RegisterDB() {
    //注册 model
    orm.RegisterModel(new(OrderInfo),new(TeamInfo), new(LendInfo), new(LendUser), new(LendBank))
    //注册驱动
    orm.RegisterDriver("mysql", orm.DRMySQL)
    //注册默认数据库
    orm.RegisterDataBase("default", "mysql", "root:root@tcp(172.16.2.133:3306)/trade?charset=utf8")//密码为空格式

}