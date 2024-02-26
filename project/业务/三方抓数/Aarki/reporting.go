package main

import (
	"context"
	"encoding/json"
	"fmt"
	glog "github.com/mao888/mao-glog"
	"gopkg.in/resty.v1"
	"time"
)

type (
	AarkiRes struct {
		Status     string        `json:"status"`
		Metrics    []string      `json:"metrics"`
		Data       []interface{} `json:"data"`
		Parameters struct {
			ByStoreIdentifier string `json:"by_store_identifier"`
			ByCountry         string `json:"by_country"`
			EndDate           string `json:"end_date"`
			ByCreative        string `json:"by_creative"`
			BySize            string `json:"by_size"`
			ByCampaign        string `json:"by_campaign"`
			Timezone          string `json:"timezone"`
			ByCampaignTag     string `json:"by_campaign_tag"`
			ByPlatform        string `json:"by_platform"`
			StartDate         string `json:"start_date"`
		} `json:"parameters"`
		Error string `json:"error"`
	}
)

func main() {
	ctx := context.Background()
	var (
		basicUrl            = "https://encore.aarki.com/dsp/api/v2"
		token               = "xlcbz3066p5ikn5fxlmj4ualdk09bx1d"
		timezone            = "US/Pacific"
		start_date          = "2024-02-20"
		end_date            = "2024-02-22"
		by_campaign         = "y"
		by_campaign_tag     = "y"
		by_country          = "y"
		by_store_identifier = "y"
		by_platform         = "y"
		by_creative         = "y"
		by_size             = "y"
	)
	url := fmt.Sprintf("%s/%s?token=%s&timezone=%s&start_date=%s&end_date=%s&by_campaign=%s&by_campaign_tag=%s&by_country=%s&by_store_identifier=%s&by_platform=%s&by_creative=%s&by_size=%s",
		basicUrl, "account_summary.json", token, timezone, start_date, end_date, by_campaign, by_campaign_tag, by_country, by_store_identifier, by_platform, by_creative, by_size)
	glog.Infof(ctx, "url:%s", url)

	resp, err := resty.New().SetRetryCount(3).R().Post(url)
	if err != nil {
		glog.Errorf(ctx, "Get err:%s", err)
		return
	}
	glog.Infof(ctx, "resp:%s", string(resp.Body()))

	if resp.StatusCode() != 200 {
		if resp.StatusCode() == 429 {
			glog.Infof(ctx, "resp.StatusCode:%d", resp.StatusCode())
			time.Sleep(1 * time.Minute)
		} else {
			glog.Errorf(ctx, "resp.StatusCode:%d", resp.StatusCode())
			return
		}
	}

	var aarkiRes AarkiRes
	err = json.Unmarshal(resp.Body(), &aarkiRes)
	if err != nil {
		glog.Errorf(ctx, "Unmarshal err:%s", err)
		return
	}
	glog.Infof(ctx, "aarkiRes:%v", aarkiRes)
	glog.Infof(ctx, "aarkiRes的长度:%d", len(aarkiRes.Data))
}
