package main

import (
	"encoding/json"
	"fmt"
	gutil "github.com/mao888/mao-gutils/http"
	"time"
)

type SearchVideo struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Data    struct {
		PageInfo struct {
			TotalNumber int `json:"total_number"`
			Page        int `json:"page"`
			PageSize    int `json:"page_size"`
			TotalPage   int `json:"total_page"`
		} `json:"page_info"`
		List []struct {
			VideoCoverUrl     string    `json:"video_cover_url"`
			Format            string    `json:"format"`
			PreviewUrl        string    `json:"preview_url"`
			FileName          string    `json:"file_name"`
			Displayable       bool      `json:"displayable"`
			Height            int       `json:"height"`
			Width             int       `json:"width"`
			BitRate           int       `json:"bit_rate"`
			CreateTime        time.Time `json:"create_time"`
			ModifyTime        time.Time `json:"modify_time"`
			Signature         string    `json:"signature"`
			Duration          float64   `json:"duration"`
			VideoId           string    `json:"video_id"`
			MaterialId        string    `json:"material_id"`
			AllowedPlacements []string  `json:"allowed_placements"`
			AllowDownload     bool      `json:"allow_download"`
			Size              int       `json:"size"`
		} `json:"list"`
	} `json:"data"`
	RequestId string `json:"request_id"`
}

func main() {

	var (
		searchVideoBaseUrl = "https://business-api.tiktok.com/open_api/v1.3/file/video/ad/search/"
		accessToken        = "be819b0981e021fd4b1914a4b769c6b320f54938"
		advertiserId       = "7306422754532868097"
		//advertiserId       = "7306422822664994817"
	)

	data := struct {
		AdvertiserID string `json:"advertiser_id"`
		Filtering    struct {
			VideoIDs []string `json:"video_ids"`
		} `json:"filtering"`
	}{
		AdvertiserID: advertiserId,
		Filtering: struct {
			VideoIDs []string `json:"video_ids"`
		}{
			VideoIDs: []string{"v10033g50000clnuao7og65od7tl6he0"},
		},
	}
	body, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Marshal err", err)
	}
	code, resp, err := gutil.HttpGetJson(searchVideoBaseUrl, body, map[string][]string{
		"Content-Type": {"application/json"},
		"Access-Token": {accessToken},
	})
	if err != nil {
		fmt.Println("Get err", err)
		return
	}
	fmt.Printf("code:%d\n", code)

	var searchVideo SearchVideo
	err = json.Unmarshal(resp, &searchVideo)
	if err != nil {
		fmt.Println("Unmarshal err", err)
		return
	}

	fmt.Printf("searchVideo:%+v\n", searchVideo)
}
