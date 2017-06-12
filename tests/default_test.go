package test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"runtime"
	"path/filepath"
	_ "beego-svr2/routers"

	"github.com/astaxie/beego"
	. "github.com/smartystreets/goconvey/convey"
	"fmt"
	"beego-svr2/util"
	"encoding/json"
)

func init() {
	_, file, _, _ := runtime.Caller(1)
	apppath, _ := filepath.Abs(filepath.Dir(filepath.Join(file, ".." + string(filepath.Separator))))
	beego.TestBeegoInit(apppath)
}

// TestGet is a sample to run an endpoint test
func TestGet(t *testing.T) {
	r, _ := http.NewRequest("GET", "/v1/object", nil)
	w := httptest.NewRecorder()
	beego.BeeApp.Handlers.ServeHTTP(w, r)

	beego.Trace("testing", "TestGet", "Code[%d]\n%s", w.Code, w.Body.String())

	Convey("Subject: Test Station Endpoint\n", t, func() {
	        Convey("Status Code Should Be 200", func() {
	                So(w.Code, ShouldEqual, 200)
	        })
	        Convey("The Result Should Not Be Empty", func() {
	                So(w.Body.Len(), ShouldBeGreaterThan, 0)
	        })
	})
}


func Test_AddOrderInfo(t *testing.T) {
	httpUrl := "http://localhost:8080/business/AddOrderInfo"
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
