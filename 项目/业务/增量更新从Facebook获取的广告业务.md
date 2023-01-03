说明:

广告系列 ： 广告组关系	1 : n

广告组 ： 广告关系		1 : n

# schedule-console-调度服务

## 业务逻辑

**全量、增量更新从Facebook获取的数据的定时任务**

- UpdateFBAds 增量更新facebook 所有广告数据

- - 1、获取当前开通自动化投放的账号（调用application服务获取Facebook账户ID）
  - 2、获取时间（根据账户ID 从redis获取）

- - - 根据账户ID 获取广告系列最新的更新时间 暂定全量更新，
    - 根据账户ID 获取广告组最新的更新时间
    - 根据账户ID 获取广告最新的更新时间

- - 3、全量、增量更新（根据前端传值 refresh == 0:刷新所有; refresh ==1: 只刷新广告组和广告; refresh ==2: 系列）

- - - 第一次时间为0，全量更新；后续只需要过滤出比上次最新更新时间大的数据，增量更新
    - 账户的新增**广告系列**

- - - - 请求Facebook获取数据，若出现错误，请求重试机制，重试3次
      - 处理Facebook返回的数据，取出自己需要用的，剩余的全部存到json对象里
      - 入库（mysql、redis）

- - - - - FB返回的结果集是按照时间降序排列的，获取本次操作的最大更新时间(第0个)：accCps[0].UpdatedAt
        - 入库，gorm的save方法。（save：可更新或新增）
        - 将最新时间set到redis
        - 如果出现错误，钉钉报警，发送到我们内部钉钉群，我们便可以及时查看错误
        - message = fmt.Sprintf("[业务异常][服务：%s][主机名：%s] %s", serverName, hostname, message)

- - - 账户的新增**广告组**

- - - - 与系列同理

- - - 账户的新增**广告**

- - - - 与系列同理

- - 4、更新完成之后，kafka 通知BI本天更新完成
  - gkafka.ProducerString(context.Background(), "adv_asset_notice", "ok")

## internal-cron

### cron-cron.go(启动定时任务)

```go
package cron

import (
    "context"
    "gitlab.ftsview.com/aircraft/schedule-console/internal/logic"
    "gitlab.ftsview.com/fotoable-go/gkafka"
    "gitlab.ftsview.com/fotoable-go/glog"

    "github.com/robfig/cron/v3"
)

var cronIns Cron

type Cron struct {
    cron *cron.Cron

    ctx context.Context
}

func InitCron() {
    // nyc, _ := time.LoadLocation("Asia/Shanghai")
    // c := cron.New(cron.WithLocation(nyc))
    c := cron.New()
    c.Start()
    cronIns.cron = c
    cronIns.ctx = context.Background()

    cronIns.StartUpdateFBAdData()
}

func (c Cron) StartCron(timer string, f func()) {
    _, err := c.cron.AddFunc(timer, f)
    if err != nil {
        glog.Errorf(c.ctx, "start cron error %w", err)
    }
}

func (c Cron) StartUpdateFBAdData() {
    // 测试时间5s一更新：@every 5s
    // 01 00 * * *
    c.StartCron("01 00 * * *", func() {
        result, _ := logic.SingletonFbCampaginLogic().UpdateFBAds(cronIns.ctx, nil)
        glog.Infof(cronIns.ctx, "本次更新广告系列数：%d, 广告组数: %d, 广告数: %d", result.Campagins, result.AdSets, result.Ads)
        //TODO 通知BI本天更新完成
        gkafka.ProducerString(context.Background(), "adv_asset_notice", "ok")
    })

    //c.StartCron("30 17 * * *", func() {
    //	if config.GlobConfig.Syncbi {
    //		logic.SingletonFBAdsetLogic().CromSyncBi(cronIns.ctx)
    //	}
    //})
}
```

## internal-logic

### logic-fb_campagic.go（广告系列logic）

```go
package logic

import (
	"context"
	"fmt"
	"gitlab.ftsview.com/aircraft/schedule-console/internal/constants"
	"gitlab.ftsview.com/aircraft/schedule-console/internal/handler/bean"
	"gitlab.ftsview.com/aircraft/schedule-console/internal/service"
	"gitlab.ftsview.com/fotoable-go/glog"
	"sync"
	"sync/atomic"
)

var (
	_fbCampaginOnce  sync.Once
	_fbCampaginLogic *FbCampaginLogic
)

type FbCampaginLogic struct {
	fbCampaginService *service.FbCampaginService	// 广告系列service
	ramblerService    *service.RamblerService		// 获取账号service
	fbAdSetService    *service.FbAdSetService		// 广告组service
	fbAdService       *service.FbAdService			// 广告service
}

func SingletonFbCampaginLogic() *FbCampaginLogic {
	_fbCampaginOnce.Do(func() {
		_fbCampaginLogic = &FbCampaginLogic{
			fbCampaginService: service.SingletonFbCampaginService(),
			ramblerService:    service.SingletonRamblerService(),
			fbAdSetService:    service.SingletonFbAdSetService(),
			fbAdService:       service.SingletonFbAdService(),
		}
	})
	return _fbCampaginLogic
}

// UpdateFBAds 增量更新facebook 所有广告数据
func (f *FbCampaginLogic) UpdateFBAds(ctx context.Context,
	fields *bean.CampaignListRequest) (bean.UpdateResult, error) {
	var (
		accountIDs []string
		refresh    int

		campaignCounts int64
		adSetCounts    int64
		adCounts       int64
	)

	if fields != nil {
		accountIDs = fields.AccountID
		refresh = fields.Refresh
	}
	// 获取当前开通自动化投放的账号
	accountList := f.ramblerService.GetFBCruiserAccount(ctx, accountIDs...)
	fmt.Println("accountList:==", accountList)
	for _, it := range accountList {
		// 获取广告系列最新的更新时间 暂定全量更新，
		// 全量更新 ，暂不Redis 设置时间戳，因为 facebook  update_time,预算更改，不会更新update_time字段问题，所以全量拉取
		campaignAt := f.fbCampaginService.GetMaxUpdateTime(ctx, "2543780585639422")
		// 获取广告组最新的更新时间
		adSetAt := f.fbAdSetService.GetMaxUpdateTime(ctx, "2543780585639422")
		// 获取广告最新的更新时间
		adAt := f.fbAdService.GetMaxUpdateTime(ctx, "2543780585639422")

		// 账户的新增广告系列
		// 根据前端传值 refresh == 0:刷新所有; refresh ==1: 只刷新广告组和广告; refresh ==2: 系列
		if refresh == 0 || refresh == 2 {
			campaigns := f.fbCampaginService.GetCampaignsByAccount(ctx, it.AccountId, constants.EmptyString,
				constants.RetryInit, nil, campaignAt)
			glog.Infof(ctx, "account: %s,Insert Campaigns success: %d条，maxTime: %d ", it.AccountId, len(campaigns), campaignAt)
			atomic.AddInt64(&campaignCounts, int64(len(campaigns)))
			f.fbCampaginService.InsertCampaignsByAccount(ctx, campaigns, it.AccountId, constants.RealBoolean)
		}
		if refresh == 0 || refresh == 1 {
			// 账户的新增广告组
			adSets := f.fbAdSetService.GetFBAdSetList(ctx, it.AccountId, constants.EmptyString, constants.RetryInit,
				nil, adSetAt) // > 1669015007 有6条数据
			glog.Infof(ctx, "account: %s,Insert AdSets success: %d条，maxTime: %d ", it.AccountId, len(adSets), adSetAt)
			atomic.AddInt64(&adSetCounts, int64(len(adSets)))
			f.fbAdSetService.InsertAdSetByAccount(ctx, adSets, it.AccountId, constants.RealBoolean)

			// 账户的新增广告
			ads := f.fbAdService.GetFBAdList(ctx, it.AccountId, constants.EmptyString, constants.RetryInit,
				nil, adAt) // > 1669015007 有6条数据
			glog.Infof(ctx, "account: %s,Insert Ads success: %d条，maxTime: %d ", it.AccountId, len(ads), adAt)
			atomic.AddInt64(&adCounts, int64(len(ads)))
			f.fbAdService.InsertAdByAccount(ctx, ads, it.AccountId, constants.RealBoolean)
		}
	}
	result := bean.UpdateResult{
		Campagins: campaignCounts,
		AdSets:    adSetCounts,
		Ads:       adCounts,
	}
	return result, nil
}

// InsertCampaigns 拉取并入库广告系列数据
//func (f *FbcampaginLogic) InsertCampaigns(ctx context.Context) (string, error) {
//	accountList := f.ramblerService.GetFBCruiserAccount(ctx)
//	for _, it := range accountList {
//		timestamp := f.fbcampaginService.GetMaxUpdateTime(ctx, it.AccountId)
//	}
//
//	return "", nil
//}
```

