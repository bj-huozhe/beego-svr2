package test

import (
	"testing"
	_ "beego-svr2/routers"
	"fmt"
	"beego-svr2/util"
	"encoding/json"
)

func Test_AddLendBank(t *testing.T) {
	httpUrl := "http://localhost:8100/lendBank/AddLendBank"
	request := make(map[string]interface{})
	request["Id"] = 1
	request["bankNo"] = "6228480402564890018"
        request["bankName"] = "ICBC"

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

func Test_FindLendBankById(t *testing.T) {
	httpUrl := "http://localhost:8100/lendBank/FindLendBankById?id=1"
	header := make(map[string]string)
	header["Cookie"] = "JSESSION:1245566"
	header["Source"] = "BBC"
	//result := common.HttpPost(request, header, httpUrl)
	result := common.HttpGet(httpUrl)
	fmt.Println("result=", result)
}


func Test_FindLendBankAll(t *testing.T) {
	httpUrl := "http://localhost:8100/lendBank/FindLendBankAll?id=59"
	result := common.HttpGet(httpUrl)
	fmt.Println("result=", result)
}

func Test_UpdateLendBank(t *testing.T) {
	httpUrl := "http://localhost:8100/lendBank/UpdateLendBank"
	request := make(map[string]interface{})
	request["Id"] = 1
	request["bankNo"] = "6228480666622220011a"
        request["bankName"] = "ABC"
	request["Version"] = 5

	json, err := json.Marshal(request)
	if err != nil {
		fmt.Println("error:", err)
	}

	header := make(map[string]string)
	header["Cookie"] = "JSESSION:1245566"
	header["Source"] = "BBC"
	result := common.HttpPostJson(string(json), header, httpUrl)
	fmt.Println("result=", result)
}






















