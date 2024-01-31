package main

import (
	"context"
	"encoding/json"
	"fmt"
	glog "github.com/mao888/mao-glog"
	"gopkg.in/resty.v1"
)

type (
	CampaignRes struct {
		JobId string `json:"jobId"`
	}
)

func main() {
	ctx := context.Background()
	var (
		chartboostCampaignUrl = "https://analytics.chartboost.com/v3/metrics/campaign"
		dateMin               = "2024-01-23"
		dataMax               = "2024-01-23"
		userId                = "5efb08d7f831010991faea60"
		userSignature         = "f34e949b5cddf3afd46bb8f5bc7e4dae3692107def924882dce53a66dc18c4cf"
		groupBy               = "app,creative"
	)
	// https://analytics.chartboost.com/v3/metrics/campaign?dateMin=2024-01-23&dateMax=2024-01-23&userId=5efb08d7f831010991faea60&userSignature=f34e949b5cddf3afd46bb8f5bc7e4dae3692107def924882dce53a66dc18c4cf&groupBy=app,creative
	requestUrl := fmt.Sprintf("%s?dateMin=%s&dateMax=%s&userId=%s&userSignature=%s&groupBy=%s",
		chartboostCampaignUrl, dateMin, dataMax, userId, userSignature, groupBy)
	resp, err := resty.New().SetRetryCount(3).R().Get(requestUrl)
	if err != nil {
		glog.Errorf(ctx, "Post err:%s", err)
		return
	}
	glog.Infof(ctx, "resp:%s", string(resp.Body()))

	var campaignRes CampaignRes
	err = json.Unmarshal(resp.Body(), &campaignRes)
	if err != nil {
		glog.Errorf(ctx, "Unmarshal err:%s", err)
		return
	}
	glog.Infof(ctx, "campaignRes:%+v", campaignRes)
}