## internal-rpc

### rpc - application_console.go

```go
package rpc

import (
	"context"
	"gitlab.ftsview.com/aircraft/schedule-console/internal/handler/bean"
	"gitlab.ftsview.com/fotoable-go/gerrors"
)
import (
	. "gitlab.ftsview.com/micro/application"
)

func FilterFBAccountID(ctx context.Context, accountIDs ...string) ([]bean.AccountID, error) {
	//filter := map[string]interface{}{"is_used_cruiser": true}
	//if len(accountIDs) > 0 {
	//	filter["account_id"] = accountIDs
	//}
	data, err := ApplicationAdapter.FilterFBAccountID(ctx, &FilterFBAccountIDReq{
		Id:              accountIDs,
		IsEnableCruiser: true,
	})
	if err != nil {
		return nil, gerrors.Wrap(err, "rpc FilterFBAccountID err")
	}
	var Ids []bean.AccountID
	for _, it := range data.Id {
		id := bean.AccountID{AccountId: it}
		Ids = append(Ids, id)
	}
	return Ids, nil
}
```

### rpc - client.go  

```go
package rpc

import (
	"fmt"
	"strings"

	"gitlab.ftsview.com/aircraft/schedule-console/internal/constants"

	"gitlab.ftsview.com/micro/application"
)

func MustInitRpc(endpoint []string) {
	addr := fmt.Sprintf("etcd://%s", strings.Join(endpoint, constants.Comma))
	application.ApplicationAdapterInit(addr)
}
```

## internal-service

### service-rambler.go（）

```go
package service

import (
	"context"
	"gitlab.ftsview.com/aircraft/schedule-console/internal/handler/bean"
	"gitlab.ftsview.com/aircraft/schedule-console/internal/rpc"
	"sync"
)

var (
	_ramblerOnce    sync.Once
	_ramblerService *RamblerService
)

type RamblerService struct {
}

func SingletonRamblerService() *RamblerService {
	_ramblerOnce.Do(func() {
		_ramblerService = &RamblerService{}
	})
	return _ramblerService
}

func (l *RamblerService) GetFBCruiserAccount(ctx context.Context, accountIDs ...string) []bean.AccountID {
	//filter := map[string]interface{}{"is_used_cruiser": true}
	//if len(accountIDs) > 0 {
	//	filter["account_id"] = accountIDs
	//}
	res, err := rpc.FilterFBAccountID(ctx, accountIDs...)
	if err != nil {
		return nil
	}
	return res
}
```

### service-fb_campagin.go

