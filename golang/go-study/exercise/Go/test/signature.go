package main

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"sort"
)

func main() {
	s1 := Signature("1760312311655437", "1688719212", "1236080353")
	s2 := Signature("1760312311655437", "1690803473", "1236080353")
	s3 := Signature("52F118BF3C8596B44E31BFD45D4B15B7", "1690803473", "63849394")

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}

// Signature 对密钥token，timestamp和nonce三个string进行签名 生成 signature
//1. 对密钥token，timestamp和nonce三个string进行升序排序
//2. 将三个string转成bytes放入一个buffer
//3. 对byte buffer进行sha1 哈希加密
//4. 哈希加密后转成string即为signature
//5. 然后将string拼在url上
func Signature(token string, timestamp string, nonce string) string {
	strList := []string{token, timestamp, nonce}
	sort.Strings(strList)
	var buffer bytes.Buffer
	for _, str := range strList {
		buffer.WriteString(str)
	}
	h := sha1.New()
	h.Write(buffer.Bytes())
	r := fmt.Sprintf("%x", h.Sum(nil))
	return r
}
