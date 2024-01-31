package main

import (
	"context"
	"encoding/csv"
	"fmt"
	glog "github.com/mao888/mao-glog"
	"gopkg.in/resty.v1"
	"io"
	"regexp"
	"strings"
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
		jobId             = "0391cd5465d59e20f7e2d4fa6028c513/2eeb5b6e-351e-42f2-8984-059623824a90"
	)
	// https://analytics.chartboost.com/v3/metrics/jobs/947231990b73fe59bc6150cd47333f6c/b7b1ca22-0977-43f9-9096-81e79d0b0e8e
	requestUrl := fmt.Sprintf("%s%s", chartboostJobsUrl, jobId)
	resp, err := resty.New().SetRetryCount(3).R().Get(requestUrl)
	if err != nil {
		glog.Errorf(ctx, "Post err:%s", err)
		return
	}
	//glog.Infof(ctx, "resp:%s", string(resp.Body()))

	// 读取resp.Body()中的数据,为csv文件,然后解析csv文件 生成[]JobDataRes
	// 解析CSV数据
	jobDataList, err := parseCSV(resp.Body())
	if err != nil {
		glog.Errorf(ctx, "parseCSV err:%s", err)
		return
	}

	// 打印解析后的数据和长度
	glog.Infof(ctx, "len(jobDataList):%d", len(jobDataList)) // 8894
}

//func parseCSV(csvData []byte) ([]JobDataRes, error) {
//	// 使用 strings.NewReader 创建一个 Reader 对象
//	reader := csv.NewReader(strings.NewReader(string(csvData)))
//
//	// 设置 LazyQuotes 和 TrimLeadingSpace 为 true，容忍引号不完全匹配和字段之间的空格
//	reader.LazyQuotes = true
//	reader.TrimLeadingSpace = true
//
//	// 读取CSV文件中的所有记录
//	records, err := reader.ReadAll()
//	if err != nil {
//		return nil, fmt.Errorf("error reading CSV: %v", err)
//	}
//
//	// 创建一个存储 JobDataRes 的切片
//	var jobDataList []JobDataRes
//
//	// 遍历CSV记录并将其映射到 JobDataRes 结构体
//	for i, record := range records {
//		// 跳过标题行
//		if i == 0 {
//			continue
//		}
//		jobData := JobDataRes{
//			Date:             record[0],
//			ToCampaignName:   record[1],
//			ToCampaignId:     record[2],
//			ToAppName:        record[3],
//			ToAppID:          record[4],
//			ToAppBundle:      record[5],
//			ToAppPlatform:    record[6],
//			FromCampaignId:   record[7],
//			FromCampaignName: record[8],
//			FromAppName:      record[9],
//			FromAppID:        record[10],
//			FromAppBundle:    record[11],
//			FromAppPlatform:  record[12],
//			CampaignType:     record[13],
//			Creative:         record[14],
//			Role:             record[15],
//			AdType:           record[16],
//			Impressions:      record[17],
//			Clicks:           record[18],
//			Installs:         record[19],
//			CTR:              record[20],
//			IR:               record[21],
//			MoneyEarned:      record[22],
//			ECPMEarned:       record[23],
//			MoneySpent:       record[24],
//			ECPMSpent:        record[25],
//			CompletedView:    record[26],
//		}
//		jobDataList = append(jobDataList, jobData)
//	}
//
//	return jobDataList, nil
//}

func parseCSV(csvData []byte) ([]JobDataRes, error) {
	// 使用 strings.NewReader 创建一个 Reader 对象
	reader := csv.NewReader(strings.NewReader(string(csvData)))

	// 设置 LazyQuotes 和 TrimLeadingSpace 为 true，容忍引号不完全匹配和字段之间的空格
	reader.LazyQuotes = true
	reader.TrimLeadingSpace = true

	// 读取CSV文件中的标题行
	_, err := reader.Read()
	if err != nil {
		return nil, fmt.Errorf("error reading CSV header: %v", err)
	}

	// 创建一个存储 JobDataRes 的切片
	var jobDataList []JobDataRes

	i := 2
	wrongNumberOfFields := 0
	lenRecordError := 0
	lenParsedError := 0
	// 遍历CSV记录并将其映射到 JobDataRes 结构体
	for {
		record, err := reader.Read()
		if err != nil {
			// 判断是否是文件结束错误
			if err == io.EOF {
				break
			}
			// 打印行数、本行数据、错误信息
			//glog.Errorf(context.Background(), "line:%d; record:%+v; err:%s", i, record, err)
			wrongNumberOfFields++
		}

		// 检查记录的长度是否符合预期
		if len(record) != 1 {
			fmt.Errorf("wrong number of fields in CSV record, expected 1, got %d", len(record))
			glog.Errorf(context.Background(), "line:%d; record:%+v; len(record):%d", i, record, len(record))
			i = i + 1
			lenRecordError++
			continue
		}

		parsed := parseRecord(record[0])
		// 确保至少有27个字段
		if len(parsed) < 27 {
			glog.Errorf(context.Background(), "line:%d; record:%+v; parsed:%+v", i, record, parsed)
			fmt.Errorf("line:%d; not enough fields in CSV record, expected at least 27, got %d", i, len(parsed))
			i = i + 1
			lenParsedError++
			continue
		}

		jobData := JobDataRes{
			Date:             parsed[0],
			ToCampaignName:   parsed[1],
			ToCampaignId:     parsed[2],
			ToAppName:        parsed[3],
			ToAppID:          parsed[4],
			ToAppBundle:      parsed[5],
			ToAppPlatform:    parsed[6],
			FromCampaignId:   parsed[7],
			FromCampaignName: parsed[8],
			FromAppName:      parsed[9],
			FromAppID:        parsed[10],
			FromAppBundle:    parsed[11],
			FromAppPlatform:  parsed[12],
			CampaignType:     parsed[13],
			Creative:         parsed[14],
			Role:             parsed[15],
			AdType:           parsed[16],
			Impressions:      parsed[17],
			Clicks:           parsed[18],
			Installs:         parsed[19],
			CTR:              parsed[20],
			IR:               parsed[21],
			MoneyEarned:      parsed[22],
			ECPMEarned:       parsed[23],
			MoneySpent:       parsed[24],
			ECPMSpent:        parsed[25],
			CompletedView:    parsed[26],
		}
		jobDataList = append(jobDataList, jobData)
		i++
	}
	glog.Infof(context.Background(), "wrongNumberOfFields:%d", wrongNumberOfFields)
	glog.Infof(context.Background(), "lenRecordError:%d", lenRecordError)
	glog.Infof(context.Background(), "lenParsedError:%d", lenParsedError)
	return jobDataList, nil
}

func parseRecord(record string) []string {
	// 使用正则表达式提取引号中的内容
	re := regexp.MustCompile("\"(.*?)\"")
	matches := re.FindAllString(record, -1)
	// 去掉每个子字符串中的双引号
	for i, match := range matches {
		matches[i] = match[1 : len(match)-1]
	}
	return matches
}
