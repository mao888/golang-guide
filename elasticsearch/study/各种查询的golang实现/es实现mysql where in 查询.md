# es实现mysql where in 查询



```go
var seriesList []*models.Series
for _, hit := range seriesSearchResult.Hits.Hits {

            var serie models.Series
            err := json.Unmarshal(hit.Source, &serie)
            if err != nil {
                return nil, err
            }
            seriesList = append(seriesList, &serie)
        }

// 系列Id集合
seriesIdList := make([]interface{}, len(seriesList))
for index, value := range seriesList {
            seriesIdList[index] = value.Id
        }

// DNA资产数
termQuery := elastic.NewBoolQuery()
termQuery := elastic.NewTermsQuery("series_id", seriesIdList...)
//termQuery.Must(elastic.NewTermsQuery("series_id", seriesIdList...))

dnasSearchResult, err := l.svcCtx.OEsClient.OClient.Count().Index("dna").Query(termQuery).Do(l.ctx)
if err != nil {
            return nil, err
        }
appResList[i].AppDnaCounts = dnasSearchResult
fmt.Println("dnasSearchResult", dnasSearchResult)
fmt.Println("appResList[i].AppDnaCounts", appResList[i].AppDnaCounts)
```