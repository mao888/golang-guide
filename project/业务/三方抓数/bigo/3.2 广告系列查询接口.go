package main

import (
	"context"
	"encoding/json"
	glog "github.com/mao888/mao-glog"
	"gopkg.in/resty.v1"
)

type (
	CampaignRes struct {
		Result struct {
			Total int `json:"total"`
			List  []struct {
				CalcStatus     int    `json:"calcStatus"`
				UpdateTime     string `json:"updateTime"`
				PromotedObject int    `json:"promotedObject"`
				RtaType        int    `json:"rtaType"`
				CalcSubStatus  int    `json:"calcSubStatus"`
				SkanApp        string `json:"skanApp"`
				BudgetMode     int    `json:"budgetMode"`
				RtaCampaignId  string `json:"rtaCampaignId"`
				CreateTime     string `json:"createTime"`
				Name           string `json:"name"`
				Id             string `json:"id"`
				Budget         string `json:"budget"`
				Status         int    `json:"status"`
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
	)
	campaignUrl := "https://api.adsbigo.com/openapi/campaign/list"

	resp, err := resty.New().SetRetryCount(3).R().
		SetHeaders(map[string]string{
			"Content-Type":  "application/json",
			"Authorization": "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IlA0SkRvOUx1emlkeUVoZlF6SG1qbHRKR0R5bC81UXNKVzBZMDhkTmkrTHM9In0.eyJhdWQiOlsib3BlbmFwaSJdLCJ1c2VyX25hbWUiOiJhZG1pbiIsInNjb3BlIjpbInJlYWQiXSwiZXhwIjoxNzA2NjAzNzE0LCJhdXRob3JpdGllcyI6WyJST0xFX0FETUlOIl0sImp0aSI6IjhkOGI4NWMwLTZjYWMtNDhiZS1hYWU4LTVjN2JhZTc5OTllNSIsImNsaWVudF9pZCI6IjI0MDEyNDQyMDM0NzY5MDg2MDk3OTIifQ.XO4EX_2MOG4Ey2QbDPaY2CKRlB7jihJuXTu4rpeFXUkv6m4nFyhuLknY9R9sgdGVnfrLyva2ZJ8EoK5GS0qgrgAdBWSCe7P0sXxMFyzMn4o_0XUbZyINe2bmvZKNF587lfa-lOPb8JkY_vR3GrTpnC1FaeMi7Mdve2INBup4myNTDa3Jq3w8RJ1ops9Q7CpTAz7gah2vYs4IIztKB_q9hpGJXDgfSgmtHLVNacPM4NSQ0LGyfuKDe4uMgQGGCk-MsY8CdlfBNZplcOpPIfA4_J4z6-YjNbpedGVZCkcy6mK0U5wWL-yd8USTrkFic9oHabicDBRVDpqtsXrxxoTTLA",
			"advertiser-id": "169055",
		}).
		SetBody(map[string]interface{}{
			"pageNo":   pageNo,
			"pageSize": pageSize,
		}).Post(campaignUrl)
	if err != nil {
		glog.Errorf(ctx, "Post err:%s", err)
		return
	}
	glog.Infof(ctx, "resp:%s", string(resp.Body()))

	var campaign CampaignRes
	err = json.Unmarshal(resp.Body(), &campaign)
	if err != nil {
		glog.Errorf(ctx, "Unmarshal err:%s", err)
		return
	}
	glog.Infof(ctx, "campaign:%+v", campaign)
}
