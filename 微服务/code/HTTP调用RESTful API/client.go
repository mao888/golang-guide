// client/selectsort.go

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//type addParam struct {
//	X int `json:"x"`
//	Y int `json:"y"`
//}
//
//type addResult struct {
//	Code int `json:"code"`
//	Data int `json:"data"`
//}

func main() {
	// 通过HTTP请求调用其他服务器上的add服务
	url := "http://127.0.0.1:9090/add"
	param := addParam{
		X: 10,
		Y: 20,
	}
	paramBytes, _ := json.Marshal(param)
	resp, _ := http.Post(url, "application/json", bytes.NewReader(paramBytes))
	defer resp.Body.Close()

	respBytes, _ := ioutil.ReadAll(resp.Body)
	var respData addResult
	json.Unmarshal(respBytes, &respData)
	fmt.Println(respData.Data) // 30
}