```go
package service

import (
	"context"
	"fmt"
	"github.com/spf13/cast"
	"gitlab.ftsview.com/aircraft/schedule-console/internal/config"
	"gitlab.ftsview.com/aircraft/schedule-console/internal/constants"
	"gitlab.ftsview.com/aircraft/schedule-console/internal/handler/bean"
	"gitlab.ftsview.com/aircraft/schedule-console/internal/store/model"
	"gitlab.ftsview.com/aircraft/schedule-console/internal/store/query"
	"gitlab.ftsview.com/fotoable-go/glog"
	"gitlab.ftsview.com/fotoable-go/gmysql"
	"gitlab.ftsview.com/fotoable-go/gredis"
	"gitlab.ftsview.com/fotoable-go/gutil"
	"gopkg.in/resty.v1"
	"strconv"
	"sync"
	"time"
)

var (
	_fbCampaginOnce    sync.Once
	_fbCampaginService *FbCampaginService
)

type FbCampaginService struct {
}

func SingletonFbCampaginService() *FbCampaginService {
	_fbCampaginOnce.Do(func() {
		_fbCampaginService = &FbCampaginService{}
	})
	return _fbCampaginService
}

type FBCursor struct {
	Before string `json:"before"`
	After  string `json:"after"`
}
type FBPaging struct {
	Cursors FBCursor `json:"cursors"`
	Next    string   `json:"next"`
}

// FBCampaignResult FaceBook获取的原始Campagin
type FBCampaignResult struct {
	Data   []bean.FBCampagin `json:"data"`
	Paging FBPaging          `json:"paging"`
}

// FbCampaignResult 处理后返回的Campagin
type FbCampaignResult struct {
	Data   []bean.FbCampagin `json:"data"`
	Paging FBPaging          `json:"paging"`
}

func (f *FbCampaginService) GetMaxUpdateTime(ctx context.Context, accountId string) int64 {
	maxTime, err := gredis.Redis(constants.RedisName).HGet(ctx, constants.CampaignMaxUpdateTime, accountId)
	if err != nil {
		glog.Error(ctx, "GetMaxUpdateTime error: %s  ", err.Error())
		return 0
	}
	return cast.ToInt64(maxTime)
}

//GetCampaignsByAccount 获取广告系列筛选项
func (f *FbCampaginService) GetCampaignsByAccount(ctx context.Context, accountId, url string, retry int,
	accCps []bean.FbCampagin, timestamp int64) []bean.FbCampagin {

	if url == constants.EmptyString {
		filter := []bean.FBFilterParams{{
			Field:    "campaign.delivery_status",
			Operator: "IN",
			Value:    []string{"active", "deleted", "archived", "inactive", "off", "pending"}}}
		url = fmt.Sprintf(`%sact_%s/campaigns/?fields=%s&access_token=%s`,
			config.GlobConfig.FbApi.BaseUrl, accountId, constants.FBCampaignFields, config.GlobConfig.FbApi.AccessToken)
		if timestamp != constants.EmptyInt {
			filter = append(filter, bean.FBFilterParams{
				Field:    "updated_time",
				Operator: "GREATER_THAN",
				Value:    strconv.FormatInt(timestamp, 10)})
		}
		url = fmt.Sprintf(`%s&filtering=%s`, url, gutil.Object2JSON(filter))
	}
	resp, err := resty.New().R().Get(url)
	// 请求重试机制 重试3次
	if err != nil {
		if retry >= constants.RetryNum {
			glog.Error(ctx, "GetCampaignsByAccount error: ", err)
		} else {
			retry++
			f.GetCampaignsByAccount(ctx, accountId, constants.EmptyString, retry, accCps, timestamp)
		}
	}
	glog.Info(ctx, "GetCampaignsByAccount: success ", accountId)
	var result FbCampaignResult
	var faceBookCampagin FBCampaignResult
	if err := gutil.JSON2ObjectE(resp.Body(), &faceBookCampagin); err != nil {
		glog.Errorf(ctx, "GetCampaignsByAccount: json to object error.accountID: %s, error: %s",
			accountId, err.Error())
	}

	result.Paging = faceBookCampagin.Paging
	for _, datum := range faceBookCampagin.Data {
		var data bean.FbCampagin
		// 更新时间
		if datum.UpdatedTime == constants.EmptyString {
			data.UpdatedAt = 0
		} else {
			updatedAt, err := time.Parse(constants.DateLayout, datum.UpdatedTime)
			if err != nil {
				return nil
			}
			data.UpdatedAt = updatedAt.Unix()
		}
		// 创建时间
		if datum.CreatedTime == constants.EmptyString {
			data.CreatedAt = 0
		} else {
			createAt, err := time.Parse(constants.DateLayout, datum.CreatedTime)
			if err != nil {
				return nil
			}
			data.UpdatedAt = createAt.Unix()
		}
		// data
		faceBookCampaginData, err := gutil.Object2JSONE(&datum)
		if err != nil {
			return nil
		}
		data.Data = faceBookCampaginData

		// 主要字段
		data.CampaignID = datum.CampaignID
		data.AccountID = datum.AccountID
		//data.UpdatedAt = updatedAt.Unix()
		data.Name = datum.Name
		//data.CreatedAt = createdAt.Unix()
		data.BidStrategy = datum.BidStrategy
		data.DailyBudget = datum.DailyBudget
		data.EffectiveStatus = datum.EffectiveStatus
		data.LifetimeBudget = datum.LifetimeBudget
		data.SmartPromotionType = datum.SmartPromotionType
		result.Data = append(result.Data, data)
	}
	accCps = append(accCps, result.Data...)
	if result.Paging.Next != constants.EmptyString {
		accCps = f.GetCampaignsByAccount(ctx, accountId, result.Paging.Next, 0, accCps, timestamp)
	}
	return accCps
}

// InsertCampaignsByAccount 广告 系列入库
func (l *FbCampaginService) InsertCampaignsByAccount(ctx context.Context, accCps []bean.FbCampagin, accountID string, types bool) {

	db := gmysql.DB(ctx, config.GlobConfig.Mysql.DBName)
	campaginTable := query.Use(db).FbCampaign

	//FB返回的结果集是按照降序排列的，获取本次操作的最大更新时间
	if len(accCps) == 0 {
		return
	}
	if types { // 增量更新 || 全部新增
		for _, i2 := range accCps {
			var accCampagin model.FbCampaign
			accCampagin.CampaignID = i2.CampaignID
			accCampagin.AccountID = i2.AccountID
			accCampagin.UpdatedAt = i2.UpdatedAt
			accCampagin.Name = i2.Name
			accCampagin.CreatedAt = i2.CreatedAt
			accCampagin.BidStrategy = i2.BidStrategy
			accCampagin.DailyBudget = i2.DailyBudget
			accCampagin.EffectiveStatus = i2.EffectiveStatus
			accCampagin.LifetimeBudget = i2.LifetimeBudget
			accCampagin.SmartPromotionType = i2.SmartPromotionType
			accCampagin.Data = i2.Data
			err := campaginTable.WithContext(ctx).Save(&accCampagin)
			if err != nil {
				glog.Error(ctx, "增量更新 || 全部新增 campaigns error: %s", err.Error())
				return
			}
		}
	}
	//全量更新 ，暂不Redis 设置时间戳，因为   facebopl  update_time,预算更改，不会更新update_time字段问题，所以全量拉取
	//maxTime, err := time.Parse(constants.DateTimeLayout, accCps[0].UpdatedAt)
	//if err != nil {
	//	gutil.DingTalkAlarm(constants.ServiceName, fmt.Sprintf("MaxTime error, %s", accCps[0].UpdatedAt))
	//	glog.Error(ctx, "MaxTime error: %s", err.Error())
	//	return
	//}
	if err := gredis.Redis(constants.RedisName).HSet(
		ctx, constants.CampaignMaxUpdateTime, accountID, cast.ToString(accCps[0].UpdatedAt)); err != nil {
		glog.Error(ctx, "set max time to redis error: %s", err.Error())
		return
	}
}
```

### service-fb_ad_set.go

