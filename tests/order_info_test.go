package test

import (
	"testing"
	_ "beego-svr2/routers"
	"fmt"
	"beego-svr2/util"
	"encoding/json"
)

func Test_AddOrderInfo(t *testing.T) {
	httpUrl := "http://localhost:8100/order/AddOrderInfo"
	request := make(map[string]interface{})
	request["Id"] = 1
	request["Name"] = "abin1"
	request["Age"] = 23

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

func Test_FindOrderById(t *testing.T) {
	httpUrl := "http://localhost:8100/order/FindOrderById?id=59"
	//httpUrl := "http://localhost:8080/business/FindById/59"
	request := make(map[string]interface{})
	request["Id"] = 1
	request["Name"] = "abin1"
	request["Age"] = 23


	//json, err := json.Marshal(request)
	//if err != nil {
	//    fmt.Println("error:", err)
	//}

	header := make(map[string]string)
	header["Cookie"] = "rules_session_id=64b713c52c7511e6a4519801a7928995"
	header["RRDSource"] = "YM"
	//result := common.HttpPost(request, header, httpUrl)
	result := common.HttpGet(httpUrl)
	fmt.Println("result=", result)
}


func Test_FindOrderAll(t *testing.T) {
	httpUrl := "http://localhost:8100/order/FindOrderAll?id=59"
	result := common.HttpGet(httpUrl)
	fmt.Println("result=", result)
}

func Test_UpdateOrder(t *testing.T) {
	httpUrl := "http://localhost:8100/order/UpdateOrder"
	request := make(map[string]interface{})
	request["Id"] = 64
	request["Name"] = "abin29"
	request["Age"] = 29
	request["Version"] = 5

	json, err := json.Marshal(request)
	if err != nil {
		fmt.Println("error:", err)
	}

	header := make(map[string]string)
	header["Cookie"] = "rules_session_id=64b713c52c7511e6a4519801a7928995"
	header["RRDSource"] = "YM"
	result := common.HttpPostJson(string(json), header, httpUrl)
	fmt.Println("result=", result)
}






















