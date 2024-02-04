package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	glog "github.com/mao888/mao-glog"
	"gopkg.in/resty.v1"
	"time"
)

type (
	LiftoffReportStatusRes struct {
		CreatedAt  time.Time `json:"created_at"`
		State      string    `json:"state"`
		Parameters struct {
			Format                        string      `json:"format"`
			Timezone                      string      `json:"timezone"`
			EventIds                      interface{} `json:"event_ids"`
			CampaignIds                   interface{} `json:"campaign_ids"`
			UseTwoLetterCountry           bool        `json:"use_two_letter_country"`
			VideoPlayMilestones           bool        `json:"video_play_milestones"`
			RemoveZeroRows                bool        `json:"remove_zero_rows"`
			StartTime                     time.Time   `json:"start_time"`
			AppIds                        interface{} `json:"app_ids"`
			EndTime                       time.Time   `json:"end_time"`
			IncludeSkanCensoredInstalls   bool        `json:"include_skan_censored_installs"`
			GroupBy                       []string    `json:"group_by"`
			IncludeRepeatEvents           bool        `json:"include_repeat_events"`
			IncludeSkanUncensoredInstalls bool        `json:"include_skan_uncensored_installs"`
			EventCount                    int         `json:"event_count"`
			IncludeReengagedEvents        bool        `json:"include_reengaged_events"`
			CohortWindow                  interface{} `json:"cohort_window"`
			CallbackUrl                   interface{} `json:"callback_url"`
		} `json:"parameters"`
		Id string `json:"id"`
	}
)

func main() {
	ctx := context.Background()
	var (
		apiKey          = "bacfa09c4f"
		apiSecret       = "U1NUhwT2c1s0GRPka9DmZg=="
		basicLiftoffUrl = "https://data.liftoff.io/api/v1/reports"

		reportId = "94ca462dfe"
	)
	// 拼接client_id和secret，并转换为字节数组
	data := []byte(apiKey + ":" + apiSecret)
	// 使用base64进行编码
	encoded := base64.StdEncoding.EncodeToString(data)
	authorization := "Basic " + encoded

	requestUrl := fmt.Sprintf("%s/%s/status", basicLiftoffUrl, reportId)
	resp, err := resty.New().SetRetryCount(3).R().
		SetHeaders(map[string]string{
			"Authorization": authorization,
		}).Get(requestUrl)
	if err != nil {
		glog.Errorf(ctx, "Post err:%s", err)
		return
	}
	glog.Infof(ctx, "resp:%s", string(resp.Body()))

	var res LiftoffReportStatusRes
	err = json.Unmarshal(resp.Body(), &res)
	if err != nil {
		glog.Errorf(ctx, "Unmarshal err:%s", err)
		return
	}
	glog.Infof(ctx, "res:%+v", res)

	if res.State == "completed" {
		glog.Infof(ctx, "reportId:%s is completed", reportId)
	}
}