```go
package service

import (
	"context"
	"fmt"
	"github.com/spf13/cast"
	"gitlab.ftsview.com/aircraft/schedule-console/internal/config"
	"gitlab.ftsview.com/aircraft/schedule-console/internal/constants"
	"gitlab.ftsview.com/aircraft/schedule-console/internal/handler/bean"
	"gitlab.ftsview.com/aircraft/schedule-console/internal/store/model"
	"gitlab.ftsview.com/aircraft/schedule-console/internal/store/query"
	"gitlab.ftsview.com/fotoable-go/glog"
	"gitlab.ftsview.com/fotoable-go/gmysql"
	"gitlab.ftsview.com/fotoable-go/gredis"
	"gitlab.ftsview.com/fotoable-go/gutil"
	"gopkg.in/resty.v1"
	"strconv"
	"sync"
	"time"
)

var (
	_fbAdSetOnce    sync.Once
	_fbAdSetService *FbAdSetService
)

type FbAdSetService struct {
}

func SingletonFbAdSetService() *FbAdSetService {
	_fbAdSetOnce.Do(func() {
		_fbAdSetService = &FbAdSetService{}
	})
	return _fbAdSetService
}

// FBAdsetResult FaceBook获取的原始广告组Adset
type FBAdsetResult struct {
	Data   []bean.FBAdSet `json:"data"`
	Paging FBPaging       `json:"paging"`
}

// FbAdsetResult 处理后返回的广告组Adset
type FbAdsetResult struct {
	Data   []bean.FbAdset `json:"data"`
	Paging FBPaging       `json:"paging"`
}

//GetMaxUpdateTime 获取广告组最新的数据更新时间
func (f *FbAdSetService) GetMaxUpdateTime(ctx context.Context, accountId string) int64 {
	maxTime, err := gredis.Redis(constants.RedisName).HGet(ctx, constants.AdSetMaxUpdateTime, accountId)
	if err != nil {
		glog.Error(ctx, "GetMaxUpdateTime error: %s  ", err.Error())
		return 0
	}
	return cast.ToInt64(maxTime)
}

//GetFBAdSetList 获取fb后台的广告组列表
func (f *FbAdSetService) GetFBAdSetList(ctx context.Context, accountId, url string, retry int, accCps []bean.FbAdset,
	timestamp int64) []bean.FbAdset {

	if url == constants.EmptyString {
		filter := []bean.FBFilterParams{{
			Field:    "adset.delivery_status",
			Operator: "IN",
			Value:    []string{"active", "deleted", "archived", "inactive", "off", "pending"}}}
		url = fmt.Sprintf(`%sact_%s/adsets/?fields=%s&access_token=%s`, config.GlobConfig.FbApi.BaseUrl,
			accountId, constants.FBAdSetFields, config.GlobConfig.FbApi.AccessToken)
		if timestamp != constants.EmptyInt {
			filter = append(filter, bean.FBFilterParams{
				Field:    "updated_time",
				Operator: "GREATER_THAN",
				Value:    strconv.FormatInt(timestamp, 10)})
		}
		url = fmt.Sprintf(`%s&filtering=%s`, url, gutil.Object2JSON(filter))
	}
	resp, err := resty.New().R().Get(url)
	// 请求重试机制 重试3次
	if err != nil {
		if retry >= constants.RetryNum {
			glog.Error(ctx, "GetAdSetsByAccount error: ", err)
		} else {
			retry++
			f.GetFBAdSetList(ctx, accountId, constants.EmptyString, retry, accCps, timestamp)
		}
	}
	glog.Info(ctx, "GetAdSetsByAccount: success ", accountId)
	var (
		result        FbAdsetResult // 处理后返回的广告组Adset
		facebookAdSet FBAdsetResult // FaceBook获取的原始广告组Adset
	)
	if err := gutil.JSON2ObjectE(resp.Body(), &facebookAdSet); err != nil {
		glog.Errorf(ctx, "GetCampaignsByAccount: json to object error.accountID: %s, error: %s",
			accountId, err.Error())
	}

	result.Paging = facebookAdSet.Paging
	for _, datum := range facebookAdSet.Data {
		adSet := bean.FbAdset{
			AdsetID:         datum.AdsetID,
			CampaignID:      datum.CampaignID,
			AccountID:       datum.AccountID,
			Name:            datum.Name,
			EffectiveStatus: datum.EffectiveStatus,
			BidStrategy:     datum.BidStrategy,
			DailyBudget:     datum.DailyBudget,
			LifetimeBudget:  datum.LifetimeBudget,
			//AttributionSpec:  datum.AttributionSpec,
			OptimizationGoal: datum.OptimizationGoal,
			//CreatedAt: 0,
			//UpdatedAt: 0
			//Data:             "",
		}
		// 更新时间
		if datum.UpdatedTime == constants.EmptyString {
			adSet.UpdatedAt = 0
		} else {
			updatedAt, err := time.Parse(constants.DateLayout, datum.UpdatedTime)
			if err != nil {
				return nil
			}
			adSet.UpdatedAt = updatedAt.Unix()
		}
		// 创建时间
		if datum.CreatedTime == constants.EmptyString {
			adSet.CreatedAt = 0
		} else {
			createAt, err := time.Parse(constants.DateLayout, datum.CreatedTime)
			if err != nil {
				return nil
			}
			adSet.UpdatedAt = createAt.Unix()
		}
		// AttributionSpec
		attributionSpec, err := gutil.Object2JSONE(&datum.AttributionSpec)
		if err != nil {
			return nil
		}
		adSet.AttributionSpec = attributionSpec
		// Data
		adSetData, err := gutil.Object2JSONE(&datum)
		if err != nil {
			return nil
		}
		adSet.Data = adSetData
		result.Data = append(result.Data, adSet)
	}
	accCps = append(accCps, result.Data...)
	if result.Paging.Next != constants.EmptyString {
		accCps = f.GetFBAdSetList(ctx, accountId, result.Paging.Next, constants.RetryInit, accCps, timestamp)
	}
	return accCps
}

// InsertAdSetByAccount 广告组信息入库
func (f *FbAdSetService) InsertAdSetByAccount(ctx context.Context, accCps []bean.FbAdset, accountID string, types bool) {

	db := gmysql.DB(ctx, config.GlobConfig.Mysql.DBName)
	adSetTable := query.Use(db).FbAdSet

	//FB返回的结果集是按照降序排列的，获取本次操作的最大更新时间
	if len(accCps) == 0 {
		return
	}
	if types { // 增量更新 || 全部新增
		for _, cp := range accCps {
			accAdSet := model.FbAdSet{
				AdsetID:          cp.AdsetID,
				CampaignID:       cp.CampaignID,
				AccountID:        cp.AccountID,
				Name:             cp.Name,
				EffectiveStatus:  cp.EffectiveStatus,
				BidStrategy:      cp.BidStrategy,
				DailyBudget:      cp.DailyBudget,
				LifetimeBudget:   cp.LifetimeBudget,
				AttributionSpec:  cp.AttributionSpec,
				OptimizationGoal: cp.OptimizationGoal,
				CreatedAt:        cp.CreatedAt,
				UpdatedAt:        cp.UpdatedAt,
				Data:             cp.Data,
			}
			err := adSetTable.WithContext(ctx).Save(&accAdSet)
			if err != nil {
				glog.Error(ctx, "增量更新 || 全部新增 adSet error: %s", err.Error())
				return
			}
		}
	}
	// 全量更新 ，暂不Redis 设置时间戳，因为   facebopl  update_time,预算更改，不会更新update_time字段问题，所以全量拉取
	//maxTime, err := time.Parse(constants.DateTimeLayout, accCps[0].UpdatedAt)
	//if err != nil {
	//	gutil.DingTalkAlarm(constants.ServiceName, fmt.Sprintf("adset MaxTime error, %s adsetID: %s",
	//		accCps[0].UpdatedAt,
	//		accCps[0].AdsetID))
	//	glog.Error(ctx, "adset MaxTime error: %s", err.Error())
	//	return
	//}
	if err := gredis.Redis(constants.RedisName).HSet(
		ctx, constants.AdSetMaxUpdateTime, accountID, cast.ToString(accCps[0].UpdatedAt)); err != nil {
		glog.Error(ctx, "adset set max time to redis error: %s", err.Error())
		return
	}
}
```

### service-fb_ad.go

