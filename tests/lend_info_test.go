package test

import (
	"testing"
	_ "beego-svr2/routers"
	"fmt"
	"beego-svr2/util"
	"encoding/json"
	"golang-test/com/basic/util/uuid"
)

func Test_AddLendInfo(t *testing.T) {
	httpUrl := "http://localhost:8100/lend/AddLendInfo"
	request := make(map[string]interface{})
	request["Id"] = 1
	request["LendId"] = uuid.GetGuid()
	request["RequestContent"] = "{'name':'abin'}"
	request["ResponseContent"] = "{'status':'SUCCESS'}"
	request["UserId"] = 1000
	request["Status"] = "INIT"

	json, err := json.Marshal(request)
	if err != nil {
		fmt.Println("error:", err)
	}

	header := make(map[string]string)
	header["Cookie"] = "JSESSION:1245566"
	header["RRDSource"] = "BBC"
	//result := common.HttpPost(request, header, httpUrl)
	result := common.HttpPostJson(string(json), header, httpUrl)
	fmt.Println("result=", result)
}

func Test_FindLendInfoById(t *testing.T) {
	httpUrl := "http://localhost:8100/lend/FindLendInfoById?id=4"
	result := common.HttpGet(httpUrl)
	fmt.Println("result=", result)
}


func Test_FindLendInfoAll(t *testing.T) {
	httpUrl := "http://localhost:8100/lend/FindLendInfoAll"
	result := common.HttpGet(httpUrl)
	fmt.Println("result=", result)
}


func Test_UpdateLendInfo(t *testing.T) {
	httpUrl := "http://localhost:8100/lend/UpdateLendInfo"
	request := make(map[string]interface{})
	request["Id"] = 1
	request["LendId"] = uuid.GetGuid()
	request["RequestContent"] = "{'name':'abin1'}"
	request["ResponseContent"] = "{'status':'FAILURE'}"
	request["UserId"] = 1009
	request["Status"] = "SUCCESS"
	request["Version"] = 5

	json, err := json.Marshal(request)
	if err != nil {
		fmt.Println("error:", err)
	}

	header := make(map[string]string)
	header["Cookie"] = "JSESSION:1245566"
	header["RRDSource"] = "BBC"
	//result := common.HttpPost(request, header, httpUrl)
	result := common.HttpPostJson(string(json), header, httpUrl)
	fmt.Println("result=", result)
}




