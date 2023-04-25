package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/resty.v1"
)

func main() {
	type tokenResp struct {
		Code              int    `json:"code"`
		Msg               string `json:"msg"`
		TenantAccessToken string `json:"tenant_access_token"`
		Expire            int    `json:"expire"`
	}

	var (
		// 请求头
		contentType = "application/json; charset=utf-8"
		// 请求体
		appId     = "cli_a365e3b762b8900d"
		appSecret = "dvmL03uDtcb7DXKDFRRxPcuxSPAX7zmg"
		// HTTP URL
		url = "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal"
		// tokenResp
		tenantTokenResp tokenResp
	)
	resp, err := resty.New().SetRetryCount(3).R().SetHeaders(map[string]string{
		"Content-Type": contentType,
	}).SetFormData(map[string]string{
		"app_id":     appId,
		"app_secret": appSecret,
	}).Post(url)
	if err != nil {
		fmt.Println("Post err", err)
		return
	}
	err = json.Unmarshal(resp.Body(), &tenantTokenResp)
	if err != nil {
		fmt.Println()
		return
	}
	fmt.Println("tenant_access_token: ", tenantTokenResp)
}