```go
package service

import (
	"context"
	"fmt"
	"github.com/spf13/cast"
	"gitlab.ftsview.com/aircraft/schedule-console/internal/config"
	"gitlab.ftsview.com/aircraft/schedule-console/internal/constants"
	"gitlab.ftsview.com/aircraft/schedule-console/internal/handler/bean"
	"gitlab.ftsview.com/aircraft/schedule-console/internal/store/model"
	"gitlab.ftsview.com/aircraft/schedule-console/internal/store/query"
	"gitlab.ftsview.com/fotoable-go/glog"
	"gitlab.ftsview.com/fotoable-go/gmysql"
	"gitlab.ftsview.com/fotoable-go/gredis"
	"gitlab.ftsview.com/fotoable-go/gutil"
	"gopkg.in/resty.v1"
	"strconv"
	"sync"
	"time"
)

var (
	_fbAdOnce    sync.Once
	_fbAdService *FbAdService
)

type FbAdService struct {
}

func SingletonFbAdService() *FbAdService {
	_fbAdOnce.Do(func() {
		_fbAdService = &FbAdService{}
	})
	return _fbAdService
}

// FBAdResult FaceBook获取的原始广告Ad
type FBAdResult struct {
	Data   []bean.FBAd `json:"data"`
	Paging FBPaging    `json:"paging"`
}

// FbAdResult 处理后返回的广告Ad
type FbAdResult struct {
	Data   []bean.FbAd `json:"data"`
	Paging FBPaging    `json:"paging"`
}

//GetMaxUpdateTime 获取广告新的数据更新时间
func (f *FbAdService) GetMaxUpdateTime(ctx context.Context, accountId string) int64 {
	maxTime, err := gredis.Redis(constants.RedisName).HGet(ctx, constants.AdMaxUpdateTime, accountId)
	if err != nil {
		glog.Error(ctx, "GetMaxUpdateTime error: %s  ", err.Error())
		return 0
	}
	return cast.ToInt64(maxTime)
}

//GetFBAdList 获取fb的广告数据
func (f *FbAdService) GetFBAdList(ctx context.Context, accountId, url string, retry int, accCps []bean.FbAd,
	timestamp int64) []bean.FbAd {

	if url == constants.EmptyString {
		filter := []bean.FBFilterParams{{
			Field:    "ad.delivery_status",
			Operator: "IN",
			Value:    []string{"active", "deleted", "archived", "inactive", "off", "pending"}}}
		url = fmt.Sprintf(`%sact_%s/ads/?fields=%s&access_token=%s`, config.GlobConfig.FbApi.BaseUrl,
			accountId, constants.FBAdFields, config.GlobConfig.FbApi.AccessToken)
		// 增量拉取数据
		if timestamp != constants.EmptyInt {
			filter = append(filter, bean.FBFilterParams{
				Field:    "updated_time",
				Operator: "GREATER_THAN",
				Value:    strconv.FormatInt(timestamp, 10)})
		}
		url = fmt.Sprintf(`%s&filtering=%s`, url, gutil.Object2JSON(filter))
	}
	resp, err := resty.New().R().Get(url)
	glog.Info(ctx, url)
	// 请求重试机制 重试3次
	if err != nil {
		if retry >= constants.RetryNum {
			glog.Error(ctx, "GetAdsByAccount error: ", err)
		} else {
			retry++
			f.GetFBAdList(ctx, accountId, constants.EmptyString, retry, accCps, timestamp)
		}
	}
	glog.Info(ctx, "GetAdsByAccount: success ", accountId)
	var (
		result     FbAdResult // 处理后返回的广告Ad
		facebookAd FBAdResult // FaceBook获取的原始广告Ad
	)

	if err = gutil.JSON2ObjectE(resp.Body(), &facebookAd); err != nil {
		glog.Errorf(ctx, "GetAdsByAccount: json to object error.accountID: %s, error: %s",
			accountId, err.Error())
	}
	result.Paging = facebookAd.Paging
	for _, datum := range facebookAd.Data {
		ad := bean.FbAd{
			AdID:            datum.AdID,
			AdsetID:         datum.AdsetID,
			CampaignID:      datum.CampaignID,
			CreativeID:      datum.Creative["id"],
			Name:            datum.Name,
			EffectiveStatus: datum.EffectiveStatus,
			//CreatedAt:       0,
			//UpdatedAt:       0,
			//Data:            "",
		}
		// 更新时间
		if datum.UpdatedTime == constants.EmptyString {
			ad.UpdatedAt = 0
		} else {
			updatedAt, err := time.Parse(constants.DateLayout, datum.UpdatedTime)
			if err != nil {
				return nil
			}
			ad.UpdatedAt = updatedAt.Unix()
		}
		// 创建时间
		if datum.CreatedTime == constants.EmptyString {
			ad.CreatedAt = 0
		} else {
			createAt, err := time.Parse(constants.DateLayout, datum.CreatedTime)
			if err != nil {
				return nil
			}
			ad.UpdatedAt = createAt.Unix()
		}
		// Data
		adData, err := gutil.Object2JSONE(&datum)
		if err != nil {
			return nil
		}
		ad.Data = adData
		result.Data = append(result.Data, ad)
	}
	accCps = append(accCps, result.Data...)
	if result.Paging.Next != constants.EmptyString {
		accCps = f.GetFBAdList(ctx, accountId, result.Paging.Next, constants.RetryInit, accCps, timestamp)
	}
	return accCps
}

//InsertAdByAccount 入库广告数据
func (f *FbAdService) InsertAdByAccount(ctx context.Context, accCps []bean.FbAd, accountID string, types bool) {

	db := gmysql.DB(ctx, config.GlobConfig.Mysql.DBName)
	adTable := query.Use(db).FbAd

	//FB返回的结果集是按照降序排列的，获取本次操作的最大更新时间
	if len(accCps) == 0 {
		return
	}
	if types { // 增量更新
		for _, cp := range accCps {
			accAd := model.FbAd{
				AdID:            cp.AdID,
				AdsetID:         cp.AdsetID,
				CampaignID:      cp.CampaignID,
				CreativeID:      cp.CreativeID,
				Name:            cp.Name,
				EffectiveStatus: cp.EffectiveStatus,
				CreatedAt:       cp.CreatedAt,
				UpdatedAt:       cp.UpdatedAt,
				Data:            cp.Data,
			}
			err := adTable.WithContext(ctx).Save(&accAd)
			if err != nil {
				glog.Error(ctx, "增量更新 || 全部新增 ad error: %s", err.Error())
				return
			}
		}
	}
	//默认情况FB返回的为降序
	//maxTime, err := time.Parse(constants.DateTimeLayout, accCps[0].UpdatedAt)
	//if err != nil {
	//	gutil.DingTalkAlarm(constants.ServiceName, fmt.Sprintf("MaxTime error, %s", accCps[0].UpdatedAt))
	//	glog.Error(ctx, "MaxTime error: %s", err.Error())
	//	return
	//}

	//插入BI库
	//f.SyncMaterial(ctx, accCps, accountID)

	glog.Infof(ctx, "accountID max update_time. %s", cast.ToString(accCps[0].UpdatedAt))
	if err := gredis.Redis(constants.RedisName).HSet(
		ctx, constants.AdMaxUpdateTime, accountID, cast.ToString(accCps[0].UpdatedAt)); err != nil {
		glog.Error(ctx, "ad set max time to redis error: %s", err.Error())
		return
	}
}
```

## internal - handler

### bean - common.go

```go
package bean

type BaseRequest struct {
	CompanyID int         `json:"company_id"`
	AppID     interface{} `json:"app_id"`
	GameID    interface{} `json:"game_id"`
}

type BasePageRequest struct {
	Page int64 `json:"page"`
	Size int64 `json:"size"`
}

type AccountID struct {
	AccountId string `json:"id"`
}
```

### bean - fb_ad.go

