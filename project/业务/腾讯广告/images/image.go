package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"gopkg.in/resty.v1"
	"net/http"
	"time"
)

func main() {
	type imageResp struct {
		Code      int    `json:"code"`
		Message   string `json:"message"`
		MessageCn string `json:"message_cn"`
		Data      struct {
			ImageId    string `json:"image_id"`
			Width      int    `json:"width"`
			Height     int    `json:"height"`
			FileSize   int    `json:"file_size"`
			Type       string `json:"type"`
			Signature  string `json:"signature"`
			PreviewUrl string `json:"preview_url"`
		} `json:"data"`
	}
	var (
		accessToken string = "77ae53631c5aa47bbf97f709fe920aa1"
		timestamp   int64  = time.Now().Unix()
		nonce       string = "qwers"

		accountId   string = fmt.Sprintf("%d", 30492333)
		upload_type string = "UPLOAD_TYPE_BYTES"
		signature   string = "d2ac5aa7833f3daf90d4b558c8e1e052"
		bytes       string
		description string = "超哥"

		basePath string = "https://api.e.qq.com/v1.1/images/add"
		imageUrl string = "https://ark-oss.bettagames.com/2023-03/d2ac5aa7833f3daf90d4b558c8e1e052.jpg"

		ylhImageResp imageResp
	)
	//now := time.Now().Unix()
	fileBytes, err := getFileBytes(imageUrl)
	if err != nil {
		fmt.Println("getFileBytes err", err)
		return
	}
	bytes = base64.StdEncoding.EncodeToString(fileBytes)
	url := fmt.Sprintf("%s?access_token=%s&timestamp=%v&nonce=%s", basePath, accessToken, timestamp, nonce)
	resp, err := resty.New().SetRetryCount(3).R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"account_id":  accountId,
			"upload_type": upload_type,
			"signature":   signature,
			"bytes":       bytes,
			"description": description,
		}).
		Post(url)
	if err != nil {
		fmt.Println("Post err", err)
		return
	}
	fmt.Println("code:", resp.StatusCode())
	err = json.Unmarshal(resp.Body(), &ylhImageResp)
	if err != nil {
		fmt.Println("Unmarshal err:", err)
		return
	}
	if resp.StatusCode() != http.StatusOK {
		fmt.Println("resp.StatusCode() != http.StatusOK")
		return
	}
	fmt.Println("ylhImageResp:", ylhImageResp)
}

func getFileBytes(netUrl string) ([]byte, error) {
	resp, err := resty.New().R().Get(netUrl)
	if err != nil {
		return nil, err
	}
	return resp.Body(), nil
}
