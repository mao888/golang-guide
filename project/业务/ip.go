package main

import (
	"context"
	"fmt"
	glog "github.com/mao888/mao-glog"
	"gopkg.in/resty.v1"
	"sync"
	"sync/atomic"
)

type GeoIPResponse struct {
	Country struct {
		ISOCode string `json:"iso_code"`
	} `json:"country"`
}

func main() {
	var (
		ctx        = context.Background()
		ipApi      = "http://ip-api.com/json/47.93.20.204"
		requestUrl = fmt.Sprintf("%s?lang=zh-CN", ipApi)
	)

	glog.Infof(ctx, "requestUrl: %v", requestUrl)

	var wg sync.WaitGroup
	var counter int32

	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			resp, err := resty.New().SetRetryCount(3).R().Get(requestUrl)
			if err != nil {
				glog.Errorf(ctx, "resty.New().SetRetryCount(3).R().Get(ipApi) failed, err: %v", err)
				return
			}
			glog.Infof(ctx, "第几次请求: %d; resp: %v", i, resp)
			atomic.AddInt32(&counter, 1)
		}(i)
	}

	wg.Wait()

	glog.Infof(ctx, "Total requests completed: %d", atomic.LoadInt32(&counter))
}
