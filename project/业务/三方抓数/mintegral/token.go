package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"time"
)

func main() {
	apiKey := "8a89923f6ad7b65c68cbe6e12249e9fd"
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	fmt.Println(timestamp)
	token := md5.Sum([]byte(apiKey + fmt.Sprintf("%x", md5.Sum([]byte(timestamp)))))
	fmt.Printf("%x\n", token)
}
