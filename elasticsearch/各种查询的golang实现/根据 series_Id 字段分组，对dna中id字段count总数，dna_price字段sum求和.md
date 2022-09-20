# 根据 series_Id 字段分组，对dna中id字段count总数，dna_price字段sum求和



### 业务需求：

**索引名**：dna、series

**关系**：一个series 有 n 个dna，series 与 dna关系为 1 ：n

**要求**：实现按照dna发行数量（系列下**dna**的总数）或者发行金额（系列下**dna**的**dna_price总数**）排序，同时支持分页功能

**思路：**根据 系列（series_Id 字段）分组，对dna中**id字段**count总数，**dna_price字段**sum求和

### dna索引结构如下：

```go
type Dna struct {
    App struct {
        AcCode     string `json:"ac_code"`
        BifUserBid string `json:"bif_user_bid"`
        Id         string `json:"id"`
        Name       string `json:"name"`
        Type       int    `json:"type"`
        ChainName  string `json:"chainName"`
    } `json:"app"`
    BifUser struct {
        Dna721ContractAddress string `json:"dna721_contract_address"`
    } `json:"bif_user"`
    Category   string    `json:"category"`
    CreateTime time.Time `json:"create_time"`
    DisplayUrl string    `json:"display_url"`
    DnaPrice   string    `json:"dna_price"`
    Id         string    `json:"id"`
    Name       string    `json:"name"`
    Number     string    `json:"number"`
    OwnerBid   string    `json:"owner_bid"`
    Series     struct {
        Issuer string `json:"issuer"`
        Name   string `json:"name"`
    } `json:"series"`
    SeriesId string `json:"series_id"`
    Status   int    `json:"status"`
    TokenBid string `json:"token_bid"`
    Url      string `json:"url"`
}
```

### es查询语句

```json
{
  "aggregations": {
    "dnaCountsGroup": {
      "aggregations": {
        "bucket_field": {
          "bucket_sort": {
            "from": 3,
            "size": 3
          }
        },
        "dna_counts_num": {
          "value_count": {
            "field": "id"
          }
        },
        "dna_price_num": {
          "sum": {
            "field": "dna_price"
          }
        }
      },
      "terms": {
        "field": "series_id",
        "order": [
          {
            "dna_counts_num": "desc"
          }
        ],
        "size": 8
      }
    }
  },
  "query": {
    "term": {
      "app.id": "09c464e4-9e11-439a-a6ce-6429022f9284"
    }
  }
  
}
```



### es查询结果

```json
"aggregations": {
    "dnaCountsGroup": {
        "doc_count_error_upper_bound": 0,
        "sum_other_doc_count": 0,
        "buckets": [
            {
                "key": "e499e06b-f135-49c8-835e-e7997e753bd0",
                "doc_count": 100,
                "dna_price_num": {
                    "value": 124.00000095367432
                },
                "dna_counts_num": {
                    "value": 100
                }
            },
            {
                "key": "db0d20bf-69d1-4d72-baa4-4c70a96eadd4",
                "doc_count": 27,
                "dna_price_num": {
                    "value": 170039.40000009537
                },
                "dna_counts_num": {
                    "value": 27
                }
            },
            {
                "key": "09a1febc-1cd9-4199-8eb9-cf55ccf64c98",
                "doc_count": 1,
                "dna_price_num": {
                    "value": 10000
                },
                "dna_counts_num": {
                    "value": 1
                }
            }
        ]
    }
}
```

### go实现

#### CountAggregations

```go
type CountAggregations struct {
	Aggregations struct {
		DnaCountsGroup struct {
			DocCountErrorUpperBound int `json:"doc_count_error_upper_bound"`
			SumOtherDocCount        int `json:"sum_other_doc_count"`
			Buckets                 []struct {
				Key         string `json:"key"`
				DocCount    int64  `json:"doc_count"`
				DnaPriceNum struct {
					Value float64 `json:"value"`
				} `json:"dna_price_num"`
				DnaCountsNum struct {
					Value int64 `json:"value"`
				} `json:"dna_counts_num"`
			} `json:"buckets"`
		} `json:"dnaCountsGroup"`
	} `json:"aggregations"`
}
```

#### SeriesReq

