/**
    @author: HuChao
    @since: 2022/8/17
    @desc: //TODO ES查询demo
**/
package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type EsInfo struct {
	Host    string `json:"host,omitempty"`
	User    string `json:"user,omitempty"`
	Pwd     string `json:"password,omitempty"`
	Timeout string `json:"timeout,omitempty"`
}

type EsCli struct {
	Es      *elastic.Client
	TimeOut string
}

type HourCountEsResponse struct {
	Aggregations struct {
		Hourcount struct {
			Buckets []struct {
				Key          string `json:"key"`
				SumHourcount struct {
					Value float64 `json:"value"`
				} `json:"sum_hour_count"`
			} `json:"buckets"`
		} `json:"hour_count"`
	} `json:"aggregations"`
}

var queryStr = `
	{
		"aggs": {
			  "hour_count": {
				  "aggs": {
					  "sum_hour_count": {
						  "sum": {
							  "field": "hour_count"
						  }
					  }
				  },
				  "terms": {
					  "field": "project"
				  }
			  }
		  },
		  "size":0
	  }
	`

func apiEsQuery(c *gin.Context) {
	esInfo := &EsInfo{
		Host:    "**",
		User:    "**",
		Pwd:     "**",
		Timeout: "10",
	}
	var (
		cli *EsCli
		err error
	)
	//初始化es
	if cli, err = InitEs(esInfo); err != nil {
		fmt.Printf("err: %s", err.Error())
		return
	}
	//查询
	var esRawResp *elastic.Response
	if esRawResp, err = QueryES(cli, "GET", "/indexName/_search", "application/json;charset=UTF-8", queryStr); err != nil {
		fmt.Printf("err: %s", err.Error())
		return
	}
	if esRawResp.StatusCode != 200 {
		fmt.Printf("err: es response code %d", esRawResp.StatusCode)
		return
	}
	//aggregations 解析结果
	esResp := &HourCountEsResponse{}
	if err = json.Unmarshal(esRawResp.Body, esResp); err != nil {
		fmt.Printf("err: %s", err.Error())
		return
	}

	fmt.Printf("esResp:%+#v", esResp.Aggregations.Hourcount.Buckets)
	res := make(map[string]interface{})
	for _, v := range esResp.Aggregations.Hourcount.Buckets {
		res[v.Key] = v.SumHourcount.Value
	}
	fmt.Printf("esResp:%+#v", res)
}

func QueryES(cli *EsCli, method, path, contentType string, body interface{}) (*elastic.Response, error) {
	reqopts := elastic.PerformRequestOptions{
		Method:      method,
		Path:        path, // build url
		Body:        body,
		ContentType: contentType,
	}
	timeout := "20"

	to, err := time.ParseDuration(timeout)
	if err != nil {
		fmt.Printf("time duration: %s", err.Error())
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), to)
	defer cancel()

	return cli.Es.PerformRequest(ctx, reqopts)
}

func InitEs(esInfo *EsInfo) (*EsCli, error) {
	to, err := time.ParseDuration(esInfo.Timeout)
	if err != nil {
		fmt.Printf("time duration: %s", err.Error())
		return nil, err
	}

	httpCli := &http.Client{
		Timeout: to,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
			TLSHandshakeTimeout: 10 * time.Second,
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 100,
		},
	}

	client, err := elastic.NewClient(elastic.SetHttpClient(httpCli),
		elastic.SetURL(esInfo.Host),
		elastic.SetSniff(false), // disable sniffing
		elastic.SetBasicAuth(esInfo.User, esInfo.Pwd),
	)
	if err != nil {
		fmt.Printf("NewClient %s", err.Error())
		return nil, err
	}

	ctx := context.Background()
	_, _, err = client.Ping(esInfo.Host).Do(ctx)
	if err != nil {
		fmt.Printf("Ping %s", err.Error())
		return nil, err
	}

	return &EsCli{
		Es:      client,
		TimeOut: esInfo.Timeout,
	}, nil
}
