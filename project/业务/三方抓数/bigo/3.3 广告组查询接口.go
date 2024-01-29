package main

import (
	"context"
	"encoding/json"
	glog "github.com/mao888/mao-glog"
	"gopkg.in/resty.v1"
)

type (
	AdSetRes struct {
		Result struct {
			Total int `json:"total"`
			List  []struct {
				AdvertiserId            string        `json:"advertiserId"`
				Age                     []interface{} `json:"age"`
				AppClickTrackAsLandUrl  int           `json:"appClickTrackAsLandUrl"`
				AttributeType           int           `json:"attributeType"`
				Bid                     string        `json:"bid"`
				Budget                  string        `json:"budget"`
				BudgetMode              int           `json:"budgetMode"`
				CalcStatus              int           `json:"calcStatus"`
				CalcSubStatus           int           `json:"calcSubStatus"`
				CampaignId              string        `json:"campaignId"`
				CanChangeDeepGoal       int           `json:"canChangeDeepGoal"`
				CanChangeDeepGoalSwitch int           `json:"canChangeDeepGoalSwitch"`
				CreateTime              string        `json:"createTime"`
				CreativeType            int           `json:"creativeType"`
				DeepBid                 string        `json:"deepBid"`
				DelayPostback           int           `json:"delayPostback"`
				DevicePriceSection      []interface{} `json:"devicePriceSection"`
				FlowControlMode         int           `json:"flowControlMode"`
				Gender                  string        `json:"gender"`
				HighestAndroidVersion   string        `json:"highestAndroidVersion"`
				HighestIosVersion       string        `json:"highestIosVersion"`
				Id                      string        `json:"id"`
				InstallOpt              int           `json:"installOpt"`
				LandingAppId            string        `json:"landingAppId"`
				LandingPage             string        `json:"landingPage"`
				LandingType             int           `json:"landingType"`
				Language                []interface{} `json:"language"`
				Location                []string      `json:"location"`
				LowestAndroidVersion    string        `json:"lowestAndroidVersion"`
				LowestIosVersion        string        `json:"lowestIosVersion"`
				MediaApp                []string      `json:"mediaApp"`
				MediaAppPlaceIds        struct {
					Imo   []interface{} `json:"imo"`
					Likee []interface{} `json:"likee"`
					Ban   []interface{} `json:"ban"`
				} `json:"mediaAppPlaceIds"`
				Name             string        `json:"name"`
				Network          []interface{} `json:"network"`
				NonEndTime       int           `json:"nonEndTime"`
				OcpcLearnStatus  int           `json:"ocpcLearnStatus"`
				OcpcOptType      int           `json:"ocpcOptType"`
				Operation        []interface{} `json:"operation"`
				OptimizationGoal int           `json:"optimizationGoal"`
				Os               int           `json:"os"`
				PkgName          string        `json:"pkgName"`
				PlacementType    int           `json:"placementType"`
				PricingType      int           `json:"pricingType"`
				SecondBid        string        `json:"secondBid"`
				StartTime        string        `json:"startTime"`
				Status           int           `json:"status"`
				TimeBucket       struct {
				} `json:"timeBucket"`
				UpdateTime string `json:"updateTime"`
			} `json:"list"`
		} `json:"result"`
		RequestId string `json:"requestId"`
		Retmsg    string `json:"retmsg"`
		Retcode   int    `json:"retcode"`
	}
)

func main() {
	var (
		ctx      = context.Background()
		pageNo   = 1
		pageSize = 200
		adsetUrl = "https://api.adsbigo.com/openapi/adset/list"
	)

	resp, err := resty.New().SetRetryCount(3).R().
		SetHeaders(map[string]string{
			"Content-Type":  "application/json",
			"Authorization": "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IlA0SkRvOUx1emlkeUVoZlF6SG1qbHRKR0R5bC81UXNKVzBZMDhkTmkrTHM9In0.eyJhdWQiOlsib3BlbmFwaSJdLCJ1c2VyX25hbWUiOiJhZG1pbiIsInNjb3BlIjpbInJlYWQiXSwiZXhwIjoxNzA2NjAzNzE0LCJhdXRob3JpdGllcyI6WyJST0xFX0FETUlOIl0sImp0aSI6IjhkOGI4NWMwLTZjYWMtNDhiZS1hYWU4LTVjN2JhZTc5OTllNSIsImNsaWVudF9pZCI6IjI0MDEyNDQyMDM0NzY5MDg2MDk3OTIifQ.XO4EX_2MOG4Ey2QbDPaY2CKRlB7jihJuXTu4rpeFXUkv6m4nFyhuLknY9R9sgdGVnfrLyva2ZJ8EoK5GS0qgrgAdBWSCe7P0sXxMFyzMn4o_0XUbZyINe2bmvZKNF587lfa-lOPb8JkY_vR3GrTpnC1FaeMi7Mdve2INBup4myNTDa3Jq3w8RJ1ops9Q7CpTAz7gah2vYs4IIztKB_q9hpGJXDgfSgmtHLVNacPM4NSQ0LGyfuKDe4uMgQGGCk-MsY8CdlfBNZplcOpPIfA4_J4z6-YjNbpedGVZCkcy6mK0U5wWL-yd8USTrkFic9oHabicDBRVDpqtsXrxxoTTLA",
			"advertiser-id": "169055",
		}).SetBody(map[string]interface{}{
		"pageNo":   pageNo,
		"pageSize": pageSize,
	}).Post(adsetUrl)
	if err != nil {
		glog.Errorf(ctx, "Post err:%s", err)
		return
	}
	glog.Infof(ctx, "resp:%s", string(resp.Body()))

	var adset AdSetRes
	err = json.Unmarshal(resp.Body(), &adset)
	if err != nil {
		glog.Errorf(ctx, "Unmarshal err:%s", err)
		return
	}
	glog.Infof(ctx, "adset:%+v", adset)
}
