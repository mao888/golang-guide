package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gopkg.in/resty.v1"
	"net/http"
	"time"
)

func main() {
	type videoResp struct {
		Code      int    `json:"code"`
		Message   string `json:"message"`
		MessageCn string `json:"message_cn"`
		Data      struct {
			VideoId int64 `json:"video_id"`
		} `json:"data"`
	}

	var (
		accessToken  string = "77ae53631c5aa47bbf97f709fe920aa1"
		basePath     string = "https://api.e.qq.com/v1.1/videos/add"
		videoUrl     string = "https://ark-oss.bettagames.com/2023-03/62044feaac661c98e56567a80e4dcb40.mp4"
		accountId    int64  = 30492333
		ylhVideoResp videoResp
		signature    string = "62044feaac661c98e56567a80e4dcb40"
		description  string = "超哥哥"
		nonce        string = "qwertyuiopwwsddddsq"
	)
	now := time.Now().Unix()
	fileBytes, err := getFileBytes2(videoUrl)
	if err != nil {
		fmt.Println("getFileBytes err", err)
		return
	}
	url := fmt.Sprintf("%s?access_token=%s&timestamp=%v&nonce=%s", basePath, accessToken, now, nonce)
	resp, err := resty.New().SetRetryCount(3).R().
		SetFileReader("video_file", "asset.mp4", bytes.NewReader(fileBytes)).
		SetFormData(map[string]string{
			"account_id":  fmt.Sprintf("%d", accountId),
			"signature":   signature,
			"description": description,
		}).
		Post(url)
	if err != nil {
		fmt.Println("Post err", err)
		return
	}
	fmt.Println("code:", resp.StatusCode())
	err = json.Unmarshal(resp.Body(), &ylhVideoResp)
	if err != nil {
		fmt.Println("Unmarshal err:", err)
		return
	}
	if resp.StatusCode() != http.StatusOK {
		fmt.Println("resp.StatusCode() != http.StatusOK")
		return
	}
	fmt.Println("ylhVideoResp: ", ylhVideoResp)
}

func getFileBytes2(netUrl string) ([]byte, error) {
	resp, err := resty.New().R().Get(netUrl)
	if err != nil {
		return nil, err
	}
	return resp.Body(), nil
}
