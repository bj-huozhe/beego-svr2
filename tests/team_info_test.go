package test

import (
	"testing"
	_ "beego-svr2/routers"
	"fmt"
	"beego-svr2/util"
	"encoding/json"
	"golang-test/com/basic/util/uuid"
)

func Test_AddTeamInfo(t *testing.T) {
	httpUrl := "http://localhost:8100/team/AddTeamInfo"
	request := make(map[string]interface{})
	request["Id"] = 1
	request["lendId"] = uuid.GetGuid()
	request["TeamName"] = "BAR"
	request["OrderId"] = 63

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

func Test_FindTeamById(t *testing.T) {
	httpUrl := "http://localhost:8100/team/FindTeamById?id=2"
	header := make(map[string]string)
	header["Cookie"] = "JSESSION:1245566"
	header["Source"] = "BBC"
	result := common.HttpGet(httpUrl)
	fmt.Println("result=", result)
}


func Test_FindTeamAll(t *testing.T) {
	httpUrl := "http://localhost:8100/team/FindTeamAll"
	result := common.HttpGet(httpUrl)
	fmt.Println("result=", result)
}


func Test_UpdateTeam(t *testing.T) {
	httpUrl := "http://localhost:8100/team/UpdateTeam"
	request := make(map[string]interface{})
	request["Id"] = 2
	request["lendId"] = uuid.GetGuid()
	request["TeamName"] = "COFFEE"
	request["OrderId"] = 48
	request["Version"] = 3

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












