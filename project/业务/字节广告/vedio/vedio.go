package main

import (
	"encoding/json"
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
		// 请求Header
		contentType string = "multipart/form-data"
		accessToken string
		//XDebugMode  int = 1
		// 请求参数
		advertiserId int64  = 1760312309087432 // 广告主ID
		uploadType   string = "UPLOAD_BY_URL"  // 视频上传方式，可选值:UPLOAD_BY_FILE: 文件上传（默认值），UPLOAD_BY_URL: 网址上传
		filename     string = ""               // 素材的文件名，可自定义素材名，不传择默认取文件名，最长255个字符。UPLOAD_BY_URL必填  注：若同一素材已进行上传，重新上传不会改名。
		videoUrl     string = ""               // 视频url地址
		//
		ttVideoResp videoResp
	)
	url := fmt.Sprintf("%s%s", openApiUrlPrefix, uri)
	resp, err := resty.New().SetRetryCount(3).R().
		SetHeaders(map[string]string{
			"Content-Type": contentType,
			"Access-Token": accessToken,
		}).
		SetFormData(map[string]string{
			"advertiser_id": fmt.Sprintf("%d", advertiserId),
			"upload_type":   uploadType,
			"filename":      filename,
			"video_url":     videoUrl,
		}).
		Post(url)
	if err != nil {
		fmt.Println("Post err", err)
		return
	}
	fmt.Println("code:", resp.StatusCode())
	err = json.Unmarshal(resp.Body(), &ttVideoResp)
	if err != nil {
		fmt.Println("Unmarshal err:", err)
		return
	}
}
