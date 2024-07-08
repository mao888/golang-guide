package main

import (
	"fmt"
	"net/http"
	"sync"
)

// checkUrl
func checkUrl(url string, wg *sync.WaitGroup, reachableChan, unreachableChan chan<- string) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		unreachableChan <- url
		return
	}
	resp.Body.Close()
	reachableChan <- url
}

func main() {
	// 给 10 个 url 例子
	urls := []string{
		"https://www.baidu.com",
		"https://www.google.com",
		"https://www.bing.com",
		"https://www.yahoo.com",
		"https://www.yandex.com",
		"https://www.duckduckgo.com",
		"https://www.ask.com",
		"https://www.aol.com",
		"https://www.ask.com",
		"https://www.aol.com",
	}

	var wg sync.WaitGroup
	reachableChan := make(chan string, len(urls))
	unreachableChan := make(chan string, len(urls))
	for _, url := range urls {
		wg.Add(1)
		go checkUrl(url, &wg, reachableChan, unreachableChan)
	}
	wg.Wait()
	close(reachableChan)
	close(unreachableChan)
	//fmt.Println("resultChan:", resultChan)

	for url := range reachableChan {
		fmt.Println("通的url:", url)
	}

	for url := range unreachableChan {
		fmt.Println("不通的url:", url)
	}
}
