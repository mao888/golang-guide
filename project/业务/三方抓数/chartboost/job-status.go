package main

import (
	"context"
	"encoding/json"
	"fmt"
	glog "github.com/mao888/mao-glog"
	"gopkg.in/resty.v1"
)

type (
	JobStatusRes struct {
		Status string `json:"status"`
	}
)

func main() {
	ctx := context.Background()
	var (
		chartboostJobsUrl = "https://analytics.chartboost.com/v3/metrics/jobs/"
		jobId             = "947231990b73fe59bc6150cd47333f6c/b7b1ca22-0977-43f9-9096-81e79d0b0e8e"
	)
	// https://analytics.chartboost.com/v3/metrics/jobs/947231990b73fe59bc6150cd47333f6c/b7b1ca22-0977-43f9-9096-81e79d0b0e8e?status=true
	requestUrl := fmt.Sprintf("%s%s?status=true", chartboostJobsUrl, jobId)
	resp, err := resty.New().SetRetryCount(3).R().Get(requestUrl)
	if err != nil {
		glog.Errorf(ctx, "Post err:%s", err)
		return
	}
	glog.Infof(ctx, "resp:%s", string(resp.Body()))

	var jobStatusRes JobStatusRes
	err = json.Unmarshal(resp.Body(), &jobStatusRes)
	if err != nil {
		glog.Errorf(ctx, "Unmarshal err:%s", err)
		return
	}
	glog.Infof(ctx, "jobStatusRes:%+v", jobStatusRes)

	if jobStatusRes.Status == "created" {
		glog.Infof(ctx, "jobStatusRes.Status:%s", jobStatusRes.Status)
		return
	}
}
