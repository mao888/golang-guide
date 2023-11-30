package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/resty.v1"
	"time"
)

type (
	TikTokUploadVideoRes struct {
		Message   string                  `json:"message"`    // 返回信息 例如：OK
		Code      int                     `json:"code"`       // 返回码 例如：0
		Data      []TikTokUploadVideoItem `json:"data"`       // 返回数据。由于兼容性原因，会返回数组而不是对象，且数组中仅有一个对象。
		RequestId string                  `json:"request_id"` // 请求的日志id，唯一标识一个请求
	}
	TikTokUploadVideoItem struct {
		VideoCoverUrl        string    `json:"video_cover_url"`         // 视频封面临时URL
		Format               string    `json:"format"`                  // 视频格式
		PreviewUrl           string    `json:"preview_url"`             // 视频预览链接
		PreviewUrlExpireTime string    `json:"preview_url_expire_time"` // 视频预览链接过期时间
		FileName             string    `json:"file_name"`               // 视频名称
		Displayable          bool      `json:"displayable"`             // 视频能否在平台中展示
		Height               int       `json:"height"`                  // 视频高度
		Width                int       `json:"width"`                   // 视频宽度
		BitRate              int       `json:"bit_rate"`                // 码率，单位bps
		CreateTime           time.Time `json:"create_time"`             // 创建时间。UTC 时间，格式：2020-06-10T07:39:14Z
		ModifyTime           time.Time `json:"modify_time"`             // 修改时间。UTC 时间，格式：2020-06-10T07:39:14Z
		Signature            string    `json:"signature"`               // 视频文件MD5
		Duration             float64   `json:"duration"`                // 视频时长，单位秒
		VideoId              string    `json:"video_id"`                // 视频ID，可用于广告投放中创建广告
		Size                 int       `json:"size"`                    // 视频大小，单位Byte
		MaterialId           string    `json:"material_id"`             // 素材ID
		AllowedPlacements    []string  `json:"allowed_placements"`      // 视频可投放版位
		AllowDownload        bool      `json:"allow_download"`          // 视频是否允许下载
		FixTaskId            string    `json:"fix_task_id"`             // 修复任务ID。仅在请求中flaw_detect和 auto_fix_enabled都设置为True，且检测到视频中的问题时返回。
		FlawTypes            []string  `json:"flaw_types"`              // 视频问题种类。 仅在请求中flaw_detect和 auto_fix_enabled都设置为True，且检测到视频中的问题时返回。
	}
)

func main() {
	var (
		uploadVideoBaseUrl = "https://business-api.tiktok.com/open_api/v1.3/file/video/ad/upload/"
		accessToken        = ""
		advertiserId       = ""
		fileName           = ""
		uploadType         = "UPLOAD_BY_URL"
		videoUrl           = ""
	)

	resp, err := resty.SetRetryCount(3).R().SetHeaders(map[string]string{
		"Content-Type": "application/json",
		"Access-Token": accessToken,
	}).SetBody(map[string]interface{}{
		"advertiser_id":     advertiserId,
		"file_name":         fileName,   // 默认值为文件名称或者URL的最后一个 路径名
		"upload_type":       uploadType, // 视频上传方式 默认值: UPLOAD_BY_FILE，枚举值: UPLOAD_BY_FILE，UPLOAD_BY_URL，UPLOAD_BY_FILE_ID， UPLOAD_BY_VIDEO_ID
		"video_url":         videoUrl,   // 视频文件的URL 当 upload_type 为UPLOAD_BY_URL 时必填
		"is_third_party":    false,      // 视频是否为第三方视频
		"flaw_detect":       false,      // 是否自动检测视频的潜在问题
		"auto_fix_enabled":  false,      // 是否自动修复检测到的问题。 默认值 : False。
		"auto_bind_enabled": false,      // 是否自动将修复后的视频上传至素材库。默认值： False。此字段只在flaw_detect和auto_fix_enabled 均设置为True时生效。
	}).Post(uploadVideoBaseUrl)
	if err != nil {
		fmt.Println("Post err", err)
		return
	}

	var videoRes TikTokUploadVideoRes
	err = json.Unmarshal(resp.Body(), &videoRes)
	if err != nil {
		fmt.Println("Unmarshal err", err)
		return
	}
	fmt.Printf("videoRes:%+v\n", videoRes)
}