```go
package bean

// FbAd mapped from table <fb_ad>
type FbAd struct {
	AdID            string `gorm:"column:ad_id;primaryKey" json:"ad_id"`                     // 广告ID
	AdsetID         string `gorm:"column:adset_id;not null" json:"adset_id"`                 // 广告组ID
	CampaignID      string `gorm:"column:campaign_id;not null" json:"campaign_id"`           // 广告系列ID
	CreativeID      string `gorm:"column:creative_id;not null" json:"creative_id"`           // 广告创意ID
	Name            string `gorm:"column:name;not null" json:"name"`                         // 广告名称
	EffectiveStatus string `gorm:"column:effective_status;not null" json:"effective_status"` // 广告状态
	CreatedAt       int64  `gorm:"column:created_at;autoUpdateTime" json:"created_at"`       // 创建时间戳
	UpdatedAt       int64  `gorm:"column:updated_at;autoCreateTime" json:"updated_at"`       // 更新时间戳
	Data            string `gorm:"column:data" json:"data"`
}

// FBAd FaceBook原始广告数据
type (
	FBAd struct {
		// ID                       primitive.ObjectID `bson:"_id" json:"-"`
		CompanyID            int                      `bson:"company_id" json:"company_id,omitempty"`
		AdID                 string                   `bson:"ad_id" json:"id,omitempty"`
		AccountID            string                   `bson:"account_id" json:"account_id,omitempty"`
		AdReviewFeedback     map[string]interface{}   `bson:"ad_review_feedback" json:"ad_review_feedback,omitempty"` // 审阅后的审阅反馈
		Adlabels             []Adlabels               `bson:"adlabels" json:"adlabels,omitempty"`                     // 与此活动关联的广告标签
		Adset                map[string]interface{}   `bson:"adset" json:"adset,omitempty"`                           // 包含此广告的广告组
		AdsetID              string                   `bson:"adset_id,omitempty" json:"adset_id,omitempty"`           // 包含此广告的光广告组ID
		BidAmount            int32                    `bson:"bid_amount,omitempty" json:"bid_amount,omitempty"`       // 广告的出价金额
		Campaign             map[string]interface{}   `bson:"campaign" json:"campaign,omitempty"`                     //广告的广告系列
		CampaignID           string                   `bson:"campaign_id" json:"campaign_id,omitempty"`
		ConfiguredStatus     string                   `bson:"configured_status" json:"configured_status,omitempty"`           //  广告的配置状态
		ConversionDomain     string                   `bson:"conversion_domain" json:"conversion_domain,omitempty"`           //  发生转换的域
		Creative             map[string]string        `bson:"creative" json:"creative,omitempty"`                             //  广告创意
		EffectiveStatus      string                   `bson:"effective_status" json:"effective_status,omitempty"`             // ad的有效状态
		IssuesInfo           []map[string]interface{} `bson:"issues_info,omitempty" json:"issues_info,omitempty"`             // 广告的问题使其无法交付
		LastUpdatedByAppId   string                   `bson:"last_updated_by_app_id" json:"last_updated_by_app_id,omitempty"` // 指示用于广告最新更新的应用程序。
		Name                 string                   `bson:"name" json:"name,omitempty"`                                     //  广告名称
		PreviewShareableLink string                   `bson:"preview_shareable_link" json:"preview_shareable_link,omitempty"` // 允许用户在不同位置预览广告的链接
		Recommendations      []map[string]interface{} `bson:"recommendations" json:"recommendations,omitempty"`               //  广告的建议
		SourceAd             map[string]interface{}   `bson:"source_ad" json:"source_ad,omitempty"`                           //  复制此广告的源广告
		SourceAdId           string                   `bson:"source_ad_id" json:"source_ad_id,omitempty"`                     //  复制此广告的源广告id
		Status               string                   `bson:"status,omitempty" json:"status,omitempty"`                       // 状态
		TrackingSpecs        []map[string]interface{} `bson:"tracking_specs" json:"tracking_specs,omitempty"`                 //  跟踪规范
		CreatedTime          string                   `bson:"created_time" json:"created_time,omitempty"`
		UpdatedTime          string                   `bson:"updated_time" json:"updated_time,omitempty"`
	}
)
```

### bean - fb_ad_set.go

