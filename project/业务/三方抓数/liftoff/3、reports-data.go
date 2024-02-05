package main

import (
	"context"
	"encoding/base64"
	"encoding/csv"
	"fmt"
	glog "github.com/mao888/mao-glog"
	"gopkg.in/resty.v1"
	"strings"
)

type (
	LiftoffReportDataRes struct {
		Date                    string `json:"date"`
		AppId                   string `json:"app_id"`
		CampaignId              string `json:"campaign_id"`
		PublisherAppStoreId     string `json:"publisher_app_store_id"`
		CreativeId              string `json:"creative_id"`
		CountryCode             string `json:"country_code"`
		PublisherName           string `json:"publisher_name"`
		Spend                   string `json:"spend"`
		Impressions             string `json:"impressions"`
		Clicks                  string `json:"clicks"`
		Installs                string `json:"installs"`
		Cpm                     string `json:"cpm"`
		Cpc                     string `json:"cpc"`
		Ctr                     string `json:"ctr"`
		Cpi                     string `json:"cpi"`
		Cpa                     string `json:"cpa"`
		AdsFirst24hEnterLevel24 string `json:"ads_first24h_enterLevel24"`
		AfPurchase              string `json:"af_purchase"`
		AdsFirst14dPurchase     string `json:"ads_first14d_purchase"`
		AdsInstall24Login       string `json:"ads_install24_login"`

		CampaignName string `json:"campaign_name"`
		CreativeName string `json:"creative_name"`
		AccountName  string `json:"account_name"`
	}
)

func main() {
	ctx := context.Background()
	var (
		reportId        = "1d728ae837"
		apiKey          = "bacfa09c4f"
		apiSecret       = "U1NUhwT2c1s0GRPka9DmZg=="
		basicLiftoffUrl = "https://data.liftoff.io/api/v1/reports"
	)
	// 拼接client_id和secret，并转换为字节数组
	data := []byte(apiKey + ":" + apiSecret)
	// 使用base64进行编码
	encoded := base64.StdEncoding.EncodeToString(data)
	authorization := "Basic " + encoded

	requestUrl := fmt.Sprintf("%s/%s/data", basicLiftoffUrl, reportId)
	resp, err := resty.New().SetRetryCount(3).R().
		SetHeaders(map[string]string{
			"Authorization": authorization,
		}).Get(requestUrl)
	if err != nil {
		glog.Infof(ctx, "Post err:%s", err)
		return
	}

	// 解析CSV文件
	liftoffReportDataResList, err := parseCSVAll(resp.Body())
	if err != nil {
		glog.Errorf(ctx, "parseCSVAll err:%s", err)
		return
	}

	glog.Infof(ctx, "liftoffReportDataResList:%d", len(liftoffReportDataResList))
}

// parseCSVAll
func parseCSVAll(csvData []byte) ([]LiftoffReportDataRes, error) {
	// 读取CSV文件中的所有记录
	reader := csv.NewReader(strings.NewReader(string(csvData)))
	reader.Comma = ','

	// 读取CSV文件中的所有记录
	records, err := reader.ReadAll()
	if err != nil {
		glog.Errorf(context.Background(), "error reading CSV: %v", err)
		return nil, err
	}

	// 创建一个存储 LiftoffReportDataRes 的切片
	var liftoffReportDataResList []LiftoffReportDataRes

	// 遍历CSV记录并将其映射到 LiftoffReportDataRes 结构体
	for i, record := range records {
		// 跳过标题行
		if i == 0 {
			continue
		}
		liftoffReportDataRes := LiftoffReportDataRes{
			Date:                    record[0],
			AppId:                   record[1],
			CampaignId:              record[2],
			PublisherAppStoreId:     record[3],
			CreativeId:              record[4],
			CountryCode:             record[5],
			PublisherName:           record[6],
			Spend:                   record[7],
			Impressions:             record[8],
			Clicks:                  record[9],
			Installs:                record[10],
			Cpm:                     record[11],
			Cpc:                     record[12],
			Ctr:                     record[13],
			Cpi:                     record[14],
			Cpa:                     record[15],
			AdsFirst24hEnterLevel24: record[16],
			AfPurchase:              record[17],
			AdsFirst14dPurchase:     record[18],
			AdsInstall24Login:       record[19],
		}
		liftoffReportDataResList = append(liftoffReportDataResList, liftoffReportDataRes)
	}
	glog.Infof(context.Background(), "总共:%d条; 实际:%d", len(records)-1, len(liftoffReportDataResList)) // 总共:6993条; 实际:6993
	return liftoffReportDataResList, nil
}
