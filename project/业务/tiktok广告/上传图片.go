package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/resty.v1"
)

type (
	TikTokUploadImageRes struct {
		Message string `json:"message"`
		Code    int    `json:"code"`
		Data    struct {
			ImageId     string `json:"image_id"`    // 图片ID，用于广告投放中创建广告
			MaterialId  string `json:"material_id"` // 素材ID
			Displayable bool   `json:"displayable"` // 图片能否在平台中展示
			Width       int    `json:"width"`       // 图片宽度
			Format      string `json:"format"`      // 图片格式
			Url         string `json:"url"`         // 图片URL，1小时有效期，过期后需重新获取
			Height      int    `json:"height"`      // 图片高度
			Signature   string `json:"signature"`   // 图片MD5
			Size        int    `json:"size"`        // 图片大小，单位byte
			FileName    string `json:"file_name"`   // 图片名称
			CreateTime  string `json:"create_time"` // 创建时间。UTC 时间，格式：2020-06-10T07:39:14Z
			ModifyTime  string `json:"modify_time"` // 修改时间。UTC 时间，格式：2020-06-10T07:39:14Z
		} `json:"data"`
		RequestId string `json:"request_id"` // 请求的日志id，唯一标识一个请求
	}
)

func main() {
	var (
		uploadVideoBaseUrl = "https://business-api.tiktok.com/open_api/v1.3/file/image/ad/upload/"
		accessToken        = "be819b0981e021fd4b1914a4b769c6b320f54938"
		advertiserId       = "7306422754532868097"
		//advertiserId       = "7306422822664994817"
		fileName   = "图片3"
		uploadType = "UPLOAD_BY_URL"
		imageUrl   = "https://ark-oss.bettagames.com/2023-11/2363e3a3f806b1bde2657fd2024946d7.jpg"
	)
	resp, err := resty.SetRetryCount(3).R().SetHeaders(map[string]string{
		"Content-Type": "application/json",
		"Access-Token": accessToken,
	}).SetBody(map[string]interface{}{
		"advertiser_id": advertiserId,
		"file_name":     fileName, // 默认值为文件名称或者URL的最后一个 路径名
		"upload_type":   uploadType,
		"image_url":     imageUrl, // 视频文件的URL 当 upload_type 为UPLOAD_BY_URL 时必填
	}).Post(uploadVideoBaseUrl)
	if err != nil {
		fmt.Println("Post err", err)
		return
	}
	fmt.Println("resp", resp)
	// {"code": 0, "message": "OK", "request_id": "202312010902214460BFDB426A924C7F2F", "data": {"signature": "2363e3a3f806b1bde2657fd2024946d7", "modify_time": "2023-12-01T09:00:24Z", "is_carousel_usable": true, "size": 1369062, "height": 2160, "format": "jpeg", "file_name": "\u56fe\u72471", "displayable": false, "image_id": "ad-site-i18n-sg/202312015d0d0400f0d6275d46d9a8cd", "image_url": "https://p21-ad-sg.ibyteimg.com/obj/ad-site-i18n-sg/202312015d0d0400f0d6275d46d9a8cd", "create_time": "2023-12-01T09:00:25Z", "width": 3240, "material_id": "7307518219009622018"}}

	var imageRes TikTokUploadImageRes
	err = json.Unmarshal(resp.Body(), &imageRes)
	if err != nil {
		fmt.Println("Unmarshal err", err)
		return
	}
	fmt.Println("imageRes", imageRes)
	// {OK 0 {ad-site-i18n-sg/202312015d0d0400f0d6275d46d9a8cd 7307518219009622018 false 3240 jpeg  2160 2363e3a3f806b1bde2657fd2024946d7 1369062 图片1 2023-12-01T09:00:25Z 2023-12-01T09:00:24Z} 202312010902214460BFDB426A924C7F}
}