```go
type SeriesReq struct {
	AppId    string `form:"appId"`
	PageNum  int    `form:"pageNum"`                             // 页码
	PageSize int    `form:"pageSize"`                            // 条数
	Field    string `form:"field,optional" validate:"omitempty"` // dnaCounts：根据发行数量排序 ； dnaPrice：根据发行金额排序
	Sort     string `form:"sort,optional" validate:"omitempty"`  // descend降序，ascend升序
}
```

#### logic:

```go
var counts CountAggregations
var result *elastic.SearchResult
var req    SeriesReq

// group
dnaCountsGroup := elastic.NewTermsAggregation().Field("series_id")
switch req.Sort {
case "descend":
	dnaCountsGroup.OrderByAggregation("dna_counts_num", false)
case "ascend":
	dnaCountsGroup.OrderByAggregation("dna_counts_num", true)
default:
	dnaCountsGroup.OrderByAggregation("dna_counts_num", false)
}
dnaCountsGroup.Size(int(total))

// search
searchSource := elastic.NewSearchSource()

// count id 字段
dnaCountsNum := elastic.NewValueCountAggregation().Field("id")
dnaCountsGroup.SubAggregation("dna_counts_num", dnaCountsNum)

// sum dna_price 字段
dnaPriceSum := elastic.NewSumAggregation().Field("dna_price")
dnaCountsGroup.SubAggregation("dna_price_num", dnaPriceSum)

bucketSort := elastic.NewBucketSortAggregation().From((req.PageNum - 1) * req.PageSize).Size(req.PageSize)
dnaCountsGroup.SubAggregation("bucket_field", bucketSort)

searchSource.Aggregation("dnaCountsGroup", dnaCountsGroup)

query := elastic.NewTermQuery("app.id", req.AppId)

result, err = l.svcCtx.OEsClient.OClient.Search().Index("dna").SearchSource(searchSource).Query(query).Do(l.ctx)
if err != nil {
	return nil, err
}

b, err := json.Marshal(result)
if err != nil {
	return nil, err
}
if err := json.Unmarshal(b, &counts); err != nil {
	return nil, err
}
fmt.Println("counts=", counts.Aggregations.DnaCountsGroup.Buckets)

for i, p := range counts.Aggregations.DnaCountsGroup.Buckets {

	seriesResList = append(seriesResList, &types.SeriesRes{
		SeriesId:  p.Key,
		DnaPrice:  p.DnaPriceNum.Value,
		DnaCounts: p.DnaCountsNum.Value,
	})

	dox, err := l.svcCtx.OEsClient.OClient.Get().Index("series").Id(p.Key).Do(l.ctx)
	if err != nil {
		return nil, err
	}
	var serie models.Series
	err = l.svcCtx.OEsClient.GetUnmarshal(&dox.Source, &serie)
	if err != nil {
		return nil, err
	}
	seriesResList[i].SeriesDescription = serie.Description
	seriesResList[i].SeriesName = serie.Name
	seriesResList[i].Issuer = serie.Issuer
	seriesResList[i].ExternalUrl = serie.ExternalUrl
}
```

### 接口请求结果：

```go
{
    "code": 0,
    "msg": "ok",
    "data": {
        "items": [
            {
                "seriesName": "萌萌图可",
                "seriesId": "e499e06b-f135-49c8-835e-e7997e753bd0",
                "externalUrl": "http://www.baidu.com/",
                "seriesDescription": "",
                "issuer": "发行方公司",
                "dnaCounts": 100,
                "dnaPrice": 124.00000095367432
            },
            {
                "seriesName": "测试集合hh2",
                "seriesId": "db0d20bf-69d1-4d72-baa4-4c70a96eadd4",
                "externalUrl": "http://www.baidu.com/",
                "seriesDescription": "萌萌",
                "issuer": "发行方公司",
                "dnaCounts": 27,
                "dnaPrice": 170039.40000009537
            },
            {
                "seriesName": "测试集合5",
                "seriesId": "09a1febc-1cd9-4199-8eb9-cf55ccf64c98",
                "externalUrl": "http://www.baidu.com/",
                "seriesDescription": "萌萌",
                "issuer": "发行方公司",
                "dnaCounts": 1,
                "dnaPrice": 10000
            }
        ],
        "total": 8,
        "pageNum": 2,
        "pageSize": 3
    }
}
```