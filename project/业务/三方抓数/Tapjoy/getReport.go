package main

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	glog "github.com/mao888/mao-glog"
	gutil "github.com/mao888/mao-gutils/json"
	"github.com/spf13/cast"
	"net/http"
)

// TapjoyReportResp tapjoy 报表结构
type TapjoyReportResp struct {
	Data struct {
		AdSet AdSet `json:"adSet"`
	} `json:"data"`
}
type AdSet struct {
	Ads []*Ads `json:"ads,omitempty"`
	//MultiRewardEngagementSettings []*MultiRewardEngagementSettings `json:"multiRewardEngagementSettings,omitempty"`
}
type Ads struct {
	Edge
	ID       string `json:"id"`
	Name     string `json:"name"`
	Insights struct {
		Reports []TapjoyReport `json:"reports"`
	} `json:"insights"`
}
type Edge struct {
	Node struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Campaign struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"campaign"`
	} `json:"node"`
}
type TapjoyReport struct {
	App struct {
		ID       string `json:"id"`
		Name     string `json:"name"`
		Platform string `json:"platform"`
		BundleID string `json:"bundleId"`
		HashedID string `json:"hashedId"`
	}
	Country            string  `json:"country"`
	CallToActionClicks []int64 `json:"callToActionClicks"`
	Impressions        []int64 `json:"impressions"`
	Conversions        []int64 `json:"conversions"`
	Spend              []int64 `json:"spend"`
}

func main() {
	ctx := context.Background()
	//DragonFarmAdventure_Android_ZXZ_US_broad_ROAS_0202- NewMap	141.5
	//nodeIDArr := []string{"1419d13e-30cd-47a4-a74a-fd1223f88a4a", "e7ba1a1f-72c3-47fb-8419-5f8b7f805053"}

	//DragonFarmAdventure_Android_ZXZ_US_broad_ROAS_0518- 推进程任务线	226.779999
	// [35d86f46-c177-40db-84d3-9a7a21e0d06f 3b138f89-cc70-4d11-a11a-986f50edde68 6083468c-1883-4231-a392-02851e1883ee 7e7d6011-fc4c-4eaf-9be0-2e98803f1ba5 d2b05ad1-b80a-407f-834a-b02aa572ba64]
	//nodeIDArr := []string{"35d86f46-c177-40db-84d3-9a7a21e0d06f", "3b138f89-cc70-4d11-a11a-986f50edde68", "6083468c-1883-4231-a392-02851e1883ee", "7e7d6011-fc4c-4eaf-9be0-2e98803f1ba5", "d2b05ad1-b80a-407f-834a-b02aa572ba64"}

	//DragonFarmAdventure_Android_ZXZ_US_broad_ROAS_Map_0801	0.5
	// [00e9107d-79ea-4a51-a3d8-cb0c58e67cc8 163bcf50-1edd-4ac7-9981-cb41d0e3f582 40a0d6c3-b5dc-4482-ad39-856409d9b2d7 58526b7e-69db-4313-84da-5f61d8164bf9]
	//nodeIDArr := []string{"00e9107d-79ea-4a51-a3d8-cb0c58e67cc8", "163bcf50-1edd-4ac7-9981-cb41d0e3f582", "40a0d6c3-b5dc-4482-ad39-856409d9b2d7", "58526b7e-69db-4313-84da-5f61d8164bf9"}

	//Fairyscapes_Android_ZXZ_KR_broad_ROAS_1123_tapjoy		139.922
	// [92edb4ed-fa4b-495c-9b7a-3593cab21ab5 b8d1df86-8c3b-49e6-8a16-d5c651028ff6]
	//nodeIDArr := []string{"92edb4ed-fa4b-495c-9b7a-3593cab21ab5", "b8d1df86-8c3b-49e6-8a16-d5c651028ff6"}

	//Fairyscapes_Android_ZXZ_US_broad_ROAS_1113_tapjoy	407.71106
	// [605e474b-95ab-4be7-b180-39aab7790121 af783aa2-ddd5-45dc-aded-0de217a59fe3 dbd9810f-5b6f-4f45-997c-2bc3bfa0db1e e8ad2af4-55e2-4849-8fc4-c7adedef1ea4]
	//nodeIDArr := []string{"605e474b-95ab-4be7-b180-39aab7790121", "af783aa2-ddd5-45dc-aded-0de217a59fe3", "dbd9810f-5b6f-4f45-997c-2bc3bfa0db1e", "e8ad2af4-55e2-4849-8fc4-c7adedef1ea4"}

	//Jackpotland_Android_ZXZ_US_ROAS_0407_DR
	//Jackpotland_Android_ZXZ_US_ROAS_0411_DR_Collect

	//Jackpotland_Android_ZXZ_US_tapjoy_ROAS_1020_level		1182.369999
	// [00e680c6-0716-42bc-9c87-9b755f13439b 09539103-329a-419b-a33f-41b3fe93c3fd 1f44ecb1-92be-45c1-9fcc-b6c74bbc9490 2201f8bc-f959-4eef-88d3-413ce479399d 40d87ec2-111c-4353-bd2f-67903bf48c88 42869116-f487-48d0-bfbe-244407dd0104 5f1a6d91-b104-4612-a631-52b945729d41 9dfa020b-afe5-4e1a-8676-c492ccf77c35 e8964420-3ba1-410a-b7da-6f0461a1f3a8 ec9955f2-dc9c-4152-b4b9-cfc841676df4 f2134063-41e3-408d-8358-ad68beebd70e]
	nodeIDArr := []string{"00e680c6-0716-42bc-9c87-9b755f13439b", "09539103-329a-419b-a33f-41b3fe93c3fd", "1f44ecb1-92be-45c1-9fcc-b6c74bbc9490", "2201f8bc-f959-4eef-88d3-413ce479399d", "40d87ec2-111c-4353-bd2f-67903bf48c88", "42869116-f487-48d0-bfbe-244407dd0104", "5f1a6d91-b104-4612-a631-52b945729d41", "9dfa020b-afe5-4e1a-8676-c492ccf77c35", "e8964420-3ba1-410a-b7da-6f0461a1f3a8", "ec9955f2-dc9c-4152-b4b9-cfc841676df4", "f2134063-41e3-408d-8358-ad68beebd70e"}

	//Wordcrush_android_lxi_us_broad_aeo_0930_Tapjoy-MRRT

	spendSum := float64(0)
	for _, nodeID := range nodeIDArr {
		spend, err := getReport(ctx, nodeID)
		if err != nil {
			glog.Errorf(ctx, "getReport err:%s", err)
		}
		spendSum += float64(spend)
	}
	glog.Infof(ctx, "spendSum:%d", spendSum/1000000)
}

func getReport(ctx context.Context, nodeID string) (int64, error) {
	var (
		tapjoyGraphqlReportQuery = "query { adSet(id: \"%s\") { ads { id name insights( timeRange: " +
			"{ from: \"%sT00:00:00Z\" until: \"%sT23:59:59Z\" } ) { reports { app { id name platform bundleId hashedId } " +
			"country callToActionClicks impressions conversions spend } } } } }"
		date             = "2024-04-27"
		tapjoyGraphqlUrl = "https://api.tapjoy.com/graphql"
		token            = "La+4QYFb+PQhwG/T3UMmAyaJPlMOSp98xGXfeFE+lhmo1N+5NJrGcAJNjs+E9L3DleuPxeRkU5jxIwNV5UAZRA=="
	)

	client := resty.New()
	client.SetRetryCount(3)
	client.OnAfterResponse(func(c *resty.Client, resp *resty.Response) error {
		if resp.StatusCode() == http.StatusServiceUnavailable { // 503状态码
			return fmt.Errorf("[%s] report 状态码返回: %d, 触发重试", "tapjoyName", resp.StatusCode())
		}
		return nil // 其他情况不触发重试
	})
	querySearch := fmt.Sprintf(tapjoyGraphqlReportQuery, nodeID, date, date)
	resp, err := client.R().
		SetBody(map[string]interface{}{"query": querySearch}).
		SetHeader("Authorization", "Bearer "+cast.ToString(token)).
		SetHeader("Content-Type", "application/json").
		Post(tapjoyGraphqlUrl)
	if err != nil {
		glog.Errorf(ctx, "[%s] error: %s", "tapjoyName", err.Error())
		return 0, err
	}
	if resp.StatusCode() != http.StatusOK {
		glog.Errorf(ctx, "[%s] status code: %d", "tapjoyName", resp.StatusCode())
		return 0, fmt.Errorf("[%s] status code: %d", "tapjoyName", resp.StatusCode())
	}
	tapjoyResp := &TapjoyReportResp{}
	if err := gutil.JSON2ObjectE(resp.Body(), &tapjoyResp); err != nil {
		glog.Infof(ctx, "[%s] unmarshal error. node_id:%s error: %s", "tapjoyName", nodeID, err.Error())
		glog.Infof(ctx, "[%s] json body is : %s", resp.Body())
		return 0, err
	}
	// 统计Spend
	spend := int64(0)
	for _, report := range tapjoyResp.Data.AdSet.Ads {
		for _, ins := range report.Insights.Reports {
			spend += ins.Spend[0]
		}
	}
	return spend, nil
}
