package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"gopkg.in/resty.v1"
	"strings"
	"time"
)

type Order struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Data    string `json:"data"`
}

func main() {

	clientId := "10007348"
	// 当前时间戳(毫秒)
	timestamp := time.Now().UnixNano() / 1e6
	token := "AGy97zKlbzEArFwNDh"
	//签名秘钥(小写md5(clientId+接口token+timestamp))
	signKey := strings.ToLower(fmt.Sprintf("%x", md5.Sum([]byte(clientId+token+fmt.Sprintf("%v", timestamp)))))

	url := "https://routine.wqxsw.com/flames/channel/query/order"

	// 2023-10-07 0:00:00	1696608000
	// 2023-10-31 23:59:59	1698767999
	// 2023-11-01 0:00:00	1698768000
	// 2023-11-13 16:59:59	1699865999

	resp, err := resty.New().SetRetryCount(3).R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"clientId":   clientId,     //clientId
			"sdate":      "1698768000", //开始时间-时间戳(秒)
			"edate":      "1699865999", //结束时间-时间戳(秒)
			"queryField": "loginTime",  //支持自定义字段查询：此字段不为空使用loginTime，否则使用ctime.loginTime[用户活跃时间]-满足每日新增、每日活跃数据统计;ctime[用户注册时间]
			"channelCodeList": []string{
				"29576",
				"29577",
				"29578",
				"29579",
				"29580",
				"33343",
			}, //渠道ID - 不传时，查询clientId下所有授权的渠道ID
			"timestamp": timestamp, //时间戳(毫秒)
			"signKey":   signKey,   //签名秘钥(小写md5(clientId+接口token+timestamp))
		}).
		Post(url)
	if err != nil {
		fmt.Println("Post err", err)
		return
	}

	var order Order
	err = json.Unmarshal(resp.Body(), &order)
	if err != nil {
		fmt.Println("Unmarshal err:", err)
		return
	}
	fmt.Printf("order:%+v\n", order)
	// order:{RetCode:0 RetMsg:成功 Data:O4jM0nylczk118F}
	// order:{RetCode:0 RetMsg:成功 Data:jdM86excIXp8IZ9}
	// order:{RetCode:0 RetMsg:成功 Data:LiMSF92VqFRAkU9}
	// 生产 order
	// 2023-10-07 0:00:00 - 2023-10-31 23:59:59		order:{RetCode:0 RetMsg:成功 Data:LAie8XsAo1FPQGM}
	// 2023-11-01 0:00:00 - 2023-11-13 16:59:59		order:{RetCode:0 RetMsg:成功 Data:pwffF360k4ldDan}
}
