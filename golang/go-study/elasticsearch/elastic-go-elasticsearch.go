/**
    @author: huChao
    @since: 2022/8/9
    @desc: //TODO https://github.com/elastic/go-elasticsearch/tree/7.10
**/
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"log"
	"strconv"
	"strings"
	"sync"
)

func main() {
	log.SetFlags(0)

	var (
		r  map[string]interface{}
		wg sync.WaitGroup
	)
	// Initialize a client with the default settings.  使用默认设置初始化客户端
	//
	// An `ELASTICSEARCH_URL` environment variable will be used when exported.
	//
	es, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://47.102.154.244:9200"},
		Username:  "elastic",
		Password:  "dna_es_Passw0rd",
	})
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	//// 1. Get cluster info 得到集群信息
	////
	//res, err := es.Info()
	//if err != nil {
	//	log.Fatalf("Error getting response: %s", err)
	//}
	//defer res.Body.Close()
	//
	//// Check response status
	//if res.IsError() {
	//	log.Fatalf("Error: %s", res.String())
	//}
	////Deserialize the response into a map. 将响应反序列化为映射
	//if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
	//	log.Fatalf("Error parsing the response body: %s", err)
	//}
	//
	////Print client and server version numbers. 打印客户端和服务器版本号
	//log.Printf("Client: %s", elasticsearch.Version)
	//log.Printf("Server, %s", r["version"].(map[string]interface{})["number"])
	//log.Println(strings.Repeat("~", 37))

	// 2. Index documents concurrently	索引文档并发
	//
	for i, title := range []string{"Test One", "Test Two"} {
		wg.Add(1)

		go func(i int, title string) {
			defer wg.Done()

			// Build the request body.  构建请求体
			data, err := json.Marshal(struct {
				Title string
			}{Title: title})
			if err != nil {
				log.Fatalf("Error marshaling document: %s", err)
			}

			// Set up the request object.  设置请求对象
			req := esapi.IndexRequest{
				Index:      "test",
				DocumentID: strconv.Itoa(i + 1),
				Body:       bytes.NewReader(data),
				Refresh:    "true",
			}

			// Perform the request with the client. 用客户端执行请求
			res, err := req.Do(context.Background(), es)
			if err != nil {
				log.Fatalf("Error getting response: %s", err)
			}
			defer res.Body.Close()

			if res.IsError() {
				log.Printf("[%s] Error indexing document ID=%d", res.Status(), i+1)
			} else {
				// Deserialize the response into a map. 将响应反序列化为映射
				var r map[string]interface{}
				if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
					log.Fatalf("Error parsing the response body: %s", err)
				} else {
					// Print the response status and indexed document version. 打印响应状态和索引文档版本
					//log.Printf("[%s] %s; version=%d", res.Status(), r["resule"], int(r["_result"].(float64)))
					println("打印响应状态和索引文档版本")
				}
			}
		}(i, title)
	}
	wg.Wait()

	log.Println(strings.Repeat("-", 37))

	// 3. Search for the indexed documents	搜索索引的文档
	//
	// Build the request body. 构建请求体
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"title": "test",
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	// Perform the search request.  执行搜索请求
	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("test"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["types"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	log.Println("Body=", r)

	// Print the response status, number of results, and request duration.  打印响应状态、结果数量和请求持续时间
	log.Printf(
		"[%s]; took: %dms",
		res.Status(),
		//int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["values"].(float64)),
		int(r["took"].(float64)),
	)
	// Print the ID and document source for each hit.  打印每个命中的ID和文档源
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	}

	log.Println(strings.Repeat("=", 37))
}
