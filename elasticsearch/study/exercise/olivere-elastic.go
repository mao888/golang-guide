/**
    @author: HuChao
    @since: 2022/8/11
    @desc: //TODO  An Elasticsearch client for the go进阶14讲 programming language.
			 TODO github: https://github.com/olivere/elastic/
             TODO https://olivere.github.io/elastic/
**/
package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/config"
	"time"
)
import "github.com/olivere/elastic"

type Tweet struct {
	User     string                `json:"user"`
	Message  string                `json:"message"`
	Retweets int                   `json:"retweets"`
	Image    string                `json:"images,omitempty"`
	Created  time.Time             `json:"created,omitempty"`
	Tags     []string              `json:"tags,omitempty"`
	Location string                `json:"location,omitempty"`
	Suggest  *elastic.SuggestField `json:"suggest_field,omitempty"`
}

const mapping = `
{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings":{
		"tweet":{
			"properties":{
				"user":{
					"type":"keyword"
				},
				"message":{
					"type":"text",
					"store": true,
					"fielddata": true
				},
				"images":{
					"type":"keyword"
				},
				"created":{
					"type":"date"
				},
				"tags":{
					"type":"keyword"
				},
				"location":{
					"type":"geo_point"
				},
				"suggest_field":{
					"type":"completion"
				}
			}
		}
	}
}`

func main() {
	client, err := elastic.NewClientFromConfig(&config.Config{
		URL:      "http://47.102.154.244:9200",
		Username: "elastic",
		Password: "dna_es_Passw0rd",
	})
	if err != nil {
		panic(err)
	}

	info, code, err := client.Ping("http://47.102.154.244:9200").Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	// Getting the ES version number is quite common, so there's a shortcut  获取ES版本号是很常见的，所以有一个快捷方式
	esversion, err := client.ElasticsearchVersion("http://47.102.154.244:9200")
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)

	// Use the IndexExists service to check if a specified index exists.  使用IndexExists服务检查指定的索引是否存在
	exists, err := client.IndexExists("twitter").Do(context.Background())
	if err != nil {
		panic(err)
	}
	if !exists {
		// Create a new index.
		createIndex, err := client.CreateIndex("twitter").BodyString(mapping).Do(context.Background())
		if err != nil {
			panic(err)
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
		}
	}

	// Index a tweet (using JSON serialization)  索引tweet(使用JSON序列化)
	tweet1 := Tweet{
		User:     "huchao",
		Message:  "123456",
		Retweets: 0,
	}
	put1, err := client.Index().
		Index("twitter").
		Type("tweet").
		Id("1").
		BodyJson(tweet1).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)

	// Index a second tweet (by string)  索引第二个tweet(按字符串)
	tweet2 := `{"user" : "huchao2", "message" : "It's a Raggy Waltz"}`
	put2, err := client.Index().Index("twitter").Type("tweet").Id("2").BodyString(tweet2).Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Indexed tweet %s to index %s, type %s\n", put2.Id, put2.Index, put2.Type)

	// Get tweet with specified ID
	get1, err := client.Get().
		Index("twitter").
		Type("tewwt").
		Id("1").
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	if get1.Found {
		fmt.Printf("Got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)
	}

	// Flush to make sure the documents got written. 为了确保文件都写好了
	_, err = client.Flush().Index("twitter").Do(context.Background())
	if err != nil {
		panic(err)
	}

	// Search with a term query  使用术语查询进行搜索
	termQuery := elastic.NewTermQuery("user", "huchao")
	searchResult, err := client.Search().
		Index("twitter").        // search in index "twitter"
		Query(termQuery).        // specify the query
		Sort("user", true).      // sort by "user" field, ascending
		From(0).Size(10).        // take documents 0-9
		Pretty(true).            // pretty print request and response JSON
		Do(context.Background()) // execute
	if err != nil {
		panic(err)
	}

	// searchResult is of type SearchResult and returns hits, suggestions,
	// and all kinds of other information from Elasticsearch.
	fmt.Printf("Query took %d milliseconds\n", searchResult.TookInMillis)

	// Update a tweet by the update API of Elasticsearch. 使用Elasticsearch的更新API更新tweet。
	// We just increment the number of retweets. 我们只是增加了转发的数量
	update, err := client.Update().
		Index("twitter").
		Type("tweet").
		Id("1").
		Script(elastic.NewScriptInline("ctx._source.retweets += params.num").Lang("painless").Param("num", 1)).
		Upsert(map[string]interface{}{"retweets": 0}).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("New version of tweet %q is now %d\n", update.Id, update.Version)

	// Delete an index.
	deleteIndex, err := client.DeleteIndex("twitter").Do(context.Background())
	if err != nil {
		panic(err)
	}
	if !deleteIndex.Acknowledged {
		// Not acknowledged
	}
}
