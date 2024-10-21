package main

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"
)

func GenerateSignature(signatureNonce string, appKey string, timestamp int64) string {
	data := fmt.Sprintf("%s%s%d", signatureNonce, appKey, timestamp)
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func main() {
	tempByte := make([]byte, 8)
	rand.Read(tempByte)
	signatureNonce := hex.EncodeToString(tempByte)
	fmt.Println("nonce:", signatureNonce)
	appKey := "8ddcffff3a80f4f4189ca1c9d4d902c3c909"
	timestamp := time.Now().Unix()
	fmt.Println("signature:", GenerateSignature(signatureNonce, appKey, timestamp))
}
