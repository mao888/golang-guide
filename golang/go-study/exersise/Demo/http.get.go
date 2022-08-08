package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type RegBasicRes struct {
	Data struct {
		DnaCount    int64 `json:"dnaCount"`
		SeriesCount int64 `json:"seriesCount"`
	} `json:"data"`
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
}

//	GetReg 请求认证平台
func GetReg() (regData *RegBasicRes, err error) {
	baseUrl := "http://192.168.2.141:30081"
	url := baseUrl + "/api/v1/statistic/perform"
	req, err := http.NewRequest("GET", url, nil) // 发起请求，没有参数时，可将bytes.NewBuffer(jsonStr)改为nil
	if err != nil {
		return nil, err
	}

	client := &http.Client{} // 处理返回结果
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("请求认证平台 code: %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	var basicResJson RegBasicRes           // 定义一个结构体，用来将结果的string格式转成json格式，便于对请求结果进行处理
	body, err := ioutil.ReadAll(resp.Body) // 读取请求结果
	if err != nil {
		return nil, err
	}

	// 请求结果string格式
	errJson := json.Unmarshal(body, &basicResJson) // 将string 格式转成json格式
	if errJson != nil {
		return nil, errJson
	}
	return &basicResJson, err
}

func main() {
	GetReg()
}
