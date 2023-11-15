package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"gopkg.in/resty.v1"
	"strings"
	"time"
)

type ChannelInfoList struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Data    struct {
		ChannelList     []string `json:"channelList"`
		ChannelInfoList []struct {
			ChannelId string `json:"channelId"`
			UserName  string `json:"userName"`
			NickName  string `json:"nickName"`
		} `json:"channelInfoList"`
	} `json:"data"`
}

func main() {

	clientId := "10007348"
	//当前时间戳(毫秒)
	timestamp := time.Now().UnixNano() / 1e6
	token := "AGy97zKlbzEArFwNDh"
	//签名秘钥(小写md5(clientId+接口token+timestamp))
	signKey := strings.ToLower(fmt.Sprintf("%x", md5.Sum([]byte(clientId+token+fmt.Sprintf("%v", timestamp)))))

	url := "https://routine.wqxsw.com/flames/channel/query/channelInfoList"
	resp, err := resty.New().SetRetryCount(3).R().
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"clientId":  clientId,  //clientId
			"timestamp": timestamp, //时间戳(毫秒)
			"signKey":   signKey,   //签名秘钥(小写md5(clientId+接口token+timestamp))
		}).
		Post(url)
	if err != nil {
		fmt.Println("Post err", err)
		return
	}

	var channelInfoList ChannelInfoList
	err = json.Unmarshal(resp.Body(), &channelInfoList)
	if err != nil {
		fmt.Println("Unmarshal err:", err)
		return
	}
	fmt.Printf("channelInfoList:%+v\n", channelInfoList)
	// channelInfoList:{RetCode:0 RetMsg:成功 Data:{ChannelList:[29576 29577 29578 29579 29580] ChannelInfoList:[{ChannelId:29576 UserName:beitadj1 NickName:无限短剧-头条直投1-贝塔} {ChannelId:29577 UserName:beitadj2 NickName:无限短剧-头lId:29578 UserName:beitadj3 NickName:无限短剧-头条直投3-贝塔} {ChannelId:29579 UserName:beitadj4 NickName:无限短剧-头条直投4-贝塔} {ChannelId:29580 UserName:beitadj5 NickName:无限短剧-头条直投5-贝塔}]}}
	// channelInfoList:{RetCode:0 RetMsg:成功 Data:{ChannelList:[29576 29577 29578 29579 29580 33343] ChannelInfoList:[{ChannelId:29576 UserName:beitadj1 NickName:无限短剧-头条直投1-贝塔} {ChannelId:29577 UserName:beitadj2 NickName:无限ChannelId:29578 UserName:beitadj3 NickName:无限短剧-头条直投3-贝塔} {ChannelId:29579 UserName:beitadj4 NickName:无限短剧-头条直投4-贝塔} {ChannelId:29580 UserName:beitadj5 NickName:无限短剧-头条直投5-贝塔} {ChannelId:33343 UserNa6-贝塔}]}}
}
