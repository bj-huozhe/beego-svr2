package request


type LendBankVo struct {
	Id int  `primary key Id`
	BankName string  `BankName`
	BankNo string `BankNo`
	Version int `Version`
}

