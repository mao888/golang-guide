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
	AdGroupRes struct {
		Code int `json:"code"`
		Data struct {
			DataList []*AdGroupInfo `json:"data_list"`
			Total    int            `json:"total"`
		} `json:"data"`
		Msg string `json:"msg"`
	}
	AdGroupInfo struct {
		CampaignName    string      `json:"campaign_name"`
		CampaignId      int         `json:"campaign_id"`
		AdgroupName     string      `json:"adgroup_name"`
		AdgroupId       int         `json:"adgroup_id"`
		PromoteId       int         `json:"promote_id"`
		PromoteName     string      `json:"promote_name"`
		ClickThroughUrl string      `json:"click_through_url"`
		BudgetTotal     float32     `json:"budget_total"`
		BudgetDaily     float32     `json:"budget_daily"`
		BidType         string      `json:"bid_type"`
		BidValue        interface{} `json:"bid_value"`
		BidCpa          []struct {
			BidValue float64 `json:"bid_value"`
			Event    string  `json:"event"`
		} `json:"bid_cpa"`
		Targeting struct {
			Region       []string      `json:"region"`
			City         []interface{} `json:"city"`
			Language     []interface{} `json:"language"`
			Os           string        `json:"os"`
			OsVersionMin interface{}   `json:"os_version_min"`
			Device       []interface{} `json:"device"`
			Network      string        `json:"network"`
		} `json:"targeting"`

		AccountID  string `json:"account_id"`
		CreateTime int64  `json:"create_time"`
		UpdateTime int64  `json:"update_time"`
		Version    int64  `json:"version"`
	}
)

func main() {
	var (
		adGroupInfoUrl = "https://scaler.taurusx.com/openapi/adgroup_info"
		accessKey      = "a7d8dece052147f598c6b57fb2fa5bcb"
		apiKey         = "e1a476536eac4e60b727b570c7140be5"

		page     = 1
		pageSize = 5
	)
	token := generateHash3(apiKey)

	url := fmt.Sprintf("%s?page=%d&page_size=%d",
		adGroupInfoUrl, page, pageSize)
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

	var adGroupRes AdGroupRes
	err = json.Unmarshal(resp.Body(), &adGroupRes)
	if err != nil {
		fmt.Println("Unmarshal err", err)
		return
	}
	fmt.Printf("adGroupRes:%+v\n", adGroupRes)
}

func generateHash3(apiKey string) string {
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
