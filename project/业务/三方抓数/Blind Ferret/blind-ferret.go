package main

import (
	"context"
	"encoding/json"
	"fmt"
	glog "github.com/mao888/mao-glog"
	"gopkg.in/resty.v1"
)

type (
	BlindFerretRes struct {
		Date                 string `json:"Date"`
		App                  string `json:"App"`
		StoreIDBundleID      string `json:"Store ID (Bundle ID)"`
		FlourishCampaignID   string `json:"Flourish Campaign ID"`
		FlourishCampaignName string `json:"Flourish Campaign Name"`
		CampaignID           string `json:"Campaign ID"`
		CampaignName         string `json:"Campaign Name"`
		SubCampaignID        string `json:"Sub-Campaign ID"`
		SubCampaignName      string `json:"Sub-Campaign Name"`
		FlourishOfferName    string `json:"Flourish Offer Name"`
		PublisherName        string `json:"Publisher Name"`
		PublisherID          string `json:"Publisher ID"`
		TrackingURL          string `json:"Tracking URL"`
		CreativeID           string `json:"Creative ID"`
		CreativeName         string `json:"Creative Name"`
		CreativeURL          string `json:"Creative URL"`
		Platform             string `json:"Platform"`
		Country              string `json:"Country"`
		ReportingTimezone    string `json:"Reporting Timezone"`
		Currency             string `json:"Currency"`
		Impressions          string `json:"Impressions"`
		Clicks               string `json:"Clicks"`
		Installs             string `json:"Installs"`
		Cost                 string `json:"Cost"`
	}
)

func main() {
	ctx := context.Background()
	var (
		format = "json"
		apiKey = "df0a831e4c69906642a435e452c40b93"
		date   = "2024-02-01"
	)
	url := fmt.Sprintf("https://engage-network.influencemobile.com/reports/v2/offers.%s?api_key=%s&date=%s", format, apiKey, date)
	resp, err := resty.New().SetRetryCount(3).R().Get(url)
	if err != nil {
		glog.Errorf(ctx, "Get err:%s", err)
		return
	}
	//glog.Infof(ctx, "resp:%s", string(resp.Body()))

	var blindFerretRes []BlindFerretRes
	err = json.Unmarshal(resp.Body(), &blindFerretRes)
	if err != nil {
		glog.Errorf(ctx, "Unmarshal err:%s", err)
		return
	}
	//glog.Infof(ctx, "blindFerretRes:%v", blindFerretRes)
	glog.Infof(ctx, "blindFerretRes的长度:%d", len(blindFerretRes))
}
