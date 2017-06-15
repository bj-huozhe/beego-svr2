package test

import (
	"testing"
	_ "beego-svr2/routers"
	"fmt"
	"beego-svr2/util"
	"encoding/json"
	"golang-test/com/basic/util/uuid"
)

func Test_AddLendUser(t *testing.T) {
	httpUrl := "http://localhost:8100/lendUser/AddLendUser"
	request := make(map[string]interface{})
	request["Id"] = 1
	request["lendId"] = uuid.GetGuid()
	request["UserName"] = "stevenjohn"
	request["Email"] = "stevenjohn@gmail.com"
	request["Age"] = 21
	request["IdNo"] = "211281199207156038"
	request["UserId"] = 1000

	json, err := json.Marshal(request)
	if err != nil {
		fmt.Println("error:", err)
	}

	header := make(map[string]string)
	header["Cookie"] = "JSESSION:1245566"
	header["Source"] = "BBC"
	//result := common.HttpPost(request, header, httpUrl)
	result := common.HttpPostJson(string(json), header, httpUrl)
	fmt.Println("result=", result)
}

func Test_FindLendUserById(t *testing.T) {
	httpUrl := "http://localhost:8100/lendUser/FindLendUserById?id=3"
	result := common.HttpGet(httpUrl)
	fmt.Println("result=", result)
}


func Test_FindLendUserAll(t *testing.T) {
	httpUrl := "http://localhost:8100/lendUser/FindLendUserAll"
	result := common.HttpGet(httpUrl)
	fmt.Println("result=", result)
}


func Test_UpdateLendUser(t *testing.T) {
	httpUrl := "http://localhost:8100/lendUser/UpdateLendUser"
	request := make(map[string]interface{})
	request["Id"] = 1
	request["lendId"] = uuid.GetGuid()
	request["UserName"] = "stevenjohnw"
	request["Email"] = "stevenjohnw@gmail.com"
	request["Age"] = 211
	request["IdNo"] = "211281199207156038-"
	request["UserId"] = 1007
	request["Version"] = 11

	json, err := json.Marshal(request)
	if err != nil {
		fmt.Println("error:", err)
	}

	header := make(map[string]string)
	header["Cookie"] = "JSESSION:1245566"
	header["Source"] = "BBC"
	//result := common.HttpPost(request, header, httpUrl)
	result := common.HttpPostJson(string(json), header, httpUrl)
	fmt.Println("result=", result)
}




