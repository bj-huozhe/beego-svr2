package request


type LendUserVo struct {
	Id int  `primary key Id`
	LendId string  `LendId`
	UserName string  `UserName`
	Email string  `Email`
	Age int  `Age`
	IdNo string  `IdNo`
	UserId int `UserId`
}

