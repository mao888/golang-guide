package main

import (
	"context"
	"encoding/csv"
	"fmt"
	glog "github.com/mao888/mao-glog"
	"gopkg.in/resty.v1"
	"io"
	"strings"
)

type MintegralRes struct {
	Date       string `json:"date"`
	OfferId    string `json:"offer_id"`
	Uuid       string `json:"uuid"`
	OfferName  string `json:"offer_name"`
	CampaignId string `json:"campaign_id"`
	SubId      string `json:"sub_id"`
	Package    string `json:"package"`
	Location   string `json:"location"`
	Currency   string `json:"currency"`
	Impression string `json:"impression"`
	Click      string `json:"click"`
	Conversion string `json:"conversion"`
	Ecpm       string `json:"ecpm"`
	Cpc        string `json:"cpc"`
	Ctr        string `json:"ctr"`
	Cvr        string `json:"cvr"`
	Ivr        string `json:"ivr"`
	Spend      string `json:"spend"`
}

func main() {

	var (
		ctx          = context.Background()
		mintegralUrl = "https://ss-api.mintegral.com/api/v2/reports/data"

		startTime       = "2024-05-01"
		endTime         = "2024-05-01"
		timezone        = "+0"
		dimensionOption = "Offer,Campaign,Sub,Location,Package"
		typeValue       = "2"
	)
	// https://ss-api.mintegral.com/api/v2/reports/data?start_time=2024-05-01&end_time=2024-05-01&timezone=+0&dimension_option=Offer,Campaign,Sub,Location,Package&type=2
	requestUrl := fmt.Sprintf("%s?start_time=%s&end_time=%s&timezone=%s&dimension_option=%s&type=%s",
		mintegralUrl, startTime, endTime, timezone, dimensionOption, typeValue)
	resp, err := resty.New().SetRetryCount(3).R().
		SetHeaders(map[string]string{
			"token":     "39cd9d8bc2b212623cb1bbefe78b114a",
			"username":  "Fotoable_MTG",
			"timestamp": "1715329009",
		}).
		Get(requestUrl)
	if err != nil {
		glog.Errorf(ctx, "Post err:%s", err)
		return
	}

	// 读取resp.Body()中的数据,为csv文件,然后解析csv文件 生成[]MintegralRes
	// 解析CSV数据

	mintegralList, err := parseCSVAll(strings.NewReader(resp.String()))
	if err != nil {
		glog.Errorf(ctx, "parseCSVAll err:%s", err)
		return
	}

	// 打印解析后的数组
	glog.Infof(ctx, "mintegralList:%+v", mintegralList)
}

func parseCSVAll(csvData io.Reader) ([]MintegralRes, error) {
	reader := csv.NewReader(csvData)
	reader.Comma = '\t'

	// 读取CSV文件中的所有记录
	records, err := reader.ReadAll()
	if err != nil {
		glog.Errorf(context.Background(), "error reading CSV: %v", err)
		return nil, err
	}

	// 创建一个存储 JobDataRes 的切片
	var m []MintegralRes

	var (
		sum = 0
	)

	// 遍历CSV记录并将其映射到 MintegralRes 结构体
	for i, record := range records {
		// 跳过标题行
		if i == 0 {
			continue
		}
		mintegral := MintegralRes{
			Date:       record[0],
			OfferId:    record[1],
			Uuid:       record[2],
			OfferName:  record[3],
			CampaignId: record[4],
			SubId:      record[5],
			Package:    record[6],
			Location:   record[7],
			Currency:   record[8],
			Impression: record[9],
			Click:      record[10],
			Conversion: record[11],
			Ecpm:       record[12],
			Cpc:        record[13],
			Ctr:        record[14],
			Cvr:        record[15],
			Ivr:        record[16],
			Spend:      record[17],
		}
		m = append(m, mintegral)
		sum++
	}
	glog.Infof(context.Background(), "总共:%d条; 实际:%d", len(records)-1, len(m))
	return m, nil
}
