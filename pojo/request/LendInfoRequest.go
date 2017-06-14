package request


type LendInfoVo struct {
	Id int  `primary key Id`
	LendId string  `LendId`
	RequestContent string  `RequestContent`
	ResponseContent string  `ResponseContent`
	Status string  `Status`
	UserId int `UserId`
	Version int `Version`
}

