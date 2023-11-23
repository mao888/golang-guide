package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	"gopkg.in/resty.v1"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

type (
	WeChatForwardResp struct {
		Data    string `json:"data"`
		Message string `json:"message"`
		Status  int    `json:"status"`
	}
)

const WeChatMiniGameForwardUrl = "https://clue.oceanengine.com/outer/wechat/applet/token/"

func main() {
	var (
		token     = "52F118BF3C8596B44E31BFD45D4B15B7"
		timestamp = time.Now().Unix()
		urlToken  = "1760312311655437"

		clickId = "1779451543470157,1783156862460969"
		openId  = "oso8q5BZxVv5QqKiJHQirpFCZ7GA"
	)

	rand.Seed(time.Now().UnixNano())
	nonce := strconv.Itoa(rand.Intn(100000000))

	s := signature(token, fmt.Sprintf("%d", timestamp), nonce)

	url := fmt.Sprintf("%s%s?timestamp=%d&nonce=%s&signature=%s",
		WeChatMiniGameForwardUrl, urlToken, timestamp, nonce, s)
	fmt.Printf("weChatMiniActiveEvent url: %s", url)

	resp, err := resty.New().SetRetryCount(3).R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"clue_token": clickId, // 用户跳转小程序时携带的clue_token参数
			"open_id":    openId,  // 微信open_id，必须回传
			"event_type": "2",     // 事件类型，string格式
			"props": map[string]interface{}{
				"pay_amount": 1,
			},
		}).
		Post(url)
	if err != nil {
		glog.Errorf("weChatMiniActiveEvent resty.New().R().Post err: %v", err)
		return
	}
	fmt.Println("resp:", string(resp.Body()))

	var weChatResp WeChatForwardResp
	err = json.Unmarshal(resp.Body(), &weChatResp)
	if err != nil {
		glog.Errorf("weChatMiniActiveEvent json.Unmarshal err: %v", err)
	}
	fmt.Printf("weChatMiniActiveEvent weChatResp: %+v", weChatResp)
}

func signature(token string, timestamp string, nonce string) string {
	strList := []string{token, timestamp, nonce}
	sort.Strings(strList)
	var buffer bytes.Buffer
	for _, str := range strList {
		buffer.WriteString(str)
	}
	h := sha1.New()
	h.Write(buffer.Bytes())
	r := fmt.Sprintf("%x", h.Sum(nil))
	return r
}
