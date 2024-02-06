package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	glog "github.com/mao888/mao-glog"
	"gopkg.in/resty.v1"
	"time"
)

type LiftoffCreativesRes struct {
	PreviewUrl         *string   `json:"preview_url"`
	Width              int       `json:"width"`
	Height             int       `json:"height"`
	CreativeType       string    `json:"creative_type"`
	FullHtmlPreviewUrl string    `json:"full_html_preview_url"`
	CreatedAt          time.Time `json:"created_at"`
	Id                 string    `json:"id"`
	Name               string    `json:"name"`
}

func main() {
	ctx := context.Background()
	var (
		apiKey          = "bacfa09c4f"
		apiSecret       = "U1NUhwT2c1s0GRPka9DmZg=="
		basicLiftoffUrl = "https://data.liftoff.io/api/v1/creatives"
	)
	// 拼接client_id和secret，并转换为字节数组
	data := []byte(apiKey + ":" + apiSecret)
	// 使用base64进行编码
	encoded := base64.StdEncoding.EncodeToString(data)
	authorization := "Basic " + encoded

	resp, err := resty.New().SetRetryCount(3).R().
		SetHeaders(map[string]string{
			"Authorization": authorization,
		}).Get(basicLiftoffUrl)
	if err != nil {
		glog.Errorf(ctx, "Post err:%s", err)
		return
	}
	//glog.Infof(ctx, "resp:%s", string(resp.Body()))

	var res []LiftoffCreativesRes
	err = json.Unmarshal(resp.Body(), &res)
	if err != nil {
		glog.Errorf(ctx, "Unmarshal err:%s", err)
		return
	}
	glog.Infof(ctx, "res的长度:%d", len(res))

	// Id 和 Name 转成一个map
	creatives := make(map[string]string)
	for _, v := range res {
		creatives[v.Id] = v.Name
	}
	glog.Infof(ctx, "creatives map的长度:%d", len(creatives))
}
