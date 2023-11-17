package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"gopkg.in/resty.v1"
	"time"
)

type (
	CampaignRes struct {
		Code int `json:"code"`
		Data struct {
			DataList []*CampaignInfo `json:"data_list"`
			Total    int             `json:"total"`
		} `json:"data"`
		Msg string `json:"msg"`
	}
	CampaignInfo struct {
		CampaignId   int     `json:"campaign_id"`
		CampaignName string  `json:"campaign_name"`
		PromoteGoal  string  `json:"promote_goal"`
		BudgetType   string  `json:"budget_type"`
		Budget       float32 `json:"budget"`

		AccountID  string `json:"account_id"`
		CreateTime int64  `json:"create_time"`
		UpdateTime int64  `json:"update_time"`
		Version    int64  `json:"version"`
	}
)

func main() {
	var (
		campaignInfoUrl = "https://scaler.taurusx.com/openapi/campaign_info"
		accessKey       = "a7d8dece052147f598c6b57fb2fa5bcb"
		apiKey          = "e1a476536eac4e60b727b570c7140be5"

		page     = 1
		pageSize = 5
	)
	token := generateHash4(apiKey)

	url := fmt.Sprintf("%s?page=%d&page_size=%d",
		campaignInfoUrl, page, pageSize)
	fmt.Println("url:", url)

	resp, err := resty.New().SetRetryCount(3).R().
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
			"access-key":   accessKey,
			"token":        token,
			"timestamp":    fmt.Sprintf("%d", time.Now().Unix()),
		}).
		Get(url)
	if err != nil {
		fmt.Println("Get err", err)
		return
	}
	fmt.Println("resp:", string(resp.Body()))

	var campaignRes CampaignRes
	err = json.Unmarshal(resp.Body(), &campaignRes)
	if err != nil {
		fmt.Println("Unmarshal err", err)
		return
	}
	fmt.Printf("campaignRes:%+v\n", campaignRes)
}

func generateHash4(apiKey string) string {
	// 获取当前时间戳
	timestamp := time.Now().Unix()

	// 对时间戳进行MD5哈希
	timestampHashInBytes := md5.Sum([]byte(fmt.Sprintf("%d", timestamp)))
	timestampHash := hex.EncodeToString(timestampHashInBytes[:])

	// 将API密钥和MD5哈希后的时间戳连接，并计算MD5哈希值
	data := fmt.Sprintf("%s%s", apiKey, timestampHash)
	hashInBytes := md5.Sum([]byte(data))
	hash := hex.EncodeToString(hashInBytes[:])

	return hash
}
