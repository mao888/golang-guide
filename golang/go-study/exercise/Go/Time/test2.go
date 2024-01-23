package main

import (
	"fmt"
	"github.com/mao888/mao-gutils/constants"
	"time"
)

func main() {
	isHour := true
	data := "2023-11-06 18:00:00"
	//data := "2023-11-06"
	loa := time.FixedZone("Asia/Shanghai", 8*60*60)
	var sdate, edate int64
	if isHour {
		date, err := time.ParseInLocation(constants.TimeYMDH, data, loa)
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		sdate = date.Unix()
		edate = date.Add(time.Minute * 60).Unix() // sdate: 1699264800 edate: 1699268400
	} else {
		date, err := time.ParseInLocation(constants.TimeYMD, data, loa)
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		sdate = date.Unix()
		edate = date.Add(time.Hour * 24).Unix() // sdate: 1699200000 edate: 1699286400
	}
	fmt.Println("sdate:", sdate, "edate:", edate)
}
