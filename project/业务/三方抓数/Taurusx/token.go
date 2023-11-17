package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

func GenerateHash(apiKey string) string {
	// 获取当前时间戳
	timestamp := time.Now().Unix()

	// 对时间戳进行MD5哈希
	timestampHashInBytes := md5.Sum([]byte(fmt.Sprintf("%d", timestamp)))
	timestampHash := hex.EncodeToString(timestampHashInBytes[:])

	// 将API密钥和MD5哈希后的时间戳连接，并计算MD5哈希值
	data := fmt.Sprintf("%s%s", apiKey, timestampHash)
	hashInBytes := md5.Sum([]byte(data))
	hash := hex.EncodeToString(hashInBytes[:])

	return hash
}

func main() {
	apiKey := "e1a476536eac4e60b727b570c7140be5"
	result := GenerateHash(apiKey)

	fmt.Println("Generated Hash:", result)
}
