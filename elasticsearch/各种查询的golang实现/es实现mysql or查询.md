# es实现mysql or查询



```go
termQuery3 := elastic.NewMatchQuery("from_user_bid", req.OwnerBid)
termQuery4 := elastic.NewMatchQuery("to_user_bid", req.OwnerBid)

boolQuery := elastic.NewBoolQuery()

boolQuery.Should(termQuery3)
boolQuery.Should(termQuery4)

searchResult, err := l.svcCtx.OEsClient.OClient.Search().Index("records").Query(boolQuery).From((req.PageNum-1)*req.PageSize).Size(req.PageSize).Sort("create_time", false).Do(l.ctx)
	if err != nil {
		return nil, err
	}
```