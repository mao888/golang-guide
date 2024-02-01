package main

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	glog "github.com/mao888/mao-glog"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"gopkg.in/resty.v1"
	"io"
)

type (
	JobDataRes struct {
		Date             string `json:"date"`
		ToCampaignName   string `json:"to_campaign_name"`
		ToCampaignId     string `json:"to_campaign_id"`
		ToAppName        string `json:"to_app_name"`
		ToAppID          string `json:"to_app_id"`
		ToAppBundle      string `json:"to_app_bundle"`
		ToAppPlatform    string `json:"to_app_platform"`
		FromCampaignId   string `json:"from_campaign_id"`
		FromCampaignName string `json:"from_campaign_name"`
		FromAppName      string `json:"from_app_name"`
		FromAppID        string `json:"from_app_id"`
		FromAppBundle    string `json:"from_app_bundle"`
		FromAppPlatform  string `json:"from_app_platform"`
		CampaignType     string `json:"campaign_type"`
		Creative         string `json:"creative"`
		Role             string `json:"role"`
		AdType           string `json:"ad_type"`
		Impressions      string `json:"impressions"`
		Clicks           string `json:"clicks"`
		Installs         string `json:"installs"`
		CTR              string `json:"ctr"`
		IR               string `json:"ir"`
		MoneyEarned      string `json:"money_earned"`
		ECPMEarned       string `json:"ecpm_earned"`
		MoneySpent       string `json:"money_spent"`
		ECPMSpent        string `json:"ecpm_spent"`
		CompletedView    string `json:"completed_view"`
	}
)

func main() {
	var (
		ctx               = context.Background()
		chartboostJobsUrl = "https://analytics.chartboost.com/v3/metrics/jobs/"
		//jobId             = "0391cd5465d59e20f7e2d4fa6028c513/2eeb5b6e-351e-42f2-8984-059623824a90"
		jobId = "dd73c0113e1ffde3a52bd4c46b2fb882/eb9c3504-27af-43a3-8597-e5863a951599"
	)
	// https://analytics.chartboost.com/v3/metrics/jobs/947231990b73fe59bc6150cd47333f6c/b7b1ca22-0977-43f9-9096-81e79d0b0e8e
	requestUrl := fmt.Sprintf("%s%s", chartboostJobsUrl, jobId)
	resp, err := resty.New().SetRetryCount(3).R().Get(requestUrl)
	if err != nil {
		glog.Errorf(ctx, "Post err:%s", err)
		return
	}
	// 读取响应体并将 UTF-16LE 转码为 UTF-8
	bodyReader := bytes.NewReader(resp.Body())
	utf16leDecoder := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewDecoder()
	utf8Reader := transform.NewReader(bodyReader, utf16leDecoder)

	// 读取resp.Body()中的数据,为csv文件,然后解析csv文件 生成[]JobDataRes
	// 解析CSV数据
	jobDataList, err := parseCSVAll(utf8Reader)
	if err != nil {
		glog.Errorf(ctx, "parseCSV err:%s", err)
		return
	}

	// 打印解析后的数组
	glog.Infof(ctx, "jobDataList:%+v", jobDataList)
}

func parseCSVAll(csvData io.Reader) ([]JobDataRes, error) {
	reader := csv.NewReader(csvData)
	reader.Comma = '\t'

	// 读取CSV文件中的所有记录
	records, err := reader.ReadAll()
	if err != nil {
		glog.Errorf(context.Background(), "error reading CSV: %v", err)
		return nil, err
	}

	// 创建一个存储 JobDataRes 的切片
	var jobDataList []JobDataRes

	var (
		sum = 0
	)

	// 遍历CSV记录并将其映射到 JobDataRes 结构体
	for i, record := range records {
		// 跳过标题行
		if i == 0 {
			continue
		}
		jobData := JobDataRes{
			Date:             record[0],
			ToCampaignName:   record[1],
			ToCampaignId:     record[2],
			ToAppName:        record[3],
			ToAppID:          record[4],
			ToAppBundle:      record[5],
			ToAppPlatform:    record[6],
			FromCampaignId:   record[7],
			FromCampaignName: record[8],
			FromAppName:      record[9],
			FromAppID:        record[10],
			FromAppBundle:    record[11],
			FromAppPlatform:  record[12],
			CampaignType:     record[13],
			Creative:         record[14],
			Role:             record[15],
			AdType:           record[16],
			Impressions:      record[17],
			Clicks:           record[18],
			Installs:         record[19],
			CTR:              record[20],
			IR:               record[21],
			MoneyEarned:      record[22],
			ECPMEarned:       record[23],
			MoneySpent:       record[24],
			ECPMSpent:        record[25],
			CompletedView:    record[26],
		}
		jobDataList = append(jobDataList, jobData)
		sum++
	}
	glog.Infof(context.Background(), "总共:%d条; 实际:%d", len(records)-1, len(jobDataList)) // 总共:6993条; 实际:6993
	return jobDataList, nil
}
