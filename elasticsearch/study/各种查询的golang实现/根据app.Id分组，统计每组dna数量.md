# 根据app.Id分组，统计每组dna数量

### 需求：

索引：dna

需求：根据app.Id 分组，统计每个app下的dna数量

### dna索引结构

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
    "dnaCounts": {
      "aggregations": {
        "countAgg": {
          "value_count": {
            "field": "id"
          }
        }
      },
      "terms": {
        "field": "app.id",
        "size": 5
      }
    }
  },
  "query": {
    "terms": {
      "app.id": [
        "e05aca13-34f6-4f3e-88d3-76698cad76c9",
        "691f84ed-8ecc-40ab-8880-4674d93c85e4",
        "c0ffecef-dad8-481c-ae05-1a55e8724327",
        "e0ffbbba-b49e-4570-a4b1-1cb74e4ad2ba",
        "f48717b2-a0a3-41c2-a978-f54299a1e42d"
      ]
    }
  },
  "size": 0
}
```



### es查询结果

```json
{
  "took": 0,
  "timed_out": false,
  "_shards": {
    "total": 2,
    "successful": 2,
    "skipped": 0,
    "failed": 0
  },
  "hits": {
    "total": {
      "value": 2942,
      "relation": "eq"
    },
    "max_score": null,
    "hits": [ ]
  },
  "aggregations": {
    "dnaCounts": {
      "doc_count_error_upper_bound": 0,
      "sum_other_doc_count": 0,
      "buckets": [
        {
          "key": "c0ffecef-dad8-481c-ae05-1a55e8724327",
          "doc_count": 2523,
          "countAgg": {
            "value": 2523
          }
        }
        ,
        {
          "key": "e0ffbbba-b49e-4570-a4b1-1cb74e4ad2ba",
          "doc_count": 374,
          "countAgg": {
            "value": 374
          }
        }
        ,
        {
          "key": "f48717b2-a0a3-41c2-a978-f54299a1e42d",
          "doc_count": 45,
          "countAgg": {
            "value": 45
          }
        }
      ]
    }
  }
}
```



### golang实现

```go
// es查询结果model
type EsCountAgg struct {
	DocCountErrorUpperBound int64 `json:"doc_count_error_upper_bound"`
	SumOtherDocCount        int64 `json:"sum_other_doc_count"`
	Buckets                 []struct {
		Key      string `json:"key"`
		DocCount int64  `json:"doc_count"`
		CountAgg struct {
			Value float64 `json:"value"`
		} `json:"countAgg"`
	} `json:"buckets"`
}

var appIds []interface{}
for _, app := range appResList {
        appIds = append(appIds, app.AppId)
    }
appTermsQuery := elastic.NewTermsQuery("app.id", appIds...)
searchDna := l.svcCtx.OEsClient.OClient.Search().Index("dna")


// 查询dna数量，按app.id分组
res, err := searchDna.Query(appTermsQuery).
	Aggregation("dnaCounts", elastic.NewTermsAggregation().Field("app.id").Size(len(appIds)).
    SubAggregation("countAgg", elastic.NewValueCountAggregation().Field("id"))).Size(0).Do(l.ctx)
if err != nil {
        return nil, err
    }
var dnaCounts models.EsCountAgg
err = json.Unmarshal(res.Aggregations["dnaCounts"], &dnaCounts)
if err != nil {
        return nil, err
    }
dnaCountM := make(map[string]int64, len(dnaCounts.Buckets))
for _, dnaCount := range dnaCounts.Buckets {
        dnaCountM[dnaCount.Key] = int64(dnaCount.CountAgg.Value)
    }
for i, app := range appResList {
        appResList[i].AppDnaCounts = dnaCountM[app.AppId]
    }
```