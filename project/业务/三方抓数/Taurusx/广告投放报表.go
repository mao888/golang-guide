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
	TaurusXRes struct {
		Code int `json:"code"`
		Data struct {
			DataList []TaurusX `json:"data_list"`
		} `json:"data"`
	}
	TaurusX struct {
		Date        string  `json:"date"`
		CampaignId  int     `json:"campaign_id"`
		AdgroupId   int     `json:"adgroup_id"`
		CreativeId  int     `json:"creative_id"`
		AdType      string  `json:"ad_type"`
		Country     string  `json:"country"`
		Clicks      float64 `json:"clicks"`
		Conversions float64 `json:"conversions"`
		Cost        float64 `json:"cost"`
		Impressions float64 `json:"impressions"`
	}
)

var (
	timestamp = time.Now().Unix()
)

func main() {
	var (
		//timestamp = time.Now().Unix()
		token string
		url   string

		baseUrl   = "https://scaler.taurusx.com/openapi/performance_data"
		accessKey = "a7d8dece052147f598c6b57fb2fa5bcb"
		apiKey    = "e1a476536eac4e60b727b570c7140be5"

		timezone = "+0"
		//startTime       = "2023-11-21"
		//endTime         = "2023-11-30"
		dimensionOption = "Campaign,Adgroup,Creative,AdType,Country"
		datas           = []string{"2023-11-21", "2023-11-22", "2023-11-23", "2023-11-24", "2023-11-25", "2023-11-26", "2023-11-27", "2023-11-28", "2023-11-29", "2023-11-30"}
	)
	token = generateHash2(apiKey)

	for i, data := range datas {
		fmt.Println("i:", i, "data:", data)

		url = fmt.Sprintf("%s?timezone=%s&start_time=%s&end_time=%s&dimension_option=%s",
			baseUrl, timezone, data, data, dimensionOption)
		fmt.Println("url:", url)

		resp, err := resty.New().SetRetryCount(3).R().
			SetHeaders(map[string]string{
				"Content-Type": "application/json",
				"access-key":   accessKey,
				"token":        token,
				"timestamp":    fmt.Sprintf("%d", timestamp),
			}).
			Get(url)
		fmt.Println("resp:", string(resp.Body()))
		if err != nil {
			fmt.Println("Get err", err)
			return
		}

		var taurusXRes TaurusXRes
		err = json.Unmarshal(resp.Body(), &taurusXRes)
		if err != nil {
			fmt.Println("Unmarshal err", err)
			return
		}
		fmt.Printf("len: %d, taurusXRes:%+v\n", len(taurusXRes.Data.DataList),
			taurusXRes)
	}
}

func generateHash2(apiKey string) string {
	// 获取当前时间戳
	//timestamp := time.Now().Unix()

	// 对时间戳进行MD5哈希
	timestampHashInBytes := md5.Sum([]byte(fmt.Sprintf("%d", timestamp)))
	timestampHash := hex.EncodeToString(timestampHashInBytes[:])

	// 将API密钥和MD5哈希后的时间戳连接，并计算MD5哈希值
	data := fmt.Sprintf("%s%s", apiKey, timestampHash)
	hashInBytes := md5.Sum([]byte(data))
	hash := hex.EncodeToString(hashInBytes[:])

	return hash
}