```go
package bean

// AdSetListRequest 广告组请求参数
type AdSetListRequest struct {
	BaseRequest
	BasePageRequest
	AccountID  []string `json:"account_ids"`
	CampaignID []string `json:"campaign_ids"`
	Status     []string `json:"status"`
	AdSetID    string   `json:"adset_id"`
	KeyWord    string   `json:"keyword"`
}

// FbAdset mapped from table <fb_adset>
type FbAdset struct {
	AdsetID          string `gorm:"column:adset_id;primaryKey" json:"adset_id"`                 // 广告组ID
	CampaignID       string `gorm:"column:campaign_id;not null" json:"campaign_id"`             // 广告系列ID
	AccountID        string `gorm:"column:account_id;not null" json:"account_id"`               // 账号ID
	Name             string `gorm:"column:name;not null" json:"name"`                           // 广告组名称
	EffectiveStatus  string `gorm:"column:effective_status;not null" json:"effective_status"`   // 广告组状态
	BidStrategy      string `gorm:"column:bid_strategy;not null" json:"bid_strategy"`           // 广告组策略
	DailyBudget      string `gorm:"column:daily_budget;not null" json:"daily_budget"`           // 日预算
	LifetimeBudget   string `gorm:"column:lifetime_budget;not null" json:"lifetime_budget"`     // 总预算
	AttributionSpec  string `gorm:"column:attribution_spec" json:"attribution_spec"`            // 归因设置
	OptimizationGoal string `gorm:"column:optimization_goal;not null" json:"optimization_goal"` // 优化目标
	CreatedAt        int64  `gorm:"column:created_at;autoUpdateTime" json:"created_at"`         // 创建时间戳
	UpdatedAt        int64  `gorm:"column:updated_at;autoCreateTime" json:"updated_at"`         // 更新时间戳
	Data             string `gorm:"column:data" json:"data"`
}

// FBAdSet FaceBook原始广告组数据
type (
	FBAdSet struct {
		// ID                       primitive.ObjectID `bson:"_id" json:"-"`
		CompanyID                    int                      `bson:"company_id" json:"company_id,omitempty"`
		CampaignID                   string                   `bson:"campaign_id" json:"campaign_id,omitempty"`
		AdsetID                      string                   `bson:"adset_id" json:"id,omitempty"`
		AccountID                    string                   `bson:"account_id" json:"account_id,omitempty"`
		AdsetSchedule                []map[string]interface{} `bson:"adset_schedule" json:"adset_schedule,omitempty"`             // 广告集计划，表示一天的交货时间表
		AssetFeedId                  string                   `bson:"asset_feed_id" json:"asset_feed_id,omitempty"`               // 包含内容以创建广告的资产源的ID
		Adlabels                     []Adlabels               `bson:"adlabels" json:"adlabels,omitempty"`                         // 与此活动关联的广告标签
		AttributionSpec              []SpecEvents             `bson:"attribution_spec" json:"attribution_spec,omitempty"`         // 转化时间窗
		BidAdjustments               map[string]interface{}   `bson:"bid_adjustments,omitempty" json:"bid_adjustments,omitempty"` // 投标调整类型到值的映射
		BidAmount                    int32                    `bson:"bid_amount,omitempty" json:"bid_amount,omitempty"`           // 广告组投放上线
		BidConstraints               map[string]interface{}   `bson:"bid_constraints,omitempty" json:"bid_constraints,omitempty"` //广告组投标限制条件
		BidInfo                      map[string]interface{}   `bson:"bid_info,omitempty" json:"bid_info,omitempty"`               //  投标目标与投标价值的关系图
		BidStrategy                  string                   `bson:"bid_strategy" json:"bid_strategy,omitempty"`                 //  竞价策略
		BillingEvent                 string                   `bson:"billing_event" json:"billing_event,omitempty"`               //  计费事件
		BudgetRemaining              string                   `bson:"budget_remaining" json:"budget_remaining,omitempty"`         // 广告组剩余预算
		Campaign                     map[string]interface{}   `bson:"campaign" json:"campaign,omitempty"`                         // 广告系列
		ConfiguredStatus             string                   `bson:"configured_status" json:"configured_status,omitempty"`       // 广告组级别的状态
		ContextualBundlingSpec       map[string]interface{}   `bson:"contextual_bundling_spec,omitempty" json:"contextual_bundling_spec,omitempty"`
		CreativeSequence             []interface{}            `bson:"creative_sequence" json:"creative_sequence,omitempty"`           // 向用户显示adgroup序列的顺序
		DailyBudget                  string                   `bson:"daily_budget" json:"daily_budget,omitempty"`                     //  竞选活动的每日预算
		DailyMinSpendTarget          string                   `bson:"daily_min_spend_target" json:"daily_min_spend_target,omitempty"` //  每日最低支出目标
		DailySpendCap                string                   `bson:"daily_spend_cap" json:"daily_spend_cap,omitempty"`               //  每日支出上限
		DestinationType              string                   `bson:"destination_type" json:"destination_type,omitempty"`             //  广告组的投放目标类型（平台）
		EffectiveStatus              string                   `bson:"effective_status" json:"effective_status,omitempty"`
		EndTime                      string                   `bson:"end_time" json:"end_time,omitempty"`                                             // 结束时间
		FrequencyControlSpecs        []map[string]interface{} `bson:"frequency_control_specs" json:"frequency_control_specs,omitempty"`               //  频率控制规格数组
		InstagramActorId             string                   `bson:"instagram_actor_id" json:"instagram_actor_id,omitempty"`                         // Instagram帐户id
		IsDynamicCreative            bool                     `bson:"is_dynamic_creative" json:"is_dynamic_creative,omitempty"`                       // 是否为动态创意广告组
		IssuesInfo                   []map[string]interface{} `bson:"issues_info,omitempty" json:"issues_info,omitempty"`                             // 错误支付信息
		LearningStageInfo            map[string]interface{}   `bson:"learning_stage_info" json:"learning_stage_info,omitempty"`                       // 有关排名或发布系统是否仍在学习此广告集的信息
		LifetimeBudget               string                   `bson:"lifetime_budget" json:"lifetime_budget,omitempty"`                               //  终身预算
		LifetimeImps                 int32                    `bson:"lifetime_imps" json:"lifetime_imps,omitempty"`                                   //  终生印象 仅适用于具有buying_type=FIXED_CPM
		LifetimeSpendCap             string                   `bson:"lifetime_spend_cap" json:"lifetime_spend_cap,omitempty"`                         // 终身开支上限
		LifetimeMinSpendTarget       string                   `bson:"lifetime_min_spend_target" json:"lifetime_min_spend_target,omitempty"`           // 终身最低花费目标
		MultiOptimizationGoalWeight  string                   `bson:"multi_optimization_goal_weight" json:"multi_optimization_goal_weight,omitempty"` // 多目标优化
		Name                         string                   `bson:"name" json:"name,omitempty"`                                                     //  广告组名称
		OptimizationGoal             string                   `bson:"optimization_goal" json:"optimization_goal,omitempty"`                           // 优化方式
		OptimizationSubEvent         string                   `bson:"optimization_sub_event" json:"optimization_sub_event,omitempty"`                 // 优化子事件
		PacingType                   []string                 `bson:"pacing_type,omitempty" json:"pacing_type,omitempty"`                             // 定义活动的步调类
		PromotedObject               map[string]interface{}   `bson:"promoted_object,omitempty" json:"promoted_object,omitempty"`                     // 宣传的对象
		Recommendations              []map[string]interface{} `bson:"recommendations,omitempty" json:"recommendations,omitempty"`                     // 活动的建议
		RecurringBudgetSemantics     bool                     `bson:"recurring_budget_semantics" json:"recurring_budget_semantics,omitempty"`         // 如果这个字段是true，您的日常开支可能会超过您的日常预算，
		ReviewFeedback               string                   `bson:"review_feedback" json:"review_feedback,omitempty"`                               //  动态创意广告评论
		RfPredictionId               string                   `bson:"rf_prediction_id,omitempty" json:"rf_prediction_id,omitempty"`                   // 范围和频率预测ID
		SourceAdset                  map[string]interface{}   `bson:"source_adset" json:"source_adset,omitempty"`                                     //  从中复制此活动的源市场活动id
		SourceAdsetId                string                   `bson:"source_adset_id,omitempty" json:"source_adset_id,omitempty"`                     // 特别广告类别
		StartTime                    string                   `bson:"start_time" json:"start_time,omitempty"`
		Status                       string                   `bson:"status,omitempty" json:"status,omitempty"`                                                     // 状态
		Targeting                    map[string]interface{}   `bson:"targeting" json:"targeting,omitempty"`                                                         // 受众定向
		TargetingOptimizationTypes   []map[string]interface{} `bson:"targeting_optimization_types,omitempty" json:"targeting_optimization_types,omitempty"`         // 将放松的选项作为优化的信号
		TimeBasedAdRotationIdBlocks  []map[string]interface{} `bson:"time_based_ad_rotation_id_blocks,omitempty" json:"time_based_ad_rotation_id_blocks,omitempty"` // 广告组ID的列表
		TimeBasedAdRotationIntervals []map[string]interface{} `bson:"time_based_ad_rotation_intervals,omitempty" json:"time_based_ad_rotation_intervals,omitempty"` // 特定广告创意在活动期间显示的日期范围
		UseNewAppClick               bool                     `bson:"use_new_app_click,omitempty" json:"use_new_app_click,omitempty"`                               // 顶线ID
		CreatedTime                  string                   `bson:"created_time" json:"created_time,omitempty"`
		UpdatedTime                  string                   `bson:"updated_time" json:"updated_time,omitempty"`
	}
	SpecEvents struct {
		EventType  string `json:"event_type" bson:"event_type"`
		WindowDays int    `json:"window_days" bson:"window_days"`
	}
)
```

### bean - fb_campagin.go

