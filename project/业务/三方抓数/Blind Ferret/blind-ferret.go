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

	BlindFerretToS3 struct {
		Date            string `json:"date"`
		App             string `json:"app"`
		BundleID        string `json:"bundle_id"`
		CampaignID      string `json:"campaign_id"`
		CampaignName    string `json:"campaign_name"`
		SubCampaignID   string `json:"sub_campaign_id"`
		SubCampaignName string `json:"sub_campaign_name"`
		CreativeID      string `json:"creative_id"`
		CreativeName    string `json:"creative_name"`
		CreativeURL     string `json:"creative_url"`
		PublisherID     string `json:"publisher_id"`
		PublisherName   string `json:"publisher_name"`
		TrackingURL     string `json:"tracking_url"`
		Country         string `json:"country"`
		Timezone        string `json:"timezone"`
		Platform        string `json:"platform"`
		Impressions     string `json:"impressions"`
		Clicks          string `json:"clicks"`
		Installs        string `json:"installs"`
		Cost            string `json:"cost"`

		CaptureDate string `json:"capture_date"`
		CreateTime  int64  `json:"create_time"`
		UpdateTime  int64  `json:"update_time"`
		Version     int64  `json:"version"`
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

	// 将vo.BlindFerretRes转换为vo.BlindFerretToS3
	var blindFerretToS3s []*BlindFerretToS3
	for _, res := range blindFerretRes {
		blindFerretToS3 := &BlindFerretToS3{
			Date:            res.Date,
			App:             res.App,
			BundleID:        res.StoreIDBundleID,
			CampaignID:      res.CampaignID,
			CampaignName:    res.CampaignName,
			SubCampaignID:   res.SubCampaignID,
			SubCampaignName: res.SubCampaignName,
			CreativeID:      res.CreativeID,
			CreativeName:    res.CreativeName,
			CreativeURL:     res.CreativeURL,
			PublisherID:     res.PublisherID,
			PublisherName:   res.PublisherName,
			TrackingURL:     res.TrackingURL,
			Country:         res.Country,
			Timezone:        res.ReportingTimezone,
			Platform:        res.Platform,
			Impressions:     res.Impressions,
			Clicks:          res.Clicks,
			Installs:        res.Installs,
			Cost:            res.Cost,
		}
		blindFerretToS3s = append(blindFerretToS3s, blindFerretToS3)
	}
	glog.Infof(ctx, "BlindFerretToS3 的长度:%d", len(blindFerretToS3s))
}
