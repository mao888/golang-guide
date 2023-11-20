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
	CreativeRes struct {
		Code int `json:"code"`
		Data struct {
			DataList []*CreativeInfo `json:"data_list"`
			Total    int             `json:"total"`
		} `json:"data"`
		Msg string `json:"msg"`
	}
	CreativeInfo struct {
		CampaignId   int    `json:"campaign_id"`
		CampaignName string `json:"campaign_name"`
		AdgroupId    int    `json:"adgroup_id"`
		AdgroupName  string `json:"adgroup_name"`
		CreativeId   int    `json:"creative_id"`
		CreativeName string `json:"creative_name"`
		AdType       string `json:"ad_type"`
		Info         struct {
			Type     string  `json:"type"`
			BitRate  float64 `json:"bit_rate"`
			Duration float64 `json:"duration"`
			W        int     `json:"w"`
			H        int     `json:"h"`
		} `json:"info"`
		Companion struct {
			Type string `json:"type"`
			W    int    `json:"w"`
			H    int    `json:"h"`
		} `json:"companion"`

		AccountID  string `json:"account_id"`
		CreateTime int64  `json:"create_time"`
		UpdateTime int64  `json:"update_time"`
		Version    int64  `json:"version"`
	}
)

func main() {
	var (
		creativeInfoUrl = "https://scaler.taurusx.com/openapi/creative_info"
		accessKey       = "a7d8dece052147f598c6b57fb2fa5bcb"
		apiKey          = "e1a476536eac4e60b727b570c7140be5"

		page     = 1
		pageSize = 50
	)
	token := generateHash5(apiKey)

	url := fmt.Sprintf("%s?page=%d&page_size=%d",
		creativeInfoUrl, page, pageSize)
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

	var campaignRes CreativeRes
	err = json.Unmarshal(resp.Body(), &campaignRes)
	if err != nil {
		fmt.Println("Unmarshal err", err)
		return
	}
	fmt.Printf("campaignRes:%+v\n", campaignRes)
}

func generateHash5(apiKey string) string {
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
