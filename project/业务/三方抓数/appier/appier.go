package main

import (
	"context"
	"encoding/json"
	"fmt"
	glog "github.com/mao888/mao-glog"
	"gopkg.in/resty.v1"
)

type (
	AppierRes struct {
		Date         string  `json:"date"`
		CampaignId   string  `json:"campaign_id"`
		CampaignName string  `json:"campaign_name"`
		Impressions  int     `json:"impressions"`
		Clicks       int     `json:"clicks"`
		Installs     int     `json:"installs"`
		Cost         float32 `json:"cost"`
		Currency     string  `json:"currency"`
		Timezone     int     `json:"timezone"`
	}
)

func main() {
	ctx := context.Background()
	var (
		basicUrl    = "https://mmp.appier.org/campaign_report"
		accessToken = "f8a5f28b4c9f405797e8359737df122e"
		startDate   = "2024-02-22"
		endDate     = "2024-02-22"
		timezone    = -8
	)
	url := fmt.Sprintf("%s?access_token=%s&start_date=%s&end_date=%s&timezone=%d",
		basicUrl, accessToken, startDate, endDate, timezone)
	glog.Infof(ctx, "url:%s", url)

	resp, err := resty.New().SetRetryCount(3).R().Get(url)
	if err != nil {
		glog.Errorf(ctx, "Get err:%s", err)
		return
	}
	glog.Infof(ctx, "resp:%s", string(resp.Body()))

	var appierRes []*AppierRes
	err = json.Unmarshal(resp.Body(), &appierRes)
	if err != nil {
		glog.Errorf(ctx, "Unmarshal err:%s", err)
		return
	}
	glog.Infof(ctx, "appierRes:%v", appierRes)
	glog.Infof(ctx, "appierRes的长度:%d", len(appierRes))
}
