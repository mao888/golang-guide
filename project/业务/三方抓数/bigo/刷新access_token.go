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

	// {"access_token":"eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IlA0SkRvOUx1emlkeUVoZlF6SG1qbHRKR0R5bC81UXNKVzBZMDhkTmkrTHM9In0.eyJhdWQiOlsib3BlbmFwaSJdLCJ1c2VyX25hbWUiOiJhZG1pbiIsInNjb3BlIjpbInJlYWQiXSwiZXhwIjoxNzA2NjAzNzE0LCJhdXRob3JpdGllcyI6WyJST0xFX0FETUlOIl0sImp0aSI6IjhkOGI4NWMwLTZjYWMtNDhiZS1hYWU4LTVjN2JhZTc5OTllNSIsImNsaWVudF9pZCI6IjI0MDEyNDQyMDM0NzY5MDg2MDk3OTIifQ.XO4EX_2MOG4Ey2QbDPaY2CKRlB7jihJuXTu4rpeFXUkv6m4nFyhuLknY9R9sgdGVnfrLyva2ZJ8EoK5GS0qgrgAdBWSCe7P0sXxMFyzMn4o_0XUbZyINe2bmvZKNF587lfa-lOPb8JkY_vR3GrTpnC1FaeMi7Mdve2INBup4myNTDa3Jq3w8RJ1ops9Q7CpTAz7gah2vYs4IIztKB_q9hpGJXDgfSgmtHLVNacPM4NSQ0LGyfuKDe4uMgQGGCk-MsY8CdlfBNZplcOpPIfA4_J4z6-YjNbpedGVZCkcy6mK0U5wWL-yd8USTrkFic9oHabicDBRVDpqtsXrxxoTTLA","token_type":"bearer","refresh_token":"eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCIsImtpZCI6IlA0SkRvOUx1emlkeUVoZlF6SG1qbHRKR0R5bC81UXNKVzBZMDhkTmkrTHM9In0.eyJhdWQiOlsib3BlbmFwaSJdLCJ1c2VyX25hbWUiOiJhZG1pbiIsInNjb3BlIjpbInJlYWQiXSwiYXRpIjoiOGQ4Yjg1YzAtNmNhYy00OGJlLWFhZTgtNWM3YmFlNzk5OWU1IiwiZXhwIjoxNzA5MTA5MzE0LCJhdXRob3JpdGllcyI6WyJST0xFX0FETUlOIl0sImp0aSI6ImQwZTZmMjc4LTgyMjctNDVjOC05MTBhLWY3ZmUwNGNmYTk0ZCIsImNsaWVudF9pZCI6IjI0MDEyNDQyMDM0NzY5MDg2MDk3OTIifQ.bdcyqGTZcxvMEYq5Zp8T_fhYicFNaTHkqvdg7oCV6Zg9h0aiE7vGeLwVr5ngItwTidVbvE1stinB591LGEwYW2O0MM0mVwQ3ht1uxBOGNIOb0Uxgnntk_Q7s4DrzEQd1b3B0PJuY6Brmcki43w8Qmtjcjib-jpSjMfvIuD2T72FWpBvjee2TniYnALZ97YS9iHtQo24Ol1Im4BiO0cPcidpwLspiXXn2iXpAetNGn7xRACBNQRuAcaIu7xXhZGnwOwT-OUSgCRnUvqLH3-6j7YvVXP91Zcd1DHfAmFJAbDe6WZ71kyaag3W_zBrQFcJkYs__3S_qzxI943xUbF4pPw","expires_in":86399,"scope":"read","jti":"8d8b85c0-6cac-48be-aae8-5c7bae7999e5"}
}