```go
package bean

import gginutil "gitlab.ftsview.com/fotoable-go/ggin-util"

// CampaignListRequest 广告系列列表请求参数
type CampaignListRequest struct {
	BaseRequest
	BasePageRequest
	Header             gginutil.HeaderToB `json:"-"`
	AccountID          []string           `json:"account_ids"`
	CampaignID         string             `json:"campaign_id"`
	Objective          string             `json:"objective"`
	SmartPromotionType string             `json:"smart_promotion_type"`
	Status             []string           `json:"status"`
	KeyWord            string             `json:"keyword"`
	Refresh            int                `json:"refresh"`
}

type UpdateResult struct {
	Campagins int64 `json:"campagins"`
	AdSets    int64 `json:"adSets"`
	Ads       int64 `json:"ads"`
}

type FbCampagin struct {
	CampaignID         string `gorm:"column:campaign_id;primaryKey" json:"campaign_id"`                 // 系列ID
	AccountID          string `gorm:"column:account_id;not null" json:"account_id"`                     // 账号ID
	Name               string `gorm:"column:name;not null" json:"name"`                                 // 系列名称
	EffectiveStatus    string `gorm:"column:effective_status;not null" json:"effective_status"`         // 状态
	SmartPromotionType string `gorm:"column:smart_promotion_type;not null" json:"smart_promotion_type"` // 广告系列类型
	BidStrategy        string `gorm:"column:bid_strategy;not null" json:"bid_strategy"`                 // 广告系列策略
	DailyBudget        string `gorm:"column:daily_budget;not null" json:"daily_budget"`                 // 日预算
	LifetimeBudget     string `gorm:"column:lifetime_budget;not null" json:"lifetime_budget"`           // 总预算
	CreatedAt          int64  `gorm:"column:created_at;autoUpdateTime" json:"created_at"`               // 创建时间戳
	UpdatedAt          int64  `gorm:"column:updated_at;autoCreateTime" json:"updated_at"`               // 更新时间戳
	Data               string `gorm:"column:data" json:"data"`
}

type FBFilterParams struct {
	Field    string      `json:"field"`
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
}

// FaceBook原始Campagin数据
type (
	FBCampagin struct {
		// ID                       primitive.ObjectID `bson:"_id" json:"-"`
		CompanyID                int                      `bson:"company_id" json:"company_id,omitempty"`
		CampaignID               string                   `bson:"campaign_id" json:"id,omitempty"`
		AccountID                string                   `bson:"account_id" json:"account_id,omitempty"`
		AdStrategyGroupId        string                   `bson:"ad_strategy_group_id" json:"ad_strategy_group_id,omitempty"`     // 广告策略组ID
		AdStrategyId             string                   `bson:"ad_strategy_id" json:"ad_strategy_id,omitempty"`                 // 广告策略ID
		Adlabels                 []Adlabels               `bson:"adlabels" json:"adlabels,omitempty"`                             // 与此活动关联的广告标签
		BidStrategy              string                   `bson:"bid_strategy" json:"bid_strategy,omitempty"`                     // 竞价策略
		BoostedObjectId          string                   `bson:"boosted_object_id,omitempty" json:"boosted_object_id,omitempty"` // 此活动关联的增强对象
		BrandLiftStudies         []map[string]interface{} `bson:"brand_lift_studies,omitempty" json:"brand_lift_studies,omitempty"`
		BudgetRebalanceFlag      bool                     `bson:"budget_rebalance_flag" json:"budget_rebalance_flag"`
		BudgetRemaining          string                   `bson:"budget_remaining" json:"budget_remaining,omitempty"`                       //剩余预算
		BuyingType               string                   `bson:"buying_type" json:"buying_type,omitempty"`                                 //  购买类型
		CanCreateBrandLiftStudy  bool                     `bson:"can_create_brand_lift_study" json:"can_create_brand_lift_study,omitempty"` //  是否提升研究
		CanUseSpendCap           bool                     `bson:"can_use_spend_cap" json:"can_use_spend_cap,omitempty"`                     //  竞选活动能否设定开支上限
		ConfiguredStatus         string                   `bson:"configured_status" json:"configured_status"`
		CreatedTime              string                   `bson:"created_time" json:"created_time,omitempty"`
		DailyBudget              string                   `bson:"daily_budget" json:"daily_budget,omitempty"` //  竞选活动的每日预算
		EffectiveStatus          string                   `bson:"effective_status" json:"effective_status,omitempty"`
		IsSkadnetworkAttribution bool                     `bson:"is_skadnetwork_attribution" json:"is_skadnetwork_attribution,omitempty"` //  设置为时true表示活动将包括SKAdNetwork，iOS 14。
		IssuesInfo               []map[string]interface{} `bson:"issues_info,omitempty" json:"issues_info,omitempty"`                     // 错误支付信息
		LastBudgetTogglingTime   string                   `bson:"last_budget_toggling_time" json:"last_budget_toggling_time,omitempty"`   // 上次预算切换时间
		LifetimeBudget           string                   `bson:"lifetime_budget" json:"lifetime_budget,omitempty"`                       //  竞选的终身预算
		Name                     string                   `bson:"name" json:"name,omitempty"`                                             //  广告系列名称
		Objective                string                   `bson:"objective" json:"objective,omitempty"`                                   //  活动目标
		PacingType               []string                 `bson:"pacing_type,omitempty" json:"pacing_type,omitempty"`                     // 定义活动的步调类
		PromotedObject           map[string]interface{}   `bson:"promoted_object,omitempty" json:"promoted_object,omitempty"`             // 宣传的对象
		Recommendations          []map[string]interface{} `bson:"recommendations,omitempty" json:"recommendations,omitempty"`             // 活动的建议
		SmartPromotionType       string                   `bson:"smart_promotion_type,omitempty" json:"smart_promotion_type,omitempty"`   // 智能促销类型
		SourceCampaign           map[string]interface{}   `bson:"source_campaign,omitempty" json:"source_campaign,omitempty"`
		SourceCampaignId         string                   `bson:"source_campaign_id" json:"source_campaign_id,omitempty"`                             //  从中复制此活动的源市场活动id
		SpecialAdCategories      []string                 `bson:"special_ad_categories,omitempty" json:"special_ad_categories,omitempty"`             // 特别广告类别
		SpecialAdCategory        string                   `bson:"special_ad_category,omitempty" json:"special_ad_category,omitempty"`                 // 特别广告类别
		SpecialAdCategoryCountry []string                 `bson:"special_ad_category_country,omitempty" json:"special_ad_category_country,omitempty"` // 特殊广告类别的国家/地区字段。
		SpendCap                 string                   `bson:"spend_cap,omitempty" json:"spend_cap,omitempty"`                                     // 开支上限
		StartTime                string                   `bson:"start_time" json:"start_time,omitempty"`
		Status                   string                   `bson:"status,omitempty" json:"status,omitempty"` // 状态
		StopTime                 string                   `bson:"stop_time" json:"stop_time,omitempty"`
		ToplineId                string                   `bson:"topline_id,omitempty" json:"topline_id,omitempty"` // 顶线ID
		UpdatedTime              string                   `bson:"updated_time" json:"updated_time,omitempty"`
	}
	Adlabels struct {
		ID          string `bson:"id" json:"id"`
		Name        string `bson:"name" json:"name"`
		CreatedTime string `bson:"created_time" json:"created_time"`
		UpdatedTime string `bson:"updated_time" json:"updated_time"`
	}
)
```

## redis

![img](https://cdn.nlark.com/yuque/0/2023/png/22219483/1672735989310-84f8a1ae-ae93-4876-bd6d-fc45fbf39c69.png)