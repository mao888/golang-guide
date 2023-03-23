package main

import (
	"fmt"
	"gopkg.in/resty.v1"
)

func main() {
	type videoResp struct {
		Code      int    `json:"code"`
		Message   string `json:"message"`
		MessageCn string `json:"message_cn"`
		Data      struct {
			VideoId string `json:"video_id"`
		} `json:"data"`
	}

	var (
		openApiUrlPrefix string = "https://ad.oceanengine.com/open_api/2/"
		uri              string = "file/video/ad/"
		// 请求参数
		advertiserId int64 = 1
	)
	url := fmt.Sprintf("%s%s", openApiUrlPrefix, uri)
	resp, err := resty.New().SetRetryCount(3).R().
		SetFormData(map[string]string{
			"advertiser_id": fmt.Sprintf("%d", advertiserId),
			"upload_type":   "",
			"filename":      "",
			"video_url":     "",
		}).
		Post(url)
	if err != nil {
		fmt.Println("Post err", err)
		return
	}
	fmt.Println("code:", resp.StatusCode())
}
