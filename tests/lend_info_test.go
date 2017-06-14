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

	json, err := json.Marshal(request)
	if err != nil {
		fmt.Println("error:", err)
	}

	header := make(map[string]string)
	header["Cookie"] = "rules_session_id=64b713c52c7511e6a4519801a7928995"
	header["RRDSource"] = "YM"
	//result := common.HttpPost(request, header, httpUrl)
	result := common.HttpPostJson(string(json), header, httpUrl)
	fmt.Println("result=", result)
}

func Test_FindLendInfoById(t *testing.T) {
	httpUrl := "http://localhost:8100/lend/FindLendInfoById?id=4"
	result := common.HttpGet(httpUrl)
	fmt.Println("result=", result)
}
