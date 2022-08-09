/**
    @author: huChao
    @since: 2022/8/9
    @desc: //TODO The official Go client for Elasticsearch
**/
package main

import (
	"github.com/elastic/go-elasticsearch/v8"
	"log"
)

func main() {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	res0, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting  response: %s", err)
	}

	defer res0.Body.Close()
	log.Println(res0)

	log.Println(elasticsearch.Version)
	log.Println(es.Info())
}
