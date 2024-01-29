package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/mao888/mao-glog"
	"gopkg.in/resty.v1"
)

type (
	tokenRes struct {
		AccessToken  string `json:"access_token"`
		TokenType    string `json:"token_type"`
		RefreshToken string `json:"refresh_token"`
		ExpiresIn    int    `json:"expires_in"`
		Scope        string `json:"scope"`
		Jti          string `json:"jti"`
	}
)

var (
	clientID     = "2401244203476908609792"
	secret       = "chv65g99dlw9"
	grantType    = "refresh_token"
	refreshToken = "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IlA0SkRvOUx1emlkeUVoZlF6SG1qbHRKR0R5bC81UXNKVzBZMDhkTmkrTHM9In0.eyJhdWQiOlsib3BlbmFwaSJdLCJ1c2VyX25hbWUiOiJhZG1pbiIsInNjb3BlIjpbInJlYWQiXSwiYXRpIjoiYzIwM2M1NzItYjU5Zi00ZDA5LWE1MjAtYWY4MTVmNTYzMWIwIiwiZXhwIjoxNzA5MTA2NDA5LCJhdXRob3JpdGllcyI6WyJST0xFX0FETUlOIl0sImp0aSI6ImI0NzkyNjUwLWM2NzEtNDU3Zi1hYjdhLWQyYTUyMzg0MzJiYSIsImNsaWVudF9pZCI6IjI0MDEyNDQyMDM0NzY5MDg2MDk3OTIifQ.ORs04KymxUoFi9RIJabUIQbdYV64u6ZfdLLNNChyFBquo6hoCLrPpknFaJ3MHcg2B6ckKXL2wxXu121x21QwTp2rR1mM5HoTK-x94c4L0xhIc6oojRoCanotSaWEWiBiGY_RzYYBne5sT-zrUp8pDjdlnNZRFNlEvTj9hPHKNWb0P-uwu2sIfDKelH5rEnjVzkWAcZ0Ivj6yLKgdOPxcLfIj_OIyUYzdWuDY5KqIlu0pWGTrFUvR1OqWIjYu6gSRQMH0NBUKv2xDdDrQSkujMC1FzL7L_qJKWdrone3bpMJuJY6xknMyOIMtl7FSuG9lbao5NsZIctEfkR41gjXrBg"
)

func main() {
	ctx := context.Background()
	// 拼接client_id和secret，并转换为字节数组
	data := []byte(clientID + ":" + secret)
	// 使用base64进行编码
	encoded := base64.StdEncoding.EncodeToString(data)
	authorization := "Basic " + encoded

	url := fmt.Sprintf("https://api.adsbigo.com/oauth/token?grant_type=%s&refresh_token=%s",
		grantType, refreshToken)
	glog.Infof(ctx, "url:%s", url)

	resp, err := resty.New().SetRetryCount(3).R().
		SetHeaders(map[string]string{
			"Content-Type":  "application/json",
			"Authorization": authorization,
		}).Post(url)
	if err != nil {
		glog.Errorf(ctx, "Post err:%s", err)
		return
	}
	glog.Infof(ctx, "resp:%s", string(resp.Body()))

	var token tokenRes
	err = json.Unmarshal(resp.Body(), &token)
	if err != nil {
		glog.Errorf(ctx, "Unmarshal err:%s", err)
		return
	}
	glog.Infof(ctx, "token:%+v", token)
}
