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
	//datas           = []string{"2023-12-01", "2023-12-02", "2023-12-03"}
	datas = []string{"2023-12-02"}
	//timestamp = time.Now().Unix()
)

func main() {
	token = generateHash2(apiKey)
	//time.Sleep(1 * time.Second)
	timestamp := time.Now().Unix()
	println("getPerformanceData时间戳", timestamp)
	for i, data := range datas {
		fmt.Println("i:", i, "data:", data)
		res, err := getPerformanceData(token, data, timestamp)
		if err != nil {
			fmt.Println("getPerformanceData err", err)
			return
		}
		fmt.Printf("len: %d, taurusXRes:%+v\n", len(res),
			res)
	}
}

func generateHash2(apiKey string) string {
	// 获取当前时间戳
	timestamp := time.Now().Unix()
	println("token时间戳", timestamp)

	// 对时间戳进行MD5哈希
	timestampHashInBytes := md5.Sum([]byte(fmt.Sprintf("%d", timestamp)))
	timestampHash := hex.EncodeToString(timestampHashInBytes[:])

	// 将API密钥和MD5哈希后的时间戳连接，并计算MD5哈希值
	data := fmt.Sprintf("%s%s", apiKey, timestampHash)
	hashInBytes := md5.Sum([]byte(data))
	hash := hex.EncodeToString(hashInBytes[:])

	return hash
}

func getPerformanceData(token, date string, timestamp int64) ([]TaurusX, error) {
	url := fmt.Sprintf("%s?timezone=%s&start_time=%s&end_time=%s&dimension_option=%s",
		baseUrl, timezone, date, date, dimensionOption)
	fmt.Println("url:", url)

	resp, err := resty.New().SetRetryCount(3).R().
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
			"access-key":   accessKey,
			"token":        token,
			"timestamp":    fmt.Sprintf("%d", timestamp),
		}).
		Get(url)
	if err != nil {
		fmt.Println("Get err", err)
		return nil, err
	}
	fmt.Println("resp:", string(resp.Body()))

	var taurusXRes TaurusXRes
	err = json.Unmarshal(resp.Body(), &taurusXRes)
	if err != nil {
		fmt.Println("Unmarshal err", err)
		return nil, err
	}
	//fmt.Printf("len: %d, taurusXRes:%+v\n", len(taurusXRes.Data.DataList),
	//	taurusXRes)
	return taurusXRes.Data.DataList, nil
}
