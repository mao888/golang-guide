package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func main() {
	// 准备要发送的数据
	data := map[string]interface{}{
		"#account_id": "testing",
		"#time":       "2024-03-01 11:35:53.648",
		"#type":       "track",
		"#event_name": "refund",
		"properties": map[string]string{
			"goods_id":    "23493774147584",
			"goods_type":  "store",
			"goods_name":  "23493774147584",
			"goods_value": "0.01",
			"order_id":    "GPA.3329-3915-8231-02382",
		},
	}

	// 将数据编码为 JSON 格式
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("JSON 编码失败: ", err)
		return
	}

	appid := "test-sdk-appid"
	client := "0"

	// 创建一个 URL 值
	apiURL := "https://global-receiver-ta.thinkingdata.cn/sync_data"

	// 准备要发送的数据
	formData := url.Values{}
	formData.Set("appid", appid)
	formData.Set("data", string(jsonData))
	formData.Set("client", client)
	formData.Set("debug", "1")

	// 发送请求
	response, err := http.PostForm(apiURL, formData)
	if err != nil {
		fmt.Println("请求发送失败:", err)
		return
	}
	if response.StatusCode != 200 {
		fmt.Println("请求发送失败:", response.StatusCode)
		return
	}
	defer response.Body.Close()

	// 读取响应内容
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("响应读取失败:", err)
		return
	}

	// Response
	var resp Response
	json.Unmarshal(body, &resp)
	if err != nil {
		fmt.Println("json.Unmarshal failed:", err)
		return
	}
	fmt.Println("code:", resp.Code)

	// 输出响应内容
	fmt.Println("响应状态码:", response.Status)
}
