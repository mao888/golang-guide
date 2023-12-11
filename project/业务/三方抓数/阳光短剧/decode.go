package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	"gopkg.in/resty.v1"
	"net/url"
	"strings"
)

type SunShineOrderCallBack struct {
	Id           int     `json:"id"`
	Ver          string  `json:"ver"`
	OutTradeNo   string  `json:"outTradeNo"`
	Discount     float64 `json:"discount"`
	Type         int     `json:"type"`
	StatusNotify int     `json:"statusNotify"`
	Ctime        int     `json:"ctime"`
	FinishTime   int     `json:"finishTime"`
	UserId       int     `json:"userId"`
	ChannelId    string  `json:"channelId"`
	Domain       int     `json:"domain"`
	SourceInfo   string  `json:"sourceInfo"`
	ChapterId    string  `json:"chapterId"`
	SourceDesc   string  `json:"sourceDesc"`
	RegisterDate string  `json:"registerDate"`
	OpenId       string  `json:"openId"`
	Os           string  `json:"os"`
	ReferralId   string  `json:"referralId"`
	Adid         string  `json:"adid"`
	FromDrId     string  `json:"fromDrId"`
	Platform     string  `json:"platform"`
	Scene        string  `json:"scene"`
	ThirdCorpId  string  `json:"thirdCorpId"`
	ThirdWxId    string  `json:"thirdWxId"`
	KdrId        string  `json:"kdrId"`
	SelfReturn   string  `json:"selfReturn"`
	ProjectId    string  `json:"projectId"`
	PromotionId  string  `json:"promotionId"`
	SchannelTime string  `json:"schannelTime"`

	DyeTime      string `json:"dyeTime"`
	XuniPay      string `json:"xuniPay"`
	MoneyBenefit string `json:"moneyBenefit"`

	ReferralName string `json:"referralName"`
	AppId        string `json:"appId"`

	Date       string `json:"date"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
	Version    int64  `json:"version"`
}

func main() {
	// URL字符串
	// user
	//urlStr := "https%3A%2F%2Fxcxoss.dzyds.com%2Fdownload%2Fchannel%2F20231102%2F10007348_alias_cpsvideo_user_1698918329171.txt%3FtaskId%3DOnNLbXpKRQLcqgN&dataType=alias_user"
	// order
	//urlStr := "https%3A%2F%2Fxcxoss.dzyds.com%2Fdownload%2Fchannel%2F20231102%2F10007348_alias_cpsvideo_order_1698920405212.txt%3FtaskId%3DO4jM0nylczk118F&dataType=alias_order"
	//urlStr := "https%3A%2F%2Fxcxoss.dzyds.com%2Fdownload%2Fchannel%2F20231103%2F10007348_alias_cpsvideo_order_1698977301982.txt%3FtaskId%3DjdM86excIXp8IZ9&dataType=alias_order"
	//urlStr := "https%3A%2F%2Fxcxoss.dzyds.com%2Fdownload%2Fchannel%2F20231103%2F10007348_alias_cpsvideo_order_1698978256361.txt%3FtaskId%3DLiMSF92VqFRAkU9&dataType=alias_order"
	//urlStr := "https%3A%2F%2Fxcxoss.dzyds.com%2Fdownload%2Fchannel%2F20231106%2F10007348_alias_cpsvideo_user_1699255364687.txt%3FtaskId%3DRoxEM3IolVybj0A&dataType=alias_user"
	//urlStr := "https%3A%2F%2Fxcxoss.dzyds.com%2Fdownload%2Fchannel%2F20231106%2F10007348_alias_cpsvideo_user_1699265409934.txt%3FtaskId%3DC3SNHmrrnvVpJNN&dataType=alias_user"
	//urlStr := "https%3A%2F%2Fxcxoss.dzyds.com%2Fdownload%2Fchannel%2F20231107%2F10007348_alias_cpsvideo_order_1699324811239.txt%3FtaskId%3DcTHeJjN5FojmWt3&dataType=alias_order"

	// 生产 user
	// 2023-10-07 0:00:00 - 2023-10-31 23:59:59
	//urlStr := "https%3A%2F%2Fxcxoss.dzyds.com%2Fdownload%2Fchannel%2F20231114%2F10007348_alias_cpsvideo_user_1699950504382.txt%3FtaskId%3DRI6brZ25kMff8aR&dataType=alias_user"
	// 2023-11-01 0:00:00 - 2023-11-13 16:59:59
	//urlStr := "https%3A%2F%2Fxcxoss.dzyds.com%2Fdownload%2Fchannel%2F20231114%2F10007348_alias_cpsvideo_user_1699950576068.txt%3FtaskId%3DxAz68fhtefUCU8s&dataType=alias_user"
	// 生产 order
	// 2023-10-07 0:00:00 - 2023-10-31 23:59:59
	//urlStr := "https%3A%2F%2Fxcxoss.dzyds.com%2Fdownload%2Fchannel%2F20231114%2F10007348_alias_cpsvideo_order_1699950606540.txt%3FtaskId%3DLAie8XsAo1FPQGM&dataType=alias_order"
	// 2023-11-01 0:00:00 - 2023-11-13 16:59:59
	//urlStr := "https%3A%2F%2Fxcxoss.dzyds.com%2Fdownload%2Fchannel%2F20231114%2F10007348_alias_cpsvideo_order_1699950654773.txt%3FtaskId%3DpwffF360k4ldDan&dataType=alias_order"
	urlStr := "https://xcxoss.dzyds.com/download/channel/20231211/10007348_alias_cpsvideo_order_1702260601725.txt"

	// 对URL进行解码
	decodedURL, err := url.QueryUnescape(urlStr)
	if err != nil {
		fmt.Println("解码失败:", err)
		return
	}
	// 打印解码后的URL
	fmt.Println("解码后的URL:", decodedURL)
	// user 解码后的URL: https://xcxoss.dzyds.com/download/channel/20231102/10007348_alias_cpsvideo_user_1698918329171.txt?taskId=OnNLbXpKRQLcqgN&dataType=alias_user
	// order 解码后的URL: https://xcxoss.dzyds.com/download/channel/20231102/10007348_alias_cpsvideo_order_1698920405212.txt?taskId=O4jM0nylczk118F&dataType=alias_order
	// order 解码后的URL: https://xcxoss.dzyds.com/download/channel/20231103/10007348_alias_cpsvideo_order_1698977301982.txt?taskId=jdM86excIXp8IZ9&dataType=alias_order
	// order 解码后的URL: https://xcxoss.dzyds.com/download/channel/20231103/10007348_alias_cpsvideo_order_1698978256361.txt?taskId=LiMSF92VqFRAkU9&dataType=alias_order
	//
	// user 解码后的URL: https://xcxoss.dzyds.com/download/channel/20231106/10007348_alias_cpsvideo_user_1699265409934.txt?taskId=C3SNHmrrnvVpJNN&dataType=alias_user
	// 生产 user
	// 2023-10-07 0:00:00 - 2023-10-31 23:59:59	 https://xcxoss.dzyds.com/download/channel/20231114/10007348_alias_cpsvideo_user_1699950504382.txt?taskId=RI6brZ25kMff8aR&dataType=alias_user
	// 2023-11-01 0:00:00 - 2023-11-13 16:59:59	 https://xcxoss.dzyds.com/download/channel/20231114/10007348_alias_cpsvideo_user_1699950576068.txt?taskId=xAz68fhtefUCU8s&dataType=alias_user
	// 生产 order
	// 2023-10-07 0:00:00 - 2023-10-31 23:59:59	 https://xcxoss.dzyds.com/download/channel/20231114/10007348_alias_cpsvideo_order_1699950606540.txt?taskId=LAie8XsAo1FPQGM&dataType=alias_order
	// 2023-11-01 0:00:00 - 2023-11-13 16:59:59	 https://xcxoss.dzyds.com/download/channel/20231114/10007348_alias_cpsvideo_order_1699950654773.txt?taskId=pwffF360k4ldDan&dataType=alias_order

	resp, err := resty.New().SetRetryCount(3).R().
		SetHeader("Content-Type", "application/json").
		Get(decodedURL)
	if err != nil {
		fmt.Println("Get err", err)
		return
	}
	fmt.Println("resp:", string(resp.Body()))

	strResp := string(resp.Body())

	// 分割 resp 字符串为每个 JSON 对象
	jsonObjects := strings.Split(strResp, "\n")
	glog.Infof("jsonObjects:%+v", jsonObjects)

	// 创建一个存储 JSON 对象的切片
	var jsonArray []SunShineOrderCallBack

	// 遍历每个 JSON 对象，解析为 map 并添加到切片中
	for _, jsonObj := range jsonObjects {
		if jsonObj == "" {
			continue
		}

		var data SunShineOrderCallBack
		err := json.Unmarshal([]byte(jsonObj), &data)
		if err != nil {
			glog.Errorf("json.Unmarshal err:%+v", err)
		}
		jsonArray = append(jsonArray, data)
		//if err := json.Unmarshal([]byte(jsonObj), &data); err == nil {
		//	jsonArray = append(jsonArray, data)
		//}
	}

	fmt.Println("=====================================")
	glog.Infof("jsonArray:%+v", jsonArray)

	// 打印 JSON 数组
	//jsonArrayStr, _ := json.Marshal(jsonArray)
	//fmt.Println(string(jsonArrayStr))

	//fmt.Println("resp:", resp)

}
