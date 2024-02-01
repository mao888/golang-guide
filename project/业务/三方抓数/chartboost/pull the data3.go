package main

import (
	"context"
	"fmt"
	glog "github.com/mao888/mao-glog"
	"gopkg.in/resty.v1"
	"strings"
)

type (
	JobDataRes3 struct {
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

	context := string(resp.Body())
	lines := strings.Split(context, "\n")
	var (
		lenFieldsError = 0
		jobDataList    []JobDataRes3
	)
	for i, line := range lines {
		if i == 0 {
			continue
		}
		//glog.Infof(ctx, "line:%s", line)
		// 将 line 转成 JobDataRes3
		// 使用制表符进行分割
		fields := strings.Split(line, "\t")
		// 确保至少有27个字段
		if len(fields) < 27 {
			glog.Errorf(ctx, "第%d行fields长度错误, expected at least 27, got %d; line内容:%s", i, len(fields), line)
			lenFieldsError++
			continue
		}
		jobData := JobDataRes3{
			Date:             fields[0],
			ToCampaignName:   fields[1],
			ToCampaignId:     fields[2],
			ToAppName:        fields[3],
			ToAppID:          fields[4],
			ToAppBundle:      fields[5],
			ToAppPlatform:    fields[6],
			FromCampaignId:   fields[7],
			FromCampaignName: fields[8],
			FromAppName:      fields[9],
			FromAppID:        fields[10],
			FromAppBundle:    fields[11],
			FromAppPlatform:  fields[12],
			CampaignType:     fields[13],
			Creative:         fields[14],
			Role:             fields[15],
			AdType:           fields[16],
			Impressions:      fields[17],
			Clicks:           fields[18],
			Installs:         fields[19],
			CTR:              fields[20],
			IR:               fields[21],
			MoneyEarned:      fields[22],
			ECPMEarned:       fields[23],
			MoneySpent:       fields[24],
			ECPMSpent:        fields[25],
			CompletedView:    fields[26],
		}
		jobDataList = append(jobDataList, jobData)
	}
	glog.Infof(ctx, "lenFieldsError:%d", lenFieldsError)
	glog.Infof(ctx, "应该长度:%d; 实际长度:%d", len(lines)-1, len(jobDataList))
	// 应该长度:9933; 实际长度:9882
	// 应该长度:7015; 实际长度:6972
}
