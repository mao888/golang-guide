package main

import (
	"context"
	"encoding/json"
	"github.com/mao888/mao-glog"
	"gopkg.in/resty.v1"
)

type (
	reportRes struct {
		Result struct {
			Total int `json:"total"`
			List  []struct {
				AdId                 int64   `json:"adId"`
				AdName               string  `json:"adName"`
				AdsetId              int64   `json:"adsetId"`
				AdvertiserId         int     `json:"advertiserId"`
				AggregateTime        string  `json:"aggregateTime"`
				CampaignId           int64   `json:"campaignId"`
				Click                int     `json:"click"`
				Country              string  `json:"country"`
				Impression           int     `json:"impression"`
				RealResultConversion int     `json:"real_result_conversion"`
				TgtPkgName           string  `json:"tgtPkgName"`
				TotalCost            float64 `json:"totalCost"`
			} `json:"list"`
		} `json:"result"`
		RequestId string `json:"requestId"`
		Retmsg    string `json:"retmsg"`
		Retcode   int    `json:"retcode"`
	}
)

func main() {
	var (
		ctx           = context.Background()
		timezone      = "-8"
		startDate     = "2024-01-20"
		endDate       = "2024-01-29"
		aggregateType = 2
		pageNo        = 1
		pageSize      = 200
	)

	reportUrl := "https://api.adsbigo.com/openapi/report/list"

	resp, err := resty.New().SetRetryCount(3).R().
		SetHeaders(map[string]string{
			"Content-Type":  "application/json",
			"Authorization": "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IlA0SkRvOUx1emlkeUVoZlF6SG1qbHRKR0R5bC81UXNKVzBZMDhkTmkrTHM9In0.eyJhdWQiOlsib3BlbmFwaSJdLCJ1c2VyX25hbWUiOiJhZG1pbiIsInNjb3BlIjpbInJlYWQiXSwiZXhwIjoxNzA2NjAzNzE0LCJhdXRob3JpdGllcyI6WyJST0xFX0FETUlOIl0sImp0aSI6IjhkOGI4NWMwLTZjYWMtNDhiZS1hYWU4LTVjN2JhZTc5OTllNSIsImNsaWVudF9pZCI6IjI0MDEyNDQyMDM0NzY5MDg2MDk3OTIifQ.XO4EX_2MOG4Ey2QbDPaY2CKRlB7jihJuXTu4rpeFXUkv6m4nFyhuLknY9R9sgdGVnfrLyva2ZJ8EoK5GS0qgrgAdBWSCe7P0sXxMFyzMn4o_0XUbZyINe2bmvZKNF587lfa-lOPb8JkY_vR3GrTpnC1FaeMi7Mdve2INBup4myNTDa3Jq3w8RJ1ops9Q7CpTAz7gah2vYs4IIztKB_q9hpGJXDgfSgmtHLVNacPM4NSQ0LGyfuKDe4uMgQGGCk-MsY8CdlfBNZplcOpPIfA4_J4z6-YjNbpedGVZCkcy6mK0U5wWL-yd8USTrkFic9oHabicDBRVDpqtsXrxxoTTLA",
		}).
		SetBody(map[string]interface{}{
			"timezone":  timezone,
			"startDate": startDate,
			"endDate":   endDate,
			"indicators": []string{
				"totalCost",
				"click",
				"impression",
				"real_result_conversion",
			},
			"breakDowns": []string{
				"campaignId",
				"adsetId",
				"adId",
				"advertiserId",
				"country",
				"tgtPkgName",
			},
			"aggregateType": aggregateType,
			"pageNo":        pageNo,
			"pageSize":      pageSize,
		}).Post(reportUrl)
	if err != nil {
		glog.Errorf(ctx, "Post err:%s", err)
		return
	}
	glog.Infof(ctx, "resp:%s", string(resp.Body()))

	var report reportRes
	err = json.Unmarshal(resp.Body(), &report)
	if err != nil {
		glog.Errorf(ctx, "Unmarshal err:%s", err)
		return
	}
	glog.Infof(ctx, "report:%+v", report)
}
